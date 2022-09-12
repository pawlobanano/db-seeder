## DB Seeder for PostgreSQL

It works on PostgreSQL and utilizes go:embed feature. 
It goes through each of the _*.sql_ file located inside the *data* directory (it reads in ascending order - hence the naming convention with number as prefix). 
It runs the query in a transaction, so in the case data is already there - it will simply run SQL roll back feature. 

Finally it can be conveniently run via make file. 
Just needs the command 
```go run main.go```
