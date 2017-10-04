package stor

import (
	"fmt"

	"github.com/zero-os/0-Disk/config"
	"github.com/zero-os/0-Disk/log"
)

// ConfigFromConfigSource creates tlog stor client config from config.Source
func ConfigFromConfigSource(source config.Source, vdiskID, privKey string, dataShards, parityShards int) (conf Config, err error) {
	// read vdisk config
	vdiskConf, err := config.ReadVdiskTlogConfig(source, vdiskID)
	if err != nil {
		log.Errorf("failed to read vdisk tlog config for vdisk `%v`: %v", vdiskID, err)
		return
	}

	// read zerostor config of this vdisk
	zsc, err := config.ReadZeroStoreClusterConfig(source, vdiskConf.ZeroStorClusterID)
	if err != nil {
		log.Errorf("failed to read ZeroStorCluster config for vdisk `%v`: %v", vdiskID, err)
		return
	}

	minServerNum := dataShards + parityShards
	serverNum := len(zsc.Servers)
	if minServerNum > len(zsc.Servers) {
		err = fmt.Errorf("number of zerostor servers (%v) is less than data + parity (%v)",
			serverNum, minServerNum)
		return
	}

	// creates stor config
	serverAddrs := func() (addrs []string) {
		for _, s := range zsc.Servers {
			addrs = append(addrs, s.Address)
		}
		return
	}()

	metaServerAddrs := func() (addrs []string) {
		for _, s := range zsc.MetadataServers {
			addrs = append(addrs, s.Address)
		}
		return
	}()

	return Config{
		VdiskID:         vdiskID,
		Organization:    zsc.IYO.Org,
		Namespace:       zsc.IYO.Namespace,
		IyoClientID:     zsc.IYO.ClientID,
		IyoSecret:       zsc.IYO.Secret,
		ZeroStorShards:  serverAddrs,
		MetaShards:      metaServerAddrs,
		DataShardsNum:   dataShards,
		ParityShardsNum: parityShards,
		EncryptPrivKey:  privKey,
	}, nil
}
