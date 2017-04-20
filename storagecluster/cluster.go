package storagecluster

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	log "github.com/glendc/go-mini-log"

	gridapi "github.com/g8os/blockstor/gridapi/gridapiclient"
)

// NewClusterConfigFactory creates a ClusterConfigFactory.
func NewClusterConfigFactory(gridapiaddress string, logger log.Logger) (*ClusterConfigFactory, error) {
	if gridapiaddress == "" {
		return nil, errors.New("NewClusterConfigFactory requires a non-empty gridapiaddress")
	}
	if logger == nil {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	return &ClusterConfigFactory{
		gridapiaddress: gridapiaddress,
		requestCh:      make(chan string),
		responseCh:     make(chan clusterConfigResponse),
		logger:         logger,
	}, nil
}

// ClusterConfigFactory allows for the creation of ClusterConfigs.
type ClusterConfigFactory struct {
	gridapiaddress string
	requestCh      chan string
	responseCh     chan clusterConfigResponse
	logger         log.Logger
}

// NewConfig returns a new ClusterConfig.
func (f *ClusterConfigFactory) NewConfig(vdiskID string) (cfg *ClusterConfig, err error) {
	if vdiskID == "" {
		err = errors.New("ClusterConfig requires a non-empty vdiskID")
		return
	}

	f.requestCh <- vdiskID
	resp := <-f.responseCh

	cfg = resp.Config
	err = resp.Error
	return
}

// Listen to incoming creation requests (send by the NewConfig method)
func (f *ClusterConfigFactory) Listen(ctx context.Context) {
	for {
		select {
		// wait for a request
		case vdiskID := <-f.requestCh:
			cfg, err := NewClusterConfig(
				f.gridapiaddress,
				vdiskID,
				f.logger,
			)
			if err != nil {
				// couldn't create cfg, early exit
				f.responseCh <- clusterConfigResponse{Error: err}
				continue
			}

			cfg.done = make(chan struct{}, 1)
			go cfg.listen(ctx)

			// all fine, return the configuration
			f.responseCh <- clusterConfigResponse{Config: cfg}

		// or until the context is done
		case <-ctx.Done():
			return
		}
	}
}

type clusterConfigResponse struct {
	Config *ClusterConfig
	Error  error
}

// NewClusterConfig creates a new cluster config
func NewClusterConfig(gridapiaddress, vdiskID string, logger log.Logger) (*ClusterConfig, error) {
	client := gridapi.NewG8OSStatelessGRID()
	client.BaseURI = gridapiaddress

	if logger == nil {
		logger = log.New(os.Stderr, "", log.LstdFlags)
	}

	cfg := &ClusterConfig{
		client:  client,
		vdiskID: vdiskID,
		logger:  logger,
		done:    make(chan struct{}, 1),
	}

	if !cfg.loadConfig() {
		return nil, errors.New("couldn't load configuration")
	}

	return cfg, nil
}

// ClusterConfig contains the cluster configuration,
// which gets reloaded based on incoming SIGHUP signals.
type ClusterConfig struct {
	client  *gridapi.G8OSStatelessGRID
	vdiskID string

	// used to log
	logger log.Logger

	// keep type, such that we can check this,
	// when reloading the configuration
	vdiskType gridapi.EnumVdiskType

	// used to get a redis connection
	servers         []gridapi.HAStorageServer
	numberOfServers int //Keep it as a seperate variable since this is constantly needed

	// used to store meta data
	metaConnectionString string

	// indicates if configuration is succesfully loaded
	loaded bool

	// mutex
	mux sync.Mutex

	// used to stop the listening
	done chan struct{}
}

// ConnectionString returns a connectionstring,
// based on a given index, which will be morphed into a local index,
// based on the available storage servers available.
func (cfg *ClusterConfig) ConnectionString(index int) (string, error) {
	cfg.mux.Lock()
	defer cfg.mux.Unlock()

	if !cfg.loaded && !cfg.loadConfig() {
		return "", errors.New("couldn't load storage cluster config")
	}

	bcIndex := index % cfg.numberOfServers
	return connectionStringFromHAStorageServer(cfg.servers[bcIndex]), nil
}

// MetaConnectionString returns the connectionstring (`<host>:<port>`),
// used to connect to the meta storage server.
func (cfg *ClusterConfig) MetaConnectionString() (string, error) {
	cfg.mux.Lock()
	defer cfg.mux.Unlock()

	if !cfg.loaded && !cfg.loadConfig() {
		return "", errors.New("couldn't load storage cluster config")
	}

	return cfg.metaConnectionString, nil
}

// Close the open listen goroutine,
// which autoreloads the internal configuration,
// upon receiving a SIGHUP signal.
func (cfg *ClusterConfig) Close() {
	if cfg.done != nil {
		cfg.done <- struct{}{}
		close(cfg.done)
		cfg.done = nil
	}
}

// listen to incoming signals,
// and reload configuration when receiving a SIGHUP signal.
func (cfg *ClusterConfig) listen(ctx context.Context) {
	cfg.logger.Info("ready to reload StorageClusterConfig upon SIGHUP receival for:", cfg.vdiskID)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP)
	defer signal.Stop(ch)

	for {
		select {
		case s := <-ch:
			switch s {
			case syscall.SIGHUP:
				cfg.logger.Infof("%q received SIGHUP Signal", cfg.vdiskID)
				func() {
					cfg.mux.Lock()
					defer cfg.mux.Unlock()
					cfg.loadConfig()
				}()
			default:
				cfg.logger.Info("received unsupported signal", s)
			}

		case <-cfg.done:
			cfg.logger.Info(
				"exit listener for StorageClusterConfig for vdisk:",
				cfg.vdiskID)
			return

		case <-ctx.Done():
			cfg.logger.Info(
				"abort listener for StorageClusterConfig for vdisk:",
				cfg.vdiskID)
			return
		}
	}
}

func (cfg *ClusterConfig) loadConfig() bool {
	cfg.loaded = false

	cfg.logger.Info("loading storage cluster config")

	// get vdisk info
	vdiskInfo, _, err := cfg.client.Vdisks.GetVdiskInfo(cfg.vdiskID, nil, nil)
	if err != nil {
		cfg.logger.Infof("couldn't get vdiskInfo: %s", err.Error())
		return false
	}

	// check vdiskType, and sure it's the same one as last time
	if cfg.vdiskType != "" && cfg.vdiskType != vdiskInfo.Type {
		cfg.logger.Infof("wrong type for vdisk %q, expected %q, while received %q",
			cfg.vdiskID, cfg.vdiskType, vdiskInfo.Type)
		return false
	}
	cfg.vdiskType = vdiskInfo.Type

	//Get information about the backend storage nodes
	storageClusterInfo, _, err := cfg.client.Storageclusters.GetClusterInfo(vdiskInfo.Storagecluster, nil, nil)
	if err != nil {
		cfg.logger.Infof("couldn't get storage cluster info: %s", err.Error())
		return false
	}

	// store information required for getting redis connections
	cfg.servers = storageClusterInfo.DataStorage
	cfg.numberOfServers = len(cfg.servers)
	if cfg.numberOfServers < 1 {
		cfg.logger.Info(
			"received no storageBackendController, while at least 1 is required")
		return false
	}

	// used to store metadata
	if len(storageClusterInfo.MetadataStorage) < 1 {
		cfg.logger.Infof("No metadata servers available in storagecluster %s", vdiskInfo.Storagecluster)
		return false
	}
	cfg.metaConnectionString = connectionStringFromHAStorageServer(storageClusterInfo.MetadataStorage[0])

	cfg.loaded = true
	return cfg.loaded
}

func connectionStringFromHAStorageServer(server gridapi.HAStorageServer) string {
	return server.Master.Ip + ":" + strconv.Itoa(server.Master.Port)
}
