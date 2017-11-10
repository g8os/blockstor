package restore

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/zero-os/0-Disk/config"
	"github.com/zero-os/0-Disk/errors"
	"github.com/zero-os/0-Disk/log"
	"github.com/zero-os/0-Disk/nbd/ardb/storage"
	"github.com/zero-os/0-Disk/tlog/tlogclient/decoder"
	"github.com/zero-os/0-Disk/tlog/tlogclient/player"
	cmdConf "github.com/zero-os/0-Disk/zeroctl/cmd/config"
)

// vdiskCfg is the configuration used for the restore vdisk command
var vdiskCmdCfg struct {
	SourceConfig config.SourceConfig
	PrivKey      string
	StartTs      int64 // start timestamp
	EndTs        int64 // end timestamp
	Force        bool
}

// VdiskCmd represents the restore vdisk subcommand
var VdiskCmd = &cobra.Command{
	Use:   "vdisk id",
	Short: "Restore a vdisk using a given tlogserver",
	RunE:  restoreVdisk,
}

func restoreVdisk(cmd *cobra.Command, args []string) error {
	// create config source
	configSource, err := config.NewSource(vdiskCmdCfg.SourceConfig)
	if err != nil {
		return err
	}
	defer configSource.Close()

	// parse positional args
	argn := len(args)
	if argn < 1 {
		return errors.New("not enough arguments")
	}
	if argn > 1 {
		return errors.New("too many arguments")
	}

	vdiskID := args[0]

	logLevel := log.InfoLevel
	if cmdConf.Verbose {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)

	err = checkVdiskExists(vdiskID, configSource)
	if err != nil {
		return err
	}

	ctx := context.Background()

	player, err := player.NewPlayer(ctx, configSource, vdiskID, vdiskCmdCfg.PrivKey)
	if err != nil {
		return err
	}

	log.Infof("restoring vdisk with start timestamp=%v end timestamp=%v",
		vdiskCmdCfg.StartTs, vdiskCmdCfg.EndTs)
	lastSeq, err := player.Replay(decoder.NewLimitByTimestamp(vdiskCmdCfg.StartTs, vdiskCmdCfg.EndTs))
	log.Infof("restore finished with last sequence = %v", lastSeq)
	return err
}

// checkVdiskExists checks if the vdisk in question already/still exists,
// and if so, and the force flag is specified, delete the vdisk.
func checkVdiskExists(vdiskID string, configSource config.Source) error {
	// check if vdisk exists
	exists, vdiskType, cluster, err := storage.VdiskExists(vdiskID, configSource)
	if err != nil {
		return errors.Wrapf(err, "couldn't check if vdisk %s already exists", vdiskID)
	}
	if !exists {
		return nil // vdisk doesn't exist, so nothing to do
	}
	if !vdiskCmdCfg.Force {
		return errors.Newf("cannot restore vdisk %s as it already exists", vdiskID)
	}

	// delete vdisk, as it exists and `--force` is specified
	deleted, err := storage.DeleteVdiskInCluster(vdiskID, vdiskType, cluster)
	if err != nil {
		return errors.Wrapf(err, "couldn't delete vdisk %s", vdiskID)
	}
	if !deleted {
		return errors.Newf("couldn't delete vdisk %s for an unknown reason", vdiskID)
	}

	// delete 0-Stor (meta)data for this vdisk
	if vdiskType.TlogSupport() {
		// TODO: also delete actual tlog meta(data) from 0-Stor cluster for the supported vdisks ?!?!
		//       https://github.com/zero-os/0-Disk/issues/147
	}

	// vdisk did exist, but we were able to delete all the exiting (meta)data
	return nil
}

func init() {
	VdiskCmd.Flags().Var(
		&vdiskCmdCfg.SourceConfig, "config",
		"config resource: dialstrings (etcd cluster) or path (yaml file)")
	VdiskCmd.Flags().StringVar(
		&vdiskCmdCfg.PrivKey,
		"priv-key", "12345678901234567890123456789012",
		"private key")
	VdiskCmd.Flags().Int64Var(
		&vdiskCmdCfg.StartTs,
		"start-timestamp", 0,
		"start UTC timestamp in nanosecond(default 0: since beginning)")
	VdiskCmd.Flags().Int64Var(
		&vdiskCmdCfg.EndTs,
		"end-timestamp", 0,
		"end UTC timestamp in nanosecond(default 0: until the end)")
	VdiskCmd.Flags().BoolVarP(
		&vdiskCmdCfg.Force,
		"force", "f", false,
		"when given, delete the vdisk if it already existed")
}
