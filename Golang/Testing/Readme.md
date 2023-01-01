
# Testing in Go 
-------------------------------------
1. All test files in go have XXXX_test.go suffix and these files can either be in the same 
   pkg or in a differnet pkg.
2. The test methods MUST have the following signature. 
    func TestXXXX(t *testing.T){
        // add your test.
    }
3. Test Pass by default.



# Code coverage 
----------------------------------
1. Running a **go test -cover** finds the part of code which where covered during Unit Testing.
2. Use 2 flags to get visual heatMaps for the code coverage **-coverprofile=c.out -covermode=count**

3. Final output: 
- go test abx_test.go -coverprofile=c.out -covermode=count 
- go tool cover -html=c.out