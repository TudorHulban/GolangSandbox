### Performance Results:
#### With Mutexes:
Benchmark_WMutex-8   	70632560	        16.5 ns/op	       0 B/op	       0 allocs/op


#### With Channels:
Benchmark_WChannels-8   	 2808022	       400 ns/op	       0 B/op	       0 allocs/op


#### With Atomic:
Benchmark_WAtomic-8   	177026658	         6.74 ns/op	       0 B/op	       0 allocs/op