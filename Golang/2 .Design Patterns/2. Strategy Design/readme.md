

- **Strategy is a concept, a useful recipe for solving a particular, recurring problem. It's not a language construct, nor is it about any one form of implementation. A closure can be used to implement Strategy one day and Observer the next day.**

- The term Strategy is mostly useful in conversations with other programmers to concisely express your intent. There's nothing magical about it.

## Defination:
- “The Strategy Pattern defines a family of algorithms,encapsulates each one, and makes them interchangeable.Strategy lets the algorithm vary independently from clients that use it.”

- **State can be considered as an extension of Strategy. Both patterns are based on composition**: they change the behavior of the context by delegating some work to helper objects. Strategy makes these objects completely independent and unaware of each other. However, State doesn’t restrict dependencies between concrete states, letting them alter the state of the context at will.