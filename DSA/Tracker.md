
# Topics              Status
--------------------------------
Stack                 [DONE]
Map & sync.Map        [DONE]
2D Matrix             [...]
Recursion             [DONE]
BackTracking          []
Slice/String          []
Searching             []
Sorting               [...]
PriorityQ/heap        []
Queue                 []
Set                   []
Trees                 []
Two Pointers          []
Graph                 []
Tries                 []
GreedyAlgorithms      []
DP                    []
LinkedList            []




sync.NoCopy : It's a runtime primitive type which is used in all the sync packages 
- It prevents any struct form being copied, the runtime checks this automatically 
- sync.WaitGroup, sync.Mutex, sync.cond all have sync.NoCopy
- sync.Pool also uses it and therefore we cannot be copied.


1. print all prefix/suffix of a string
2. string matching algorithm - like KMP
3. All main sorting algorithms.
4. 


Sort string characters? → slices.Sort([]rune)
Sort slice of strings? → sort.Strings([]string)


- Make vs New 
- Representing inf
x := math.Inf(-1)
if math.IsInf(x, -1) {
	fmt.Println("hurrey")
}
fmt.Println(x)




- Date: (7,8,9,10): 
	-> Complete Golang revision and practice interview questions (Morning 2 hrs daily)
	-> Composition vs Inheritance
	-> Coupling vs cohesion
	-> DSA : 2D Matrix and BackTracking

- Date:(11,12,13,14,15,16,17): 
	-> Design patterns
		- SOLID Principle
		- Singleton 
		- Decorator
		- Strategy
		- Builder aka functional Options
		- Factory
	-> Parking lot problem
	-> rate-limiter problem
