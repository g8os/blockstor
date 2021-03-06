package backup

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zero-os/0-Disk/errors"
	"github.com/zero-os/0-Disk/log"
	"github.com/zero-os/0-Disk/nbd/ardb/backup"

	cmdconfig "github.com/zero-os/0-Disk/zeroctl/cmd/config"
)

// DescribeSnapshotCmd represents the describe-snapshot subcommand
var DescribeSnapshotCmd = &cobra.Command{
	Use:   "snapshot snapshotID",
	Short: "describe a snapshot",
	RunE:  describeSnapshot,
}

// describe only configuration
// see `init` for more information
// about the meaning of each config property.
var describeVdiskCmdCfg struct {
	PrettyPrint bool
}

func describeSnapshot(cmd *cobra.Command, args []string) error {
	logLevel := log.ErrorLevel
	if cmdconfig.Verbose {
		logLevel = log.DebugLevel
	}
	log.SetLevel(logLevel)

	// parse the position arguments
	err := parseDescribePosArguments(args)
	if err != nil {
		return err
	}

	// create backup storage config based on our flags
	backupStoragDriverConfig := createBackupStorageConfigFromFlags()

	// read snapshot's header
	header, err := backup.ReadSnapshotHeader(
		vdiskCmdCfg.SnapshotID, backupStoragDriverConfig,
		&vdiskCmdCfg.PrivateKey, vdiskCmdCfg.CompressionType)
	if err != nil {
		return err
	}

	var info SnapshotInfo

	info.SnapshotID = header.Metadata.SnapshotID
	info.BlockSize = header.Metadata.BlockSize
	info.Size = info.BlockSize * header.DedupedMap.Count
	info.Created = header.Metadata.Created
	info.Version = header.Metadata.Version.String()

	if header.Metadata.Source.VdiskID != "" {
		info.Source = &SnapshotSourceInfo{
			VdiskID:   header.Metadata.Source.VdiskID,
			BlockSize: header.Metadata.Source.BlockSize,
			Size:      header.Metadata.Source.Size,
		}
	}

	var bytes []byte
	if describeVdiskCmdCfg.PrettyPrint {
		bytes, err = json.MarshalIndent(info, "", "  \t")
	} else {
		bytes, err = json.Marshal(info)
	}
	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}

// SnapshotInfo describes a snapshot,
// using both required and optional information.
type SnapshotInfo struct {
	SnapshotID string              `json:"snapshotID"`
	BlockSize  int64               `json:"blockSize"`
	Size       int64               `json:"size"`
	Created    string              `json:"created,omitempty"`
	Source     *SnapshotSourceInfo `json:"source,omitempty"`
	Version    string              `json:"version,omitempty"`
}

// SnapshotSourceInfo describes optional information about
// the source of a snapshot.
type SnapshotSourceInfo struct {
	VdiskID   string `json:"vdiskID,omitempty"`
	BlockSize int64  `json:"blockSize,omitempty"`
	Size      int64  `json:"size,omitempty"`
}

func parseDescribePosArguments(args []string) error {
	// validate pos arg length
	argn := len(args)
	if argn < 1 {
		return errors.New("not enough arguments")
	} else if argn > 1 {
		return errors.New("too many arguments")
	}

	vdiskCmdCfg.SnapshotID = args[0]
	return nil
}

func init() {
	DescribeSnapshotCmd.Long = DescribeSnapshotCmd.Short + `

A snapshot will be described in JSON format and written to the STDOUT.
The printed JSON object can have following properties:

+ "snapshotID": the identifier of the snapshot, as defined when exporting a vdisk;
+ "blockSize": the size (in bytes) of each (plain, decompressed and deduped) block that make up the snapshot's data;
+ "size": the total (plain and decompressed) data size (in bytes) of the snapshot, as in "blockSize * blockCount";
+ "created": indicates when this snapshot was created (date+time in format RFC3339);
+ "version": tool version that was used to create this snapshot;
+ "source": information about the vdisk that was exported to create this snapshot;

Note that the snapshot size does not equal a vdisk's size.
A vdisk's (actual) size is defined by its blocksize and the biggest block index stored for that vdisk.
Meaning that if you have a blocksize of 4096 (bytes) and the biggest block index stored is 1000,
the actual size of your vdisk at that moment will be 4100096 bytes or 4004 KiB.
A vdisk's actual size is computed using the following formula:

	vdiskActualSize = vdiskBlockSize * (maxVdiskBlockIndex+1)

This is different from the (actual) size of a snapshot that we output as part of a snapshot's description.
The size of a snapshot is simply the total size that is used to store the snapshot,
and will almost always be a lot lower than the actual size of the vdisk that would be created by importing this vdisk.
The reason being is that we do not care about the actual spreading of the blocks (in terms of their block index),
when computing that size, and instead only care about the size of each block stored and how many blocks we have stored.
A snapshot's size is computed using the following formula:

	snapshotSize = snapshotBlockSize * snapshotBlockCount

Also note that in both the vdisk size and snapshot size we do not take into account
the metadata as part of the size. This is because the metadata is not important in this context,
as we only care about the actual blocks (content) when transferring between the snapshot and vdisk storage format.
The snapshot size is not important, neither accurate, and is computed on the fly while executing this command.

Remember to use the same (snapshot) name,
crypto (private) key and the compression type,
as you used while exporting the backup in question.

The crypto (private) key has a required fixed length of 32 bytes.
If the snapshot wasn't encrypted, no key should be given,
giving a key in this scenario will fail the describe.

  The FTP information is given using the --storage flag,
here are some examples of valid values for that flag:
	+ localhost:22;
	+ ftp://1.2.3.4:200;
	+ ftp://1.2.3.4:200/root/dir;
	+ ftp://user@127.0.0.1:200;
	+ ftp://user:pass@12.30.120.200:3000;
	+ ftp://user:pass@12.30.120.200:3000/root/dir;

Alternatively you can also give a local directory path to the --storage flag,
to backup to the local file system instead.
This is also the default in case the --storage flag is not specified.

  When the --storage flag contains an FTP storage config and at least one of 
--tls-server/--tls-cert/--tls-insecure/--tls-ca flags are given, 
FTPS (FTP over SSL) is used instead of a plain FTP connection. 
This enables describing backups in a private and secure fashion,
discouraging eavesdropping, tampering, and message forgery.
When the configured server does not support FTPS an error will be returned.
`

	DescribeSnapshotCmd.Flags().VarP(
		&vdiskCmdCfg.CompressionType, "compression", "c",
		"the compression type to use, options { lz4, xz }")
	DescribeSnapshotCmd.Flags().VarP(
		&vdiskCmdCfg.PrivateKey, "key", "k",
		"an optional 32 byte fixed-size private key used for encryption when given")

	DescribeSnapshotCmd.Flags().VarP(
		&vdiskCmdCfg.BackupStorageConfig, "storage", "s",
		"ftp server url or local dir path to read the snapshot's header from")

	DescribeSnapshotCmd.Flags().BoolVar(
		&describeVdiskCmdCfg.PrettyPrint, "pretty", false,
		"pretty print output when this flag is specified")

	DescribeSnapshotCmd.Flags().BoolVar(
		&vdiskCmdCfg.TLSConfig.InsecureSkipVerify,
		"tls-insecure", false,
		"when given FTP over SSL will be used without cert verification")
	DescribeSnapshotCmd.Flags().StringVar(
		&vdiskCmdCfg.TLSConfig.ServerName,
		"tls-server", "",
		"certs will be verified when given (required when --tls-insecure is not used)")
	DescribeSnapshotCmd.Flags().StringVar(
		&vdiskCmdCfg.TLSConfig.CertFile,
		"tls-cert", "",
		"PEM-encoded file containing the TLS Client cert (FTPS will be used when given)")
	DescribeSnapshotCmd.Flags().StringVar(
		&vdiskCmdCfg.TLSConfig.KeyFile,
		"tls-key", "",
		"PEM-encoded file containing the private TLS client key")
	DescribeSnapshotCmd.Flags().StringVar(
		&vdiskCmdCfg.TLSConfig.CAFile,
		"tls-ca", "",
		"optional PEM-encoded file containing the TLS CA Pool (defaults to system pool when not given)")

}
