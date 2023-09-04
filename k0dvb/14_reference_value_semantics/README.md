# Reference & Value Semantics

When to use pointers and values, and what is the trade-off?

Pointers: shared and not copied
Values: copied and not shared

Value semantics lead to higher integrity, particularly with concurrency; don't share.

Pointer semantics <i>may</i> be more efficient

### Common uses of pointers:

- some objects can't be copied safely (mutex)
- some objects are too large to copy efficiently
- some methods need to change (mutate) the receiver (later)
- when decoding protocol data into an object, e.g. json unmarshal
- when using a pointer to signal a `null` object, e.g. tree structures when child nodes have no children.

### Must not copy:
- Any `struct` with a mutex must be passed by reference.
```go
type Employee struct {
    mu sync.Mutex
    Name string
    ...
}
func do(emp *Employee) {
    emp.mu.Lock()
    defer emp.mu.Unlock()
}
```

### May copy:
- Any small `struct` under 64 bytes should be copied.
- Go routinely copies string and slice descriptors.
```go
type Widget struct {
    ID      int
    Count   int
}

func Expend(w Widget) Widget {
    w.Count--
    return w
}
```

### Stack allocation:
- Stack allocation is more efficient.
- Accessing a variable directly is more efficient than following a pointer.
- Accessing a dense seq of data is more efficient that sparse data (an array is faster than a linked list, etc).

### Heap allocation:
- Go would prefer to allocate on the stack, but sometimes can't:
1. a function returns a pointer to a local obj
2. a local obj is captured in the function closure
3. a pointer to a local obj is sent via a channel
4. any obj is assigned into an interface
5. any obj whose size is variable at runtime (`slice`).

Tip: build with the flag `-gcflags -m=2` to see escape analysis.

### For loops:
```go
for i, thing := range things {
    // thing is a copy
}

for i := range things {
    // things is mutated
    things[i].which = whatever
}
```
### Slice safety:
Anytime a function mutates a slice that's passed, we must return a copy.
- Slice's backing array may be reallocated to grow.
```go
func update(things []thing) []things {
    ...
    things = append(things, x)
    return things
}
```

Keeping a pointer to an element of a slice is risky.
```go
type user struct { name string; count int } 

func addTo(u *user) {
    u.count++ 
}

func main() {
    users := []user{{"alice", 0}, {"bob", 0}} // len 2, cap 2
    alice := &users[0]          // risky, pointer to first slice ele
    amy := user{"amy", 1}       
    users = append(users, amy)  // new user added, users reallocates
    addTo(alice)                // alice is likely a stale pointer
    fmt.Println(users)          // alice's count will be 0
}
```

