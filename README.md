Assignment. 
Go through all the actions by next week:
 Please create a separate executable to generate files where you define the size and the count and where to place them the text is coming from https://www.lipsum.com/
CLI command should look like this 
createfiles --size=100MB --count=100 path=/data
write an app that reads all the text files in a directory and count the number of words in all the files in a directory.
requirements

1. first version of the app: reading the files in sequential order and do a benchmark test for it. 
The command should look like this
Calcsize –path=/data
support material
https://hackernoon.com/how-to-write-benchmarks-in-golang-like-an-expert-0w1834gs
https://pkg.go.dev/testing#hdr-Benchmarks
Notes:




2. second version of the app: use go routine per file and per line and accumulate the count in atomic counter and create benchmark test for it, the decision perfile or perline is passed by a command argument. 
The command should look like this
Calcsize1 –path=/data –perfile=true –perline=false
support material
https://gobyexample.com/atomic-counters
https://pkg.go.dev/sync/atomic
Notes:
Atomic vars should always be used with add and load. 


Fan In -> https://techmaster.vn/posts/36200/concurrency-trong-go-fan-in-pattern
Fan Out -> https://medium.com/a-journey-with-go/go-fan-out-pattern-in-our-etl-9357a1731257

3. third version of the app: use a worker pool of configurable size and accumulate the count in an atomic counter and create a benchmark test for it.
The command should look like this
Calcsize1 –path=/data –wpsize=10
support material
https://gobyexample.com/worker-pools 
https://itnext.io/explain-to-me-go-concurrency-worker-pool-pattern-like-im-five-e5f1be71e2b0
https://medium.com/code-chasm/go-concurrency-pattern-worker-pool-a437117025b1
advanced http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
https://golangbot.com/buffered-channels-worker-pools/
https://github.com/mileusna/viber




