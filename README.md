Script for collecting metrics from `clickhouse` by `telegraf` to influxdb

build: 
GOOS=linux GOARCH=amd64 go build -o chMon chMon.go

Example of `telegraf` `input_exec` config in `.conf` file.
Change the credentials in the configuration file `config.toml` to your own to access ch-server
Place config.toml with credentials near binary.

If you have replication in your clickhouse cluster - there will be an additional field in output: "replicationQueue".

Select measurement `exec_clickhouse` in `grafana`.

Example of collected measurements in `JSON`:

```
{  
   "asynchronous_metrics":{  
      "generic.heap_size":943718400.0,
      "tcmalloc.pageheap_free_bytes":860438528.0,
      "ReplicasMaxQueueSize":0.0,
      "ReplicasMaxRelativeDelay":0.0,
      "ReplicasMaxMergesInQueue":0.0,
      "tcmalloc.current_total_thread_cache_bytes":24384888.0,
      "ReplicasSumMergesInQueue":0.0,
      "tcmalloc.thread_cache_free_bytes":24384888.0,
      "ReplicasMaxInsertsInQueue":0.0,
      "Uptime":489868.0,
      "tcmalloc.pageheap_unmapped_bytes":0.0,
      "MaxPartCountForPartition":12.0,
      "ReplicasMaxAbsoluteDelay":0.0,
      "tcmalloc.transfer_cache_free_bytes":10431616.0,
      "UncompressedCacheBytes":9860.0,
      "MarkCacheFiles":33.0,
      "UncompressedCacheCells":33.0,
      "ReplicasSumInsertsInQueue":0.0,
      "ReplicasSumQueueSize":0.0,
      "tcmalloc.central_cache_free_bytes":2386872.0,
      "MarkCacheBytes":528.0,
      "generic.current_allocated_bytes":46076496.0
   },
   "events":{  
      "ZooKeeperBytesSent":191949445,
      "IOBufferAllocs":6230,
      "UncompressedCacheMisses":66,
      "ZooKeeperClose":15,
      "ZooKeeperCreate":56,
      "IOBufferAllocBytes":1884823189,
      "MergesTimeMilliseconds":478,
      "CompressedReadBufferBytes":64375556,
      "RWLockAcquiredReadLocks":12936277,
      "SelectedParts":1,
      "MergedUncompressedBytes":84052946,
      "FunctionExecute":24,
      "ZooKeeperMulti":347317,
      "ZooKeeperExceptions":347870,
      "ZooKeeperInit":16,
      "FileOpen":4231,
      "ReadBufferFromFileDescriptorRead":6022,
      "ZooKeeperTransactions":1064947,
      "ZooKeeperExists":10983,
      "ReplicatedPartMerges":8,
      "ZooKeeperRemove":9,
      "WriteBufferFromFileDescriptorWriteBytes":8833563,
      "CannotRemoveEphemeralNode":16,
      "SelectedRanges":1,
      "LeaderElectionAcquiredLeadership":8,
      "ReadCompressedBytes":10947163,
      "CreatedReadBufferOrdinary":900,
      "RWLockReadersWaitMilliseconds":13,
      "MergedRows":686706,
      "ReadBufferFromFileDescriptorReadBytes":13346302,
      "CompressedReadBufferBlocks":2217,
      "ZooKeeperWatchResponse":8,
      "ZooKeeperBytesReceived":42992762,
      "MarkCacheMisses":33,
      "SelectQuery":130,
      "ReplicaPartialShutdown":59,
      "ZooKeeperWaitMicroseconds":11093397193,
      "WriteBufferFromFileDescriptorWrite":465,
      "SelectedMarks":1,
      "Query":139,
      "ZooKeeperList":358190,
      "ContextLock":2670566,
      "CreatedWriteBufferOrdinary":217,
      "ZooKeeperSet":8,
      "ZooKeeperGet":348384
   },
   "metrics":{  
      "InterserverConnection":0,
      "ReadonlyReplica":8,
      "ReplicatedSend":0,
      "ContextLockWait":0,
      "DistributedSend":0,
      "QueryThread":0,
      "Processes":1,
      "BackgroundPoolTask":0,
      "RWLockActiveWriters":0,
      "LeaderReplica":0,
      "ReplicatedChecks":0,
      "RWLockActiveReaders":1,
      "TCPConnection":1,
      "MemoryTracking":8704,
      "ZooKeeperSession":1,
      "LeaderElection":0,
      "DictCacheRequests":0,
      "RWLockWaitingWriters":0,
      "Read":1,
      "ZooKeeperWatch":0,
      "Merge":0,
      "QueryPreempted":0,
      "DiskSpaceReservedForMerge":0,
      "HTTPConnection":0,
      "RWLockWaitingReaders":0,
      "OpenFileForRead":0,
      "ZooKeeperRequest":0,
      "StorageBufferBytes":0,
      "ReplicatedFetch":0,
      "OpenFileForWrite":0,
      "MemoryTrackingForMerges":0,
      "Query":1,
      "EphemeralNode":0,
      "StorageBufferRows":0,
      "Write":0,
      "DelayedInserts":0,
      "SendExternalTables":0,
      "Revision":54385,
      "MemoryTrackingInBackgroundProcessingPool":0
   }
}
```
