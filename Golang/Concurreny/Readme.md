
# Concurrency in GO
- Concurrency is about design , it's designing your program as a collection of independent process to eventually run in parallel, such that the outcome is ALWAYS the same no matter if it ran sequentially or parrallely .
- Concurrency != Parallelism
- Concurrency doesn't make your program faster ONLY parallesim does.
- Concurrency leads to parrellism  
2. sync pkg contains    
    - WaitGroup (Done)
    - mutex & RWMutex (Done)
    - pool
    - locker
    - Map 
    - Cond
3. sync/atomic pkg.

- Go works on fork-join model.
- Go doesn't have user-level threads it only as GO-routines

## WaitGroup gotcha (sudden unforeseen problems)
--------------------------------------------

- If wg.Done() is not called equalt to #goroutines in wg.Add() then this would cause DeadLOCK. Why because the main goroutine is waiting  infinitely for other goroutiens to return which completed or crashed and would never return.

- If the **wg.Done() > wg.Add()** -> this will lead to panic beacuse the counter would become < 0 (negative).
    - example:
    func main(){
        var wg sync.Waitgroup
        wg.Done()
    }

- Waitgroup are ALWAYS passed by REFERECNE else panic would happen.
- wg.Add(10) is better than adding wg.Add(1) in a for-loop