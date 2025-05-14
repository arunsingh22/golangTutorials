
Go runtime Sceduler:
- The globalRunnbleQueue is check after every 61 ticks(prime number) and when
  the localRunnableQueue is empty.
- The localRunnableQueue len is 256.
- The goroutine do not have any ID unlike threads 
- The goroutine stack is dynamic starting with 2KB of size. 
[https://medium.com/@sanilkhurana7/understanding-the-go-scheduler-and-looking-at-how-it-works-e431a6daacf]

[https://levelup.gitconnected.com/goroutine-scheduler-revealed-youll-never-see-goroutines-the-same-way-again-3c159b01c25a]

Go GC:
- It's concurrent, tri-color mark and sweep algorithm.
- It used dedicated/seperate OS threads which can run on mulitple cores

# Why would GC impact system performance
- If the rate of memory allocation in the mutator (main program) is very high, then Go GC will start to “steal” more goroutines from the mutator to assist with the marking phase. This has two effects — firstly, it speeds up the GC process by providing more resources, and secondly, it takes away resources from the mutator, which slows down the rate of memory allocation. This is important to ensure that the rate of memory allocation does exceed the rate of memory cleanup, which could cause the heap to grow out of control, potentially resulting in out-of-memory crashes.

- When the garbage collector starts to steal resources from the main program, it can start to have a significant impact on the performance of the main program, since CPU resources are limited. This typically manifests in the form of “tail latency”, i.e, the higher percentiles of latency (p99, p999, etc.) compared to the average latency, and can have an adverse effect on user experience. A user will remember the worst or slowest experiences more than the average request, and this can cause user dissatisfaction. 

- [https://bwoff.medium.com/understanding-gos-garbage-collection-415a19cc485c]
