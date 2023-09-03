# Functions, Parameters and Defer
[Video](https://www.youtube.com/watch?v=wj0hUjRHkPs&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=9)

### Functions
- Are first class objects, you can do anything with it.
- Almost anything can be define inside a function.
- Functions have signature, order & type of it's params, and return type.

```go
var try func(string, int) string
func Do(a string, b int) string {
    ...
}
```

### Param terms:
Params are distinguished via two terms:
- A function declaration lists <b>formal</b> parameters.
- A function call has <b>actual</b> parameters, i.e. arguments specific to the function call.

Two param passing models:
- A param is passed by value if the function gets a copy.
    - numbers, bool, arrays, structs
- A param is passed by reference if the function can modify it.
    - pointer, strings, slices, maps, channels.

- From a techinical view, <b>every param is passed by value</b>..
    1. bc, the formal param is a local var in the function.
        if it's metally 'passed by reference' it's still copying the descriptor to underlying data.
    2. tf, it gets assigned/copied to the value of the actual param.
    3. so, if it's a pointer/descriptor, then shared

### Return values:
Have to peren all return values `(x, y, error, ...)`

### Recursion:
Anything done recurisively can be done with a for loop, but with book keeping. In recursion, the callstack/stack-frames does the book keeping.
Recursion is ultimately slower, bc it needs to make all the stack-frames.
- Think tree and graph traversal.
- Protect against infinite recursion.

### `defer`:
The defered function call is guaranteed to fun at function exit FIFO.
Examples:
1. close a file we opened
2. close a socket/HTTP request we made
3. unlock a mutex
4. make sure something gets saved before we're done..

The `defer` statement captures a function call to run later.
It operates at function scope, not block scope, when the function exits. 
- e.g. if in a `for` loop in a function, the defered function won't be triggered until the whole loop is executed at the end of the function!

```go
func main() {
    f, err := os.Open("my_file.txt")
    if err != nil {
        ...
    }
    defer f.Close()
    ... // do something with the file.
}
```

Unlike a `closure`, `defer` copies arguments to the deffered call.
Example:
```go
func main() {
    a := 10
    defer fmt.Println(a) // a gets copied here..
    // <!> it's not a reference 

    a = 11
    fmt.Println(a)
} // >> 11, 10
```

### `defer` gotcha & 'naked return':
- Params have names, but can also give names to return values.
- Named return values act as local variables to the function, so can be assigned to within the function body.
- This allows for a naked return, as the return value is named.
```go
func do() (a int) { // named return value
    defer func() {
        a = 2
    }()             // is called when naked return is hit below
    a = 1
    return          // naked return
} // returns 2
```
