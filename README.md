# building-blocks-app
App created from building blocks

# Rationale
Brooks (1975) distinguishes two forms of complexity associated with software development. 
The essential complexity is the complexity that needs to be dealt with in order to write the necessary business logic.
The accidental complexity is the complexity that is caused by getting the code on production, getting certain systems to interact, etc. In other words, everything that does not directly contribute to the business logic, but is necessary to make the whole thing work. 

Oftentimes, I am faced with the situation where the majority of my time (and of my team's time) is spent on conquering the accidental complexity. 
Ideally, the accidental complexity would be minimal and the majority of the developer time would be spent on essential complexity, therefore providing direct business value. 

One of the examples where a lot of time is spent in accidental complexity is when spinning up a new (micro)service, batch job or other function. If 80% of the skeleton code would already be written, the developer would have more time to spend on encoding the actual business logic. Therefore, sometimes people copy an already existing piece of code (potentially another service/batch job) and then start removing everything that's not part of the intended functionality. 

The building-blocks idea provides the same mechanism, but then the other way around: 
A building block repo consists of pre-written code that does one thing and does it well. You can include one or multiple building blocks into a new (or existing) project and adapt it to your needs. 


## How to include a building block
1. Clone the building block to a directory:  
e.g. `git clone git@github.com:ggerritsen/postgresql-block repository`
2. Remove the .git directory and .gitignore file:  
`cd repository && rm -rf .git && rm .gitignore`
3. Update the package name of the building block files to reflect the new path  
e.g. `sed -i '' 's/package main/package <INSERT NEW NAME>/g' $(grep -Ril package *)`
4. Use the building block's code in the app and edit it where necessary


### Next steps
- properly vendor in dependencies (using go mod)
- deploy to heroku/AWS

## Handy commands
- Create a document via kafka:  
`./bin/kafka-console-producer.sh --broker-list localhost:9092 --topic test`  
`{"name":"testDocument","description":"test description"}`  
- Create a document via http:  
`curl -v -X POST -d "name=testDocument;description=testDescription" http://localhost:8081/documents`
- Retrieve a document via http:  
`curl http://localhost:8081/documents/1`

### To become a full-fledged app
- more tests
- graceful shutdown
- proper logging
- auto-recovery from errors 
- validation

### Sources
- [Hexagonal architecture](https://dzone.com/articles/hexagonal-architecture-what-is-it-and-how-does-it)
- [Fred Brooks (1975), The mythical man-month](https://en.wikipedia.org/wiki/The_Mythical_Man-Month)
