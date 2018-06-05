# G Search
## Introduction
This Project creates a web server (listening on port 8080)  with a`GET /search` endpoint that will return the most appropriate 20 items given `searchTerm`, `lat` (latitude) and `lng` (longitude).

 e.g. `/search?searchTerm=camera&lat=51.948&lng=0.172943` 



## Installation 

### Requirement 

`go version 1.8.1 and higher`

### Download 

#### Go Get

Run `go get github.com/ObsidianRock/backend-challenge` and go to the new directory and run `go install` 

This will put the executable in the directory `$GOPATH/bin ` ( this should be in your system environment variables)

From there onward, you can simple run the `backend-challenge.exe` and will spin up a server listening on `port 8080`

#### Alternatively 

Download the executable from the **release** section on the GitHub project 



##### Possible Errors from installation 

On `windows `  

```bash
# github.com/mattn/go-sqlite3
exec: "gcc": executable file not found in %PATH%gcc
```

The project uses `go-sqlite3` as SQLite driver. go-sqlite3 uses cgo hence need a C compiler. 

On windows, Install TDM-GCC. 



#### Testing 

go the directory `$GOPATH/src/ObsidianRock/backend-challenge` and run the command: 

`go test ./...` to run all tests in current directory and all of its subdirectories.



## Packages

There are four main packages in the project, 

### util 

The *Tokenizer* function takes a string input and returns a slice of discrete words from the input string, the function will also remove any non-alphanumeric characters and lower case the word:

For example 

sentence  = "Canon EF 24-70mm "

tokens = ["canon",  "ef", "2470mm"]

refer to package test for more examples. 

### prep

This package takes user input from URL query to create the SQL query. 

The query parameters is checked against a map of predefined search terms (**searchTerm**, **lat** and**lng**), anything else raises an error.  The search terms is **tokenized** and a SQL prepared statement is created dynamically based on the number of token words.  The prepared statement is used to prevent SQL injection. 

example of generated query: 

`SELECT * FROM items WHERE item_name LIKE ? or item_name LIKE ? or etc.. `

The SQL **LIKE** operator is used to search for a specified pattern in a column.  

In this case it will be `LIKE %token_word%`  to find any value that that have "token_word" in any position

See the package tests for examples. 

### weigh 

The weigh package will take item and give it an importance based on two factors; **distance** and **rank**. 

The **distance** in meters between the **item** and the search **lat** and **lng** is calculated using the  [haversine formula](https://en.wikipedia.org/wiki/Haversine_formula).  For this project, the open source project [tile38](https://github.com/tidwall/tile38/blob/7e9871bb69561d4ba3a63407a64a29b5cf7e1b14/pkg/geojson/geo/geo.go) has a good function which implements this formula and used in this project. The **rank** determines the relevance of an item based on the number of time an item name matches the tokenized search terms. 

Take a look at the package tests for examples. 

### query 

The query package will execute the SQL query with the parameters, due to the use of the SQL  `LIKE` parameter, the token words needs to be pre and post appended with the symbol `%` .  

The **distance** and **rank** of each item will also be calculated. 

### sorting 

To keep it simple, the result is sorted based on decreasing rank and increasing distance. 

See the package tests for examples. 



## Improvements 

### Pagination 

To keep it simple, the project omits pagination. To include pagination, a suggestion would to track some form of metadata in the form of offset and limit, so the output JSON will include those and next query will include the metadata to get the next two most prominent results. 

### Performance 

According to this [report](https://www.vividcortex.com/blog/2014/11/19/analyzing-prepared-statement-performance-with-vividcortex/), using prepared statement could lead to performance degrading due fact that each query is generating a new prepared statement. The improvement would be to remove the prepared statement as the *database/sql* query function creates prepares statements under the hood according to this [guide](http://go-database-sql.org/prepared.html). 

### Database/PostgreSQL

SQLite a fantastic database for testing a prototyping, however using a more robust database such a PostgreSQL for production would be better as well as the fact that it can be extended to include spatial database support with PostGIS is awesome. 

