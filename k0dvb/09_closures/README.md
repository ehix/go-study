# Closures
[Video](https://www.youtube.com/watch?v=US3TGA-Dpqo&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=10)

Functions that live inside functions, and refer to the enclosing function's data.

Closures are types of function call that have additional data from outside the function, from another functions scope, that's closed over by reference, similar to a parameter.

The gotcha is, bc we're closing over by reference, if the closure if executed async (often much later on), it's possible that the variable that's closed will mutate.

### Variable scope vs lifetime:
Scope is static, based on the code at compile time.
Lifetime depends on program execution (runtime).
    Go compiler does escape analysis to define lifetime, adding to the heap (interpretered languages use the heap all the time rather than the stack.)

### What is a `closure`?:
A `closure` is when a func inside another func "closes over" one or more local variables of the outer function.
- Think about it as being the same as descriptors.

```go
func fib() func() int { // returns a func that returns an int
    a, b := 0, 1        // a and b declared here
    return func() int { // anon function
        a, b = b, a + b // uses a and b, didn't declare them
        return b        // the closure is what's being returned
    }
}

f := fib    // f is a function pointer
f := fib()  // f is a closure (has pointer to anon func, and env where &a and &b are referenced)
```

