# REST(Representational State Transfer Protocol) API
- It contains all the information(Method, Header, payload) needed to fulfil the request.
- No any session info stored in server before and after the request
- Even server don't know that from which to whom we are communicating

### net/http :
**Core Concepts of net/http**
1. **Request** (*http.Request): Represents an incoming HTTP request received by a server or an outgoing HTTP request sent by a client. It contains details like:

* Method (GET, POST, PUT, DELETE, etc.)
* URL (path, query parameters)
* Header (key-value pairs like Content-Type, User-Agent)
* Body (the request payload)
* Host, RemoteAddr, etc.

2. Response (http.ResponseWriter): An interface used by an HTTP server handler to construct an HTTP response. You write status codes,    headers, and the response body to it.

3. Client (*http.Client): Used to make outgoing HTTP requests to other web servers.

4. Server (*http.Server): Used to listen for incoming HTTP requests and dispatch them to appropriate handlers.

5. Handler (http.Handler interface): A key concept. An http.Handler is anything that implements the ServeHTTP(ResponseWriter, *Request) method. This method is where you write your logic to process an HTTP request and send a response.

--- 
### A. Building a Web Server
This is typically done using `http.HandleFunc` and `http.ListenAndServe`.

1. `http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))`
**What it does:** Registers a function (your "handler function") to handle HTTP requests for a specific URL path pattern. It's a convenience wrapper for http.Handle that allows you to use a regular function instead of needing to define a struct that implements http.Handler.
**Scenario:** Defining routes and their corresponding logic for your web application.

2. `http.ListenAndServe(addr string, handler http.Handler)`
**What it does:** Starts an HTTP server that listens on the specified addr (e.g., ":8080" for all interfaces on port 8080).
- If handler is nil, it uses http.DefaultServeMux (the default multiplexer where http.HandleFunc registers its routes).
- It blocks until the server stops (e.g., by an interrupt signal or an error).
**Scenario:** Launching your web server.

