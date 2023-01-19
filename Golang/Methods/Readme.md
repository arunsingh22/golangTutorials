##                          Functions 
-----------------------------------------------------------------

Func are "first class" objects; you can:
-   Define them - even inside another func
-   Create anonymous function literals
-   Pass them as func params / return values
-   Store them in variables
-   Store them in slices and maps(but not as keys)
-   Store them as field values in struct
-   Send and recievve them in channels
-   Write methods against a func type.
-   Compare a func var against nil 


##              Parameter Passing:
-----------------------------------------------------------------

### By Value:
-   numbers
-   bool
-   arrays
-   struts

### By reference:
-   strings (but they are immutable)
-   slices
-   maps
-   channels
-   things passed by pointer (&x)

**NOTE** The ultimate TRUTH:
- All params are passed by copying something (ie by value), if the thing copied is a ptr or descriptor, then the shared backing store (array,hashtable etc) can be changed through it, Thus we think of it as "by reference". 
