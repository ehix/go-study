# Basic Types

## Keywords and symbols:
- 25 keywords:
```go
break       default        func     interface   select
case        defer          go       map         struct
chan        else           goto     package     switch
const       fallthrough    if       range       type
continue    for            import   return      var
```

- Wont list all operators...

## Predeclared identifiers:
- Constants:
```go
true    false   iota    nil
```
- Types:
```go
int         int8    int16       int32       int64
uint        uint8   uint16      uint32      uint64  uintptr
float32     float64 complex64   complex128
bool        byte    rune        string      error
```

- Functions:
```go
make  len  cap  new  append  copy  close  delete
complex  real  imag
panic  recover
```

## Types

- Machine native
    ```a := 2``` name/address of memory location
    no interpreter
    makes it performant
    source -> compiler; no abstraction.

- Integers
    ```a :=2 ``` will be treated as `int`, 64 bit.
- Real: `float` no generic, 64 bit.
    Not money, but scientific.
- Complex: imaginary `float`.

- Declarations:
```go
var a int

var (
    b = 2 // assume int
    f = 2.01 // assume float64
)

// inside functions:
// short declaration operator
c := 2
```
- Show types and values:
```go
//  8  -> format
// [1] -> reuse x
fmt.Printf("%8T %[1]v\n", x)
```

- Simple types:
    - bool can be false or true, cannot convert to/from int
    - error has one function Error()
        an error can be nil or non-nil
    - pointers are physical addresses, logically opaque.
        a pointer may be nil or non-nil

- No uninitalisated variables in `go`, will have a "zero" value

- Constants are immutable to be safe for concurrency.


