##                      Online Documents to Read
------------------------------------------------------------------------------------------
1. https://medium.com/rungo : MUST READ 

2. Structures in GO: https://medium.com/rungo/structures-in-go-76377cc106a2 
3. Interfaces in GO: https://medium.com/rungo/interfaces-in-go-ab1601159b3a 


4. Error Handling in GO: 
    https://medium.com/rungo/error-handling-in-go-f0125de052f0#:~:text=Go%20does%20not%20provide%20conventional%20try%2Fcatch%20method%20to%20handle,as%20a%20normal%20return%20value.&text=error%20is%20a%20built%2Din,check%20for%20the%20nil%20condition. 

###  It pushes to origin's URL the local feature2 branch as remotefeature2 branch.
-   git push origin feature2:remotefeature2

- git log --oneline
- git remote -v 
- git 

# -a flag lets skip git add. cmd and directly commits things.
# git commit -a -m "first commit"

# git branch 
# git branch --merged
# git branch --no-merged
# git checkout -b <brancName> 
# git branch <branchName>
# git checkout <branchName>


## NOTE:
---------------------------------------------------------
- Nil is a type of zero: It indicates the absence of something and almost every built-ins are safe: len,cap,range
- "Make the zero value useful" : Rob Pike

- var s []int // nil slice
- var m map[string]int // nil map
- l:= len(s) // length of nil slice is 0

## creates a slice of len and cap 5 and makes s point to it.
- var s = make([]int,5) 

## 
- var m = make(map[string]interface{}) 
- preincrement and predecrement not supported in golang.



