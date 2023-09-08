# shop-discount
A sample project for implementing discount feature in an online shop 

## Discount Methodolody
Two main ways for implementing discounting features were considered:
1. **Rule Based**: By expressing disocunt policies in some DSL and building a knowledge base with neccessary facts. The rules can described in Drools DSL and used by inference engine of [grule](https://github.com/hyperjumptech/grule-rule-engine). Other option which seemed more appropriate was describing rules in Prolog and using [ichiban/prolog](https://github.com/ichiban/prolog) for inference engine.

2. **Adhoc Implementation**: By writing the code of each policy by hand. When the number of policies are small, this approach is drastically simpler compared to the previous one. By using this approach one can write tests for the policies in golang itself.

Because the purpose of this excersie was writing clean and maintanable code, second approach was chosen. However, the interface of disocunt service was designed in such a way that a rule based implementation could easily replace the adhoc one.

Each policy can be easily added to the sytem by implementing the `AdHocDiscount` interface.
## Architecture
The overall architecture of the shop can be anything from a monolithic one to a microservice one.
The overall structure of the code allows switching between either option seamlessly.
Services other than discount (e.g. order, user, product, cart and ...) can be implemented in the same app (monolithic) or just be a consumer of other microservices through something like gRPC or REST.

## Code Structure
All of the codes related to a single feature is placed in the package itself rather than seprating it by technical concerns to improve readability of the code base.
Each bussiness feature package is divided into 4 separate layers which are:
- Domain
- Service
- Repository
- Controller
 
Packages can only talk to each other through service layer.
There are files corresponding to the layers. The implementation of these layers, are located in a special file (e.g. adhock discount).

Some pacakges which do not correspond to bussiness features (internal services), are not following the above 4-layering rule to keep the code simple. 

## Tests
Some test cases were added as simple test cases to demonstrate unit testing and integration testing.
The test cases are by no mean exhaustive.
 
Current code coverage for discount package is: **50%**
## Run
Currently running the application is only possible through test, because much of the necessary codes are just dummy implementations.

## Time Spent
**Total** ~ 20h
- Studying alternative approaches ~ 3h
- Implementing adhoc approach ~ 12h
- Writing test: ~ 4h
- Writing documentation and comments on exported IDs ~ 1h

## Some Random Questions
-  *Why no dependency injection framework was chosen?*
  I personally belive Golang is simple enough to let one discard using a dependecy injection framework altogether and do the task of injection dependecies yourself as it is done in the code base.

- *Why is there a lock manager?*
 Because there were no assumption on the underlying persistence layer and the capabilities that it provides, a lock store was provided to prevent some problems (e.g. double spending a discount). This can be replaced with the underlying storage's transactional features such as serializable isolation level.
