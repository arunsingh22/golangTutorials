
# LLD Topics          Status
-------------------------------
SOLID/KISS/DRY       [DONE]
Factory Pattern      [DONE]
Singleton Pattern    [DONE]
Strategy Pattern     [DONE]
Observer Pattern     [DONE]
Decorator Pattern    [...]
--------------------------------
Rate Limiter         [...] [https://medium.com/@choudharys710/lld-machine-coding-w-implementation-rate-limiter-34f87e74120f]
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

# Now in any object-oriented design interview, you interviewer is typically looking for three things:

## 1. How you list down requirements, especially core features?

e.g. If your problem statement is “Design a Parking Lot” then your core features will be park() and unpark() methods

if your problem statement is “Design a restaurant food order and rating system like zomato, swiggy, uber eats etc” then your core features will be

orderFood()
rateOrder()
display list of restaurants based on their rating or popularity

## 2. How you break your problem statement in multiple classes

I always find it easier to start listing entities and their corresponding entity managers(if required) first. e.g. For restaurant food ordering and rating system your entities can be Restaurant, order, FoodItem etc and their corresponding managers will be RestaurantsManager, OrdersManager etc.

## 3. How you use design patterns to solve the core features
The most common design patterns that you will come across in a low level design interview are Strategy, Factory, Singleton and Observer. You should be familiar with their implementation and different use cases where they can be used. We will see some of those use cases in a moment.

## 4. A fourth topic is also discussed if you have done well in above three steps.

Handling multi-threading. There will be discussion on use of locks, synchronization features and thread safe data structures for your design to work correctly in a multi-threaded environment.