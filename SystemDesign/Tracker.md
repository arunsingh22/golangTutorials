
# LLD Topics          Status
-------------------------------
SOLID/KISS/DRY       [DONE]
Factory Pattern      [DONE]
Singleton Pattern    [DONE]
Strategy Pattern     [DONE]
Observer Pattern     [DONE]
Decorator Pattern    [...]
--------------------------------
Rate Limiter         [...] 
[https://medium.com/@choudharys710/lld-machine-coding-w-implementation-rate-limiter-34f87e74120f]
ParkingLot           []
TinyURL              []
ElevatorProblem      [] 
--------------------------------
SplitWiseApp         []
BookMyShow           []
VendingMachine       []
LoggingSystem        []
Meeting Scheduler    []
Cache                []
ATMSystem            []
CarRentalSystem      []
Tic-tac-Toe          []
Snake&Ladder         []
Chess                []
InventoryMgmtSystem  []


[https://codewitharyan.com/system-design/low-level-design]

- KISS : Keep it simple stupid - 
Over-complication can happen in 2 places 
1. Design complication this leads to scalability and maintainbility and can lead to 
   tight coupling.
2. Implementaion complication can lead to tigh coupling and poor readablity and            maintainbility.
3. **AVOID OVER-ENGINEERING** : Premature Abstraction and premature optimizaiton



Q: How many connections can a single Web/API Server can handle/accept 
   - What about servers like Nginx how much they can accept


# LLD
[https://medium.com/better-programming/how-to-ace-the-low-level-design-interview-3f1be6401070]

## Requirement gathering
   - Every requirement will translate into action in the system. Let’s take an example of designing a Social Network. One of the requirements would be to “Add a person as a friend”. In the software world, this is introducing a method AddFriend in one of the classes.
   - **Scope the requirements so that you would be able to finish the implementation in the given duration.**
   - Make a list of all the features that you have identified during this step.

# Actors and use cases
   - Once the requirements are clear, you should identify the actors in the system.
     Next, you note down the behavior and use cases of each actor.
   - For eg:- In the case of a chess game, there would be two players. Each of the players would take turns. One of them would be assigned black pieces and the other would get white ones.
   - In this step, you will list down the objects involved in the system. Find the relationships between them. This step will help you to come up with the correct data structure for organizing the data.

## NOTE: 
   The interview doesn't expect any kind of seq diagrams, activity diagrams etc.
   but you can draw a rough diagram for your own understanding.
