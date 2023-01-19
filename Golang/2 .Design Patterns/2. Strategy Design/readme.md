
## Defination:
- “The Strategy Pattern defines a family of algorithms,it encapsulates each one of them , and makes them interchangeable, which means we can plug and play with any algorithm without changing the client code.
Strategy lets the algorithm vary independently from clients that use it.”

- **Strategy is a concept, a useful recipe for solving a particular, recurring problem. It's not a language construct, nor is it about any one form of implementation. A closure can be used to implement Strategy one day and Observer the next day.**

- **State can be considered as an extension of Strategy. Both patterns are based on composition**: they change the behavior of the context by delegating some work to helper objects. Strategy makes these objects completely independent and unaware of each other. However, State doesn’t restrict dependencies between concrete states, letting them alter the state of the context at will.