package main

//Writing a basic HTTP server is easy using the net/http package.

import (
    "fmt"
    "net/http"
)

//A fundamental concept in net/http servers is handlers.
//A handler is an object implementing the http.Handler interface. 
//A common way to write a handler is by using the http.
//HandlerFunc adapter on functions with the appropriate signature.

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}


//Functions serving as handlers take a http.ResponseWriter and a http.Request as arguments.
//The response writer is used to fill in the HTTP response.
//Here our simple response is just “hello\n”.
//This handler does something a little more sophisticated by 
//reading all the HTTP request headers and echoing them into the response body.



func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}


func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":8090", nil)
}
