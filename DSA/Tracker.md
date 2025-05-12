
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
PriorityQ/heap        [...]
Queue                 []
LinkedList            []
Set                   []
Trees                 []
Two Pointers          []
Graph                 []
GreedyAlgorithms      []
DP                    []
Tries                 []




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