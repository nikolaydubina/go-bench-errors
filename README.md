# Benchmarking Go errors

How does _creation_, _serialization_, and _wrapping_ changes with different errors?

Notation
* `fmt.Errorf` you wrap like this `fmt.Errorf("my message %s with data %f: %w", err)`
* `DelayedError` you wrap underlying error into field but do not call `e.Err.Error()` immediately
* `errors.New` you create error with string passed directly into struct, no fmt calls made.

Results
1. `fmt.Errorf` wrapping calls `Error` immediately. Delaying `Error` call is not utilized
2. `errors.New` is very fast when no complex formatting (floats/reflection/iterables) is necessary
3. `DelayedError` is 10x faster when `Error` calls are not done, since `Error` calls are lazily delayed
4. `Error` call for `fmt.Errorf` and for `DelayedError` takes same time

Benchmarks
```
$ go test -bench=. -benchtime=5s -benchmem ./...
goos: darwin
goarch: amd64
pkg: github.com/nikolaydubina/go-bench-errors
cpu: VirtualApple @ 2.50GHz
BenchmarkWrap_DelayedError_float64_Message_2wrap-10     	14211424	       419 ns/op	     216 B/op	       8 allocs/op
BenchmarkWrap_DelayedError_float64_2wrap-10             	99927066	        58 ns/op	      96 B/op	       2 allocs/op
BenchmarkWrap_DelayedError_float64_Message_5wrap-10     	 5403166	      1097 ns/op	     736 B/op	      20 allocs/op
BenchmarkWrap_DelayedError_float64_5wrap-10             	38846422	       153 ns/op	     240 B/op	       5 allocs/op
BenchmarkWrap_DelayedError_float64_Message_10wrap-10    	 2672426	      2246 ns/op	    2072 B/op	      40 allocs/op
BenchmarkWrap_DelayedError_float64_10wrap-10            	19889925	       299 ns/op	     480 B/op	      10 allocs/op
BenchmarkWrap_DelayedError_float64_Message_50wrap-10    	  394473	     14857 ns/op	   35814 B/op	     200 allocs/op
BenchmarkWrap_DelayedError_float64_50wrap-10            	 4079131	      1467 ns/op	    2400 B/op	      50 allocs/op
BenchmarkWrap_fmt_Errorf_float64_Message_2wrap-10       	13853347	       430 ns/op	     176 B/op	       6 allocs/op
BenchmarkWrap_fmt_Errorf_float64_2wrap-10               	13966632	       429 ns/op	     176 B/op	       6 allocs/op
BenchmarkWrap_fmt_Errorf_float64_Message_5wrap-10       	 5036449	      1191 ns/op	     664 B/op	      15 allocs/op
BenchmarkWrap_fmt_Errorf_float64_5wrap-10               	 5040122	      1189 ns/op	     664 B/op	      15 allocs/op
BenchmarkWrap_fmt_Errorf_float64_Message_10wrap-10      	 2418132	      2485 ns/op	    2048 B/op	      30 allocs/op
BenchmarkWrap_fmt_Errorf_float64_10wrap-10              	 2414413	      2484 ns/op	    2048 B/op	      30 allocs/op
BenchmarkWrap_fmt_Errorf_float64_Message_50wrap-10      	  353810	     16856 ns/op	   38867 B/op	     150 allocs/op
BenchmarkWrap_fmt_Errorf_float64_50wrap-10              	  353611	     16808 ns/op	   38867 B/op	     150 allocs/op
BenchmarkSingle_fmt_Errorf_float64_Message-10           	27763376	       215 ns/op	      72 B/op	       3 allocs/op
BenchmarkSingle_fmt_Errorf_float64-10                   	30912550	       192 ns/op	      72 B/op	       3 allocs/op
BenchmarkSingle_fmt_Errorf_int_Message-10               	46369344	       129 ns/op	      51 B/op	       3 allocs/op
BenchmarkSingle_fmt_Errorf_int-10                       	56102511	       106 ns/op	      51 B/op	       3 allocs/op
BenchmarkSingle_errors_New_int_Message-10               	100000000	        53 ns/op	       4 B/op	       1 allocs/op
BenchmarkSingle_errors_New_int-10                       	163252323	        36 ns/op	       4 B/op	       1 allocs/op
PASS
ok  	github.com/nikolaydubina/go-bench-errors	149.863s
```
