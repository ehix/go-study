# Methods and Interfaces

[Video](https://www.youtube.com/watch?v=W3ZWbhQF6wg&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=18)

An `interface` specifies abstract behaviour in terms of **methods** (method set).

`Stringer` is satisfied by any object that has a `String()` method.
```go
type Stringer interface { // in fmt
    String() string
}
```
*Concrete* types offer methods that satify the interface.

### Methods are type-bound functions

A **method** is special type of function ( from Oberon-2).

Can put methods onto any user-bound type.

It has a **receiver** parameter before the function name param. This can be compared to `self` in python.


```go
type IntSlice []int

func (is IntSlice) String() string {
    var strs []string

    for _, v := range is {
        strs = append(strs, strconv.Itoa(v))
    }

    return "[" + strings.Join(strs, ";") + "]"
}
```
- has-a -> composition
- is-a  -> substitutability

### Why interfaces?

An interface specifies required behaviour by a **method set**, one or more methods.

Any type that implements that method set satisfies the interface. You won't see an `implements` key word. No type will declate itself to implement something explicity.

This is known as **structual typing** or "duck" typing.

Think about interfaces from the consumer-side not the provider-side. For example, a function needs a param to provide some behaviour, the function (consumer) want's x interface, with y method set to provide the behaviour.

Without, we'd have to write many functions for many concrete types, possibly coupled to them.

Example without interfaces, and therefore no way to abstract behaviours:
```go
// File, Buffer, Socket would need to be concrete types.
func OutputToFile(f *File, ...) {...}
func OutputToBuffer(b *Buffer, ...) {...}
func OutputToSocket(s *Socket, ...) {...}
```
But instead we can do:
```go
// In io package.
// LCD is the byte type, it is concrete.
type Writer interface {
    Write([]byte) (int, error)
}

func OutputTo(w io.Writer, ...){...}
```

### Not just structs

A method may be defined on any **user-declared** (named) type (with some limitations).

That means methods can't be declared on `int` bc it's a **built-in** type, but...
```go
type MyInt int
func (i Int) String() string {
    ...
}
```
The same method name may be bound to different types.

### Receivers

A method may take a *pointer* or *value* receiver, but not both.

```go
type Point struct {
    X, Y float64
}

func (p Point) Offset(x, y float64) Point {
    // this method gets a copy, p is not changed.
    return Point{p.x+x, p.y+y}
}

func (p *Point) Move(x, y float64) {
    // this mothod gets an address, p is changed.
    p.x += x
    p.y += y
}
```
Taking a pointer allows the method to change the receiver (original object).

### Interfaces and substitution

All the methods must be present to satisfy the interface, therefore, it's good to keep interfaces small.

```go
// Interfaces:
var w       io.Writer          // has 1 methods, Write.
var rwc     io.ReadWriteCloser // has 3 methods Read, Write, Close.

w = os.Stdout           // OK: *os.File has Write method
w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
w = time.Second         // ERROR: no Write method

rwc = os.Stdout         // OK: *os.File has all 3 methods
rwc = new(bytes.Buffer) // ERROR: no Close method

w = rwc                 // OK: io.ReadWriteCloser has Write, is a broader interface.
rwc = w                 // ERROR: io.Writer has no Close method, is a narrower interface.
```

### Interface satisfiability

The **receiver** must be of the right type (pointer or value).
```go
type IntSet struct { /* ... */ }

func (*IntSet) String() string
// Can't assign String method below
var _ - IntSet{}.String()   // ERROR: String needs *IntSet (got literal, which doesn't have an address/pointer)

// Can assign String method here, as assigned variables have addresses/pointers required by the interface.
var s IntSet
var _ = s.String()          // OK: s is a variable; &s used automatically

// Define String method to work on a pointer receiver
var _ fmt.Stringer = &s     // OK
var _ fmt.Stringer = s      // ERROR: no String method
```

### Composition

`io.ReadWriter` is actually defined by Go as two interfaces:
```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// Can define another interface with composition of other interfaces:
type ReadWriter interface {
    Reader
    Writer
}
```
Keep interfaces small, one method, then compose interfaces from those if we need more complex behaviours.

### Interface declarations

#### All methods for a given type must be declared in the same package.

Static typing means we want to know all the methods a type has at compile time.

- We **can't extend** a type by adding methods in a different pkg
- We **can extend** the **type** in a new pkg through embedding:
```go
type Bigger struct {
    // Bigger is-not a Big
    // Bigger has-a Big
    my.Big                  // get all Big methods via promotion
}

func (b Bigger) DoIt() {
    ...                     // and add another method here
}
```

