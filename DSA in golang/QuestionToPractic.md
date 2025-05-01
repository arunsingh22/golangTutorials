

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