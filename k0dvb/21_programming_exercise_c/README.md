# Programming Exercise C

[Video](https://www.youtube.com/watch?v=YUaruvHkXio&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=22)

Exercise 7.11 from GOPL:web front-end for a database

Add additional handers to example gopl.io/ch7/http4 (or start from scratch) so that clients can create, read, update, and delete database entries.

For example, a request of the form: `/update?item=socks&price=6`
will update the price of an item in the inventory and report an error if the item does not exist or if the price is invalid.

We will *ignore* the race conditions for the purpose of this exercise, as there is a concurrency issue. Will be modified in later exercises.

Matt's solution: https://github.com/matt4biz/go-class-exer-7.11

- Start from scratch.
- Build a webserver for a store in OO style, with a method value handler.
- Have a DB with items and prices.
- Fill in method bodies for listing and updating prices, adding and removing items.

A method value closes over the receiver to become a plain function in terms of the regular params in it's param list once it's closed over the reciever.

If we write a method against some type tha t takes a ResponseWriter and a *Request, it can become a handler.

**Example usage:**
```shell
$ go run ./part1 &
[1] 14074

$ curl http://localhost:8080/list
shoes: $50.00
socks: $5.00

$ curl http://localhost:8080/read?item=socks
item socks has price $5.00

$ curl http://localhost:8080/update?item=socks\&price=6
new price $6.00 for socks

$ kill %1
[1]+  Terminated: 15          go run ./part1
```

### Refs:
https://gobyexample.com/url-parsing

### Interfaces in HTTP (revisited)
```go
// Go built-in webserver is looking for this.
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

// Type that simply wraps a function type.
type HandlerFunc func(ResponseWriter, *Request)

// Delcare method on HandlerFunc that satisfies the Handler interface.
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}

// Handler matches type HanderFunc and so interface Handler
// so the HTTP framework can call ServeHTTP on it.

func handler(w, http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world from %s\n", r.URL.Path[1:])
}
```