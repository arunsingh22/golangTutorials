## Interfaces in Golang
<!-- Read for better understanding -->
<!-- https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c -->
----------------------------------------------------------------------------------
In Go we’ve two concepts related to interfaces:

- **Interface — set of methods** required to implement such interface. It’s defined using keyword interface.
example: 
```type I interface {
    m1()
    m2(int)
    m3(int) int
    m4() int
}
```
- Interface type **(interface{})**— variable of interface type which can hold any value implementing particular interface
Let’s discuss these topics in next two se
- **interface{}** is base type for all TYPES in golang therefore every type in go implicitly inherits interface{} type.
    ```
    fmt.Println(a ...interface{})
    ```

2. Because interface is also a type , we can create variables for it and make calls to functions BUT this will give us runtime Exception.
``` No compile-time error BUT Runtime error
	panic: runtime error: invalid memory address or nil pointer dereference
	var x area
	fmt.Println(x.Area(5))
```
3. **In Go, an interface cannot implement other interfaces or extend them, but we can create a new interface by merging two or more interfaces.Therefore only composition of interface is allowed.**