/*Package statistics defines a statistics logging API specific for the nbd server

IOPS and throughput statistics loggers

For logging IOPS and throughput statistics, 2 functions are provided:
StartIOPSThroughputRead for logging the read statistics and StartIOPSThroughputWrite for logging the write statistics.
These functions are used for sending IOPS and throughput statistics based on how many bytes that have been sent to the logger.

When callling these functions, 3 goroutines are created. One that listens for incoming data and collects/aggregates it,
another that at a predetermined interval calculates the IOPS ((part of) blocks per second) and throughput (in kB/s)
based on the data from the aggregating goroutine and broadcasts them using the BroadcastStatistics function from the log package.
and when a valid config source (config.Source) is provided another goroutines is lauched to listen for updates on the cluster ID.
If the source is not valid or nil, an error will be logged and the cluster ID will be omitted
from the tags when the statistics are broadcasted.

The functions return a Logger interface that allows for interaction with the logger.
Logger.Send provides a way to push data (amount of bytes written or read) to the logger.
Logger.Stop cancels the internal context and closes the running goroutines.

usage example:
	blockSize := int64(4096)
	vdiskID := "testVdisk"
	readLogger := StartIOPSThroughputRead(nil, vdiskID, blockSize)
	defer readLogger.Close()

	// log that only half a block was read (2048 bytes)
	readLogger.Send(2048)

	// with an interval set at 1 minute, this will output (when interval is reached):
	// 10::vdisk.iops.read@virt.testVdisk:0.008333333333333333|A
	// 10::vdisk.throughput.read@virt.testVdisk:0.03333333333333333|A
*/
package statistics