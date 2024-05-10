# Application Programming Interfaces

## What are they?

Application Programming Interfaces (API for short) are a means of communication between two different software systems, programs, and applications. 
It defines the methods and data formats that the two systems will use to communicate with each other.
It is essentially a way of defining how two computers, or more for that matter, will send resources and data back and forth to each other.

## Golang API Development

Golang has a built-in HTTP package that allows an ease of development when it comes to creating simple APIs.

This is done by using the net/http package which provides a lot of functionality for creating HTTP servers.

A quick example of how to create an HTTP server that returns hello world on the root route:

```go
package main

import (
    "net/http"
    "log"
)

func main() {
    mux := http.NewServeMux()
    
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World"))
    })
    
    server := &http.Server{
        Addr:   ":8080",
        Handler: mux,
    }
    
    log.Fatal(server.ListenAndServe())
}
```

## RESTful APIs

RESTful APIs are a type of API that follows the REST architectural style. 

An example of a RESTful API is the `std-api/` folder in this repository. This API is a simple RESTful API that allows you to create, read, update, and delete posts. 

The `crud-api/` folder in this repository is another example of a RESTful API that allows you to create, read, update, and delete users from a local MySQL database.
This API uses the `go-chi` router package to create the routes and handlers for the API. 

The differences between these two API's are the way they are structured and the way they handle requests. One uses the standard library, and the other one uses an idiomatic approach of creating routes and handlers (with slight differences in the database and querying).
