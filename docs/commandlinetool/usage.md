# Using the 0-Disk Command Line Tool

```
$ g8stor -h
g8stor manages the Zero-OS visks

Find more information at github.com/zero-os/0-Disk/g8stor.

Usage:
  g8stor [command]

Available Commands:
  copy        Copy a Zero-OS vdisk
  delete      Delete a Zero-OS resource
  help        Help about any command
  list        List Zero-OS resources
  restore     Restore a Zero-OS resource
  version     Output the version information

Flags:
  -h, --help      help for g8stor
  -v, --verbose   log available information

Use "g8stor [command] --help" for more information about a command.
```

## Examples

Config file used in examples where the config file is used:

```yaml
storageClusters:
  clusterA:
    dataStorage:
      - address: localhost:6379
    metaDataStorage:
      address: localhost:6379
  clusterB:
    dataStorage:
      - address: localhost:6380
    metaDataStorage:
      address: localhost:6380
vdisks:
  vdiskA:
    blockSize: 4096
    size: 1
    storageCluster: clusterA
    type: boot
  vdiskC:
    blockSize: 4096
    size: 1
    storageCluster: clusterB
    type: boot
```

Note that the examples below don't show all available flags.
Please use `g8stor [command] --help` to see the flags of a specific command,
`g8stor copy vdisk --help` will for example show all information available for the
command used to copy a vdisk.

## Copy a vdisk

To copy `vdiskA` as a new vdisk (`vdiskB`) on the _same_ storage cluster (`clusterA`), I would do:

```
$ g8stor copy vdisk vdiskA vdiskB
```

Which would be the same as the more explicit version:

```
$ g8stor copy vdisk vdiskA vdiskB clusterA --config config.yml
```

To copy `vdiskA` as a new vdisk (`vdiskA`) on a _different_ storage cluster (`clusterB`), I would do:

```
$ g8stor copy vdisk vdiskA vdiskA clusterB
```

The following command would be illegal, and abort with an error:

```
$ g8stor copy vdisk vdiskA vdiskA
```

## Delete vdisks

To delete all vdisks listed in the config file:

```
$ g8stor delete vdisks
```

Which is the less explicit version of:

```
$ g8stor delete vdisks --config config.yml
```

To delete only 1 (or more) vdisks, rather then all, we can specify their id(s):

```
$ g8stor delete vdisks vdiskC --config.yml
```

With this knowledge we can write the first delete example even more explicit:

```
$ g8stor delete vdisks vdiskA vdiskC --config.yml
```

The following would succeed for the found vdisk, but log an error for the other vdisk as that one can't be found:

```
$ g8stor delete vdisks foo vdiskA # vdiskA will be deleted correctly, even though foo doesn't exist
```

### Restore a (deduped or nondeduped) vdisk

Restore vdisk `a`:

```
$ g8stor restore vdisk a
```

**Note**: this requires that you have a `config.yml` file in the current working directory.

### List all available vdisks

List vdisks available on `localhost:16379`:

```
$ g8stor list vdisks localhost:16379
```

#### WARNING

This command is slow, and might take a while to finish!
It might also decrease the performance of the ARDB server
in question, by locking the server down for each operation.

## Legacy Examples

### Copy metadata of a deduped vdisk

vdisk `a` and `b` are in the same ardb (`localhost:16379`):

```
$ g8stor copy deduped a b localhost:16379
```

vdisk `a` and `b` are in different ardbs (`localhost:16379` -> `localhost:16380`):

```
$ g8stor copy deduped a b localhost:16379 localhost:16380
```

vdisk `a` and `b` are in different ardb databases (`localhost:16379 DB=0` -> `localhost:16379 DB=1`):

```
$ g8stor copy deduped a b localhost:16379 --targetdb 1
```

### Delete metadata of a deduped vdisk

Delete vdisk `a`:

```
$ g8stor delete deduped a localhost:16379
```