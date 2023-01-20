Usecase of  middleware
-------------------
Logging 
Authentication and Authorization 
Rate Limiting 
Caching 
Header manipulation 
Request/Reponse interception.(validation of header, params, body, or changing the protocol)


## Unit Testing
---------------------
### What is not a unit test 
1. If it talks to database 
2. If it communicates over network 
3. If it touchs to Filesystem 
4. If it can run in parallel with other unit test. 


### How unit test helps 

- go test -v 
- go test -run . -v // . represents all TestXXX in _test.go file.
- go test -run Ints -v // this only runs TestInts in _test.go file. 
- go test -run Foo -v  // this only runs TestFoo in _test.go file. 

### SubTests in Table of testing 
- with subtests we can run very specific test cases 
example : go test -run Ints/number -v 