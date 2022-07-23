## Interfaces in Golang
----------------------------------------------------------------------------------

1. **interface{}** is base type for all TYPES in golang therefore every type in go implicitly inherits interface{} type.
    ```
    fmt.Println(a ...interface{})
    ```

2. Because interface is also a type , we can create variables for it and make calls to functions BUT this will give us runtime Exception.
``` No compile-time error BUT Runtime error
	panic: runtime error: invalid memory address or nil pointer dereference
	var x area
	fmt.Println(x.Area(5))
```
3. **In Go, an interface cannot implement other interfaces or extend them, but we can create a new interface by merging two or more interfaces.**