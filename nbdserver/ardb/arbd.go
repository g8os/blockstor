package ardb

import (
	"errors"
	"fmt"
	"log"

	"golang.org/x/net/context"

	gridapi "github.com/g8os/blockstor/gridapi/gridapiclient"
	"github.com/g8os/blockstor/nbdserver/lba"
	"github.com/g8os/blockstor/nbdserver/storage"
	"github.com/g8os/gonbdserver/nbd"
	"github.com/garyburd/redigo/redis"
)

// shared constants
const (
	// DefaultLBACacheLimit defines the default cache limit
	DefaultLBACacheLimit = 20 * 1024 * 1024 // 20 mB
)

// NewBackendFactory creates a new Backend Factory,
// which is used to create a Backend, without having to work with global variables.
func NewBackendFactory(pool *RedisPool, scConfigFactory *storage.ClusterConfigFactory, gridAPIAddress string, lbaCacheLimit int64) (*BackendFactory, error) {
	if pool == nil {
		return nil, errors.New("NewBackendFactory requires a non-nil RedisPool")
	}
	if scConfigFactory == nil {
		return nil, errors.New("NewBackendFactory requires a non-nil storage.ClusterConfigFactory")
	}
	if gridAPIAddress == "" {
		return nil, errors.New("NewBackendFactory requires a non-empty gridAPIAddress")
	}

	return &BackendFactory{
		backendPool:     pool,
		scConfigFactory: scConfigFactory,
		gridAPIAddress:  gridAPIAddress,
		lbaCacheLimit:   lbaCacheLimit,
	}, nil
}

//BackendFactory holds some variables
// that can not be passed in the exportconfig like the pool of ardbconnections
// I hate the factory pattern but I hate global variables even more
type BackendFactory struct {
	backendPool     *RedisPool
	scConfigFactory *storage.ClusterConfigFactory
	gridAPIAddress  string
	lbaCacheLimit   int64
}

//NewBackend generates a new ardb backend
func (f *BackendFactory) NewBackend(ctx context.Context, ec *nbd.ExportConfig) (backend nbd.Backend, err error) {
	volumeID := ec.Name

	// create storage cluster config,
	// which is used to dynamically reload the configuration
	storageClusterCfg, err := f.scConfigFactory.NewConfig(volumeID)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}

	redisProvider, err := newRedisProvider(f.backendPool, storageClusterCfg)

	//Get information about the volume
	g8osClient := gridapi.NewG8OSStatelessGRID()
	g8osClient.BaseURI = f.gridAPIAddress
	log.Println("[INFO] Starting volume", volumeID)
	volumeInfo, _, err := g8osClient.Volumes.GetVolumeInfo(volumeID, nil, nil)
	if err != nil {
		log.Println("[ERROR]", err)
		return
	}

	var storage backendStorage
	blockSize := int64(volumeInfo.Blocksize)

	switch volumeInfo.Volumetype {
	case gridapi.EnumVolumeVolumetypedb, gridapi.EnumVolumeVolumetypecache:
		storage = newNonDedupedStorage(volumeID, blockSize, redisProvider)
	case gridapi.EnumVolumeVolumetypeboot:
		cacheLimit := f.lbaCacheLimit
		if cacheLimit < lba.BytesPerShard {
			cacheLimit = DefaultLBACacheLimit
		}

		volumeSize := int64(volumeInfo.Size)
		blockCount := volumeSize / blockSize
		if volumeSize%blockSize > 0 {
			blockCount++
		}

		var vlba *lba.LBA
		vlba, err = lba.NewLBA(
			volumeID,
			blockCount,
			cacheLimit,
			redisProvider,
		)
		if err != nil {
			log.Println("[ERROR]", err)
			return
		}

		storage = newDedupedStorage(volumeID, blockSize, redisProvider, vlba)
	default:
		err = fmt.Errorf("Unsupported volume type: %s", volumeInfo.Volumetype)
	}

	backend = &Backend{
		blockSize:         blockSize,
		size:              uint64(volumeInfo.Size),
		storage:           storage,
		storageClusterCfg: storageClusterCfg,
	}
	return
}

// newRedisProvider creates a new redis provider
func newRedisProvider(pool *RedisPool, storageClusterCfg *storage.ClusterConfig) (*redisProvider, error) {
	if pool == nil {
		return nil, errors.New(
			"no redis pool is given, while one is required")
	}
	if storageClusterCfg == nil {
		return nil, errors.New(
			"no storage cluster config is given, while one is required")
	}

	return &redisProvider{
		redisPool:         pool,
		storageClusterCfg: storageClusterCfg,
	}, nil
}

// redisProvider allows you to get a redis connection from a pool
// using a modulo index
type redisProvider struct {
	redisPool         *RedisPool
	storageClusterCfg *storage.ClusterConfig
}

// GetRedisConnection from the underlying pool, using a modulo index
func (rp *redisProvider) RedisConnection(index int) (conn redis.Conn, err error) {
	connString, err := rp.storageClusterCfg.ConnectionString(index)
	if err != nil {
		return
	}

	conn = rp.redisPool.Get(connString)
	return
}

// MetaRedisConnection implements lba.MetaRedisProvider.MetaRedisConnection
func (rp *redisProvider) MetaRedisConnection() (conn redis.Conn, err error) {
	connString, err := rp.storageClusterCfg.MetaConnectionString()
	if err != nil {
		return
	}

	conn = rp.redisPool.Get(connString)
	return
}