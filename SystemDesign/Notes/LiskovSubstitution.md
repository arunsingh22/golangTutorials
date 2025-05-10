
# Liskov Substitution
---------------------------
The letter 'L' in the SOLID principles stands for the Liskov Substitution Principle (LSP).   

In simple terms, LSP means that if you have a base type (or interface), you should be able to use any of its subtypes (or implementations) without knowing it and without breaking your program's behavior.   

Imagine you have a function that works with a Shape interface, which has a method to calculate the area. If you have specific shapes like Rectangle and Circle that implement this Shape interface, the LSP states that you should be able to pass either a Rectangle or a Circle to your function, and it should work correctly without any unexpected issues. The function shouldn't need to know the specific type of shape it's dealing with.   

How much is LSP applicable to Golang?

The Liskov Substitution Principle is highly applicable to Golang, even though Go doesn't have traditional class-based inheritance like some other object-oriented languages. Go achieves polymorphism and code reuse through interfaces and composition, and LSP is a crucial concept in ensuring these mechanisms are used correctly.

Here's how LSP applies in Go:

Interfaces: Go's interfaces define contracts. Any type that implements all the methods of an interface satisfies that interface. LSP in Go means that if a type T implements an interface I, then any code that uses a variable of type I should be able to work correctly with a value of type T without needing to know the specific type T. The behavior of T should be consistent with the expectations set by the interface I.   
Composition: Go favors composition over inheritance. When you compose types (embed one struct within another), the embedded type's methods become available on the outer type. LSP in this context means that if a type B embeds a type A, and you have code that works with A, you should be able to use B in its place if the relevant methods of A are still accessible and behave as expected through B.   
Behavioral Consistency: The key to LSP is behavioral consistency. Subtypes or implementations should not introduce unexpected side effects or violate the contracts defined by their base types or interfaces. For example, if an interface defines a method to "get" a value, an implementation shouldn't "set" or delete that value as a side effect.
Method Signatures: When implementing an interface, the method signatures (names, parameters, and return types) must match the interface definition. Furthermore, the behavior implied by the method name and its purpose in the interface should be upheld by the implementing type.   
