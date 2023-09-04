# Networking with HTTP

[Video](https://www.youtube.com/watch?v=Q-uy0FS6RwU&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=15)

Go std library has many pkgs for making web servers.
Including:
- Client & server sockets
- Route multiplexing
- HTTP and HTML, including HTML templates
- JSON and other data formats
- Cryptographic security
- SQL database access
- Compression utilities
- Image generation

### Go HTTP design:
- An HTTP handler function is an instance of an interface.
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

// Method declaration, on a function.
// Not just restricted to classes/structs.
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}

// The HTTP framework can call a method on a function type
// Member of the HandlerFunc type, structurally matches.
// HTTP package will be able to handler this.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world! from %s\n", r.UTL.Path[1:])
}
```
