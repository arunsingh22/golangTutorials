
<!-- https://refactoring.guru/design-patterns/singleton/go/example#example-1 -->
There are other methods of creating a singleton instance in Go:

**init function**
We can create a single instance inside the init function. This is only applicable if the early initialization of the instance is ok. **The init function is only called once per file in a package, so we can be sure that only a single instance will be created.**

sync.Once
The sync.Once will only perform the operation once.

## Drawbacks:
------------------------------
1. It violates the single reponsiblity principle as the class is doing more than 1 thing 
2. Special care is required to make it thread safe. 
3. Unit testing becomes difficult as mocking requires more than 1 instance.
