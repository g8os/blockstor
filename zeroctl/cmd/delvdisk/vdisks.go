package delvdisk

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zero-os/0-Disk/config"
	"github.com/zero-os/0-Disk/log"
	cmdconfig "github.com/zero-os/0-Disk/zeroctl/cmd/config"
)

var vdisksCfg struct {
	ConfigPath string
	Force      bool
}

// VdisksCmd represents the vdisks delete subcommand
var VdisksCmd = &cobra.Command{
	Use:   "vdisks [vdiskid...]",
	Short: "Delete one, multiple or all vdisks",
	RunE:  deleteVdisks,
}

func deleteVdisks(cmd *cobra.Command, args []string) error {
	logLevel := log.ErrorLevel
	if cmdconfig.Verbose {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	log.Infof("loading config %s...", vdisksCfg.ConfigPath)

	cfg, err := config.ReadConfig(vdisksCfg.ConfigPath, config.NBDServer)
	if err != nil {
		return err
	}

	vdisks, err := getVdisks(cfg, args)
	if err != nil {
		return err
	}

	// store all deduped and nondeduped vdisks
	// in a map, where the map-key is the ardb's connection if
	// and the values are the ids of the vdisks stored on that connection string
	dedupedVdisksMetadata := make(map[config.StorageServerConfig][]string)
	nondedupedVdisks := make(map[config.StorageServerConfig][]string)

	var vdiskids []string
	var storageType config.StorageType

	log.Info("sorting all target vdisks by storage type and connection")
	for vdiskID, vdisk := range vdisks {
		// storageCluster is guaranteed to exist by the config module
		cluster := cfg.StorageClusters[vdisk.StorageCluster]

		switch storageType = vdisk.StorageType(); storageType {
		case config.StorageDeduped:
			// metadataStorage is guaranteed to exist for deduped storage
			vdiskids = dedupedVdisksMetadata[*cluster.MetadataStorage]
			dedupedVdisksMetadata[*cluster.MetadataStorage] = append(vdiskids, vdiskID)

		case config.StorageNonDeduped:
			for _, storage := range cluster.DataStorage {
				vdiskids = nondedupedVdisks[storage]
				nondedupedVdisks[storage] = append(vdiskids, vdiskID)
			}

		default: // shouldn't happen
			return fmt.Errorf("invariant: vdisk %s has unknown storage type %d",
				vdiskID, storageType)
		}
	}

	if len(dedupedVdisksMetadata) > 0 {
		log.Info("deleting metadata of selected deduped vdisks...")
		for cfg, vdiskids := range dedupedVdisksMetadata {
			err = deleleDedupedVdisksMetadata(vdisksCfg.Force, cfg, vdiskids...)
			if err != nil {
				return err
			}
		}
	}

	if len(nondedupedVdisks) > 0 {
		log.Info("deleting data of selected nondeduped vdisks...")
		for cfg, vdiskids := range nondedupedVdisks {
			err = deleleNondedupedVdisks(vdisksCfg.Force, cfg, vdiskids...)
			if err != nil {
				return err
			}
		}
	}

	log.Info("done")
	return nil
}

func getVdisks(cfg *config.Config, args []string) (map[string]config.VdiskConfig, error) {
	if len(args) == 0 {
		return cfg.Vdisks, nil
	}

	// create a vdisk map, so we only have each id once
	vdiskids := make(map[string]struct{}, len(args))
	for _, vdiskid := range args {
		vdiskids[vdiskid] = struct{}{}
	}

	log.Info("retreiving given vdisks from config file...")

	// collect all vdisk configurations (once)
	vdisks := make(map[string]config.VdiskConfig)
	for candidateID := range vdiskids {
		for vdiskID := range cfg.Vdisks {
			if vdiskID == candidateID {
				vdisks[vdiskID] = cfg.Vdisks[vdiskID]
				delete(cfg.Vdisks, vdiskID)
				delete(vdiskids, vdiskID)
				break
			}
		}

		if len(cfg.Vdisks) == 0 {
			break
		}
	}

	if len(vdisks) == 0 {
		return nil, errors.New("no vdisks could be found for the given vdiskids")
	}

	if !vdisksCfg.Force && len(vdiskids) > 0 {
		// abort non-forced command,
		// in case least one given vdisk could not be found

		message := "following vdisk(s) could not be found in the config file:"
		for vdiskID := range vdiskids {
			message += " " + vdiskID + ","
		}
		return nil, errors.New(message[:len(message)-1])
	}

	// log an error for each invalid vdisk id
	for vdiskID := range vdiskids {
		log.Errorf("vdisk %s could not be found and will thus not be deleted", vdiskID)
	}

	return vdisks, nil
}

func init() {
	VdisksCmd.Long = VdisksCmd.Short + `

When no vdiskids are specified,
all vdisks listed in the config file will be deleted.

WARNING: until issue #88 has been resolved,
  only the metadata of deduped vdisks can be deleted by this command.
  Nondeduped vdisks have no metadata, and thus are not affected by this issue.
`

	VdisksCmd.Flags().StringVar(
		&vdisksCfg.ConfigPath, "config", "config.yml",
		"zeroctl config file")
	VdisksCmd.Flags().BoolVarP(
		&vdisksCfg.Force, "force", "f", false,
		"when enabled non-fatal errors are logged instead of aborting the command")
}
