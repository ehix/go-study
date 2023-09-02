# Arrays, Slices, and Maps

[Video](https://www.youtube.com/watch?v=T0Xymg0_aSU&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=6)

## Composite types:
- `array`: `[4]int`
- `slice`: `[]int` a variable length array
- `map`: `map[string]int` a hashmap or dict

## Arrays
- They're not used alot, due to the fact they have a fixed size.
- <b>Arrays are passed by value</b>, thus elements are copied.

```go
var a [3]int
var b [3]int{0, 0, 0}
var c [...]{0, 0, 0}   // sized by initializer
var d [3]int
d = b                   // elements copdies

var m [...]int{1, 2, 3, 4}
c = m                   // type mismatch, size 4 and 3
```

## Slice
- Similar to a string, has a descriptor and points to memory.
- Always has an array behind it.
- descriptor has length and a capacity.

```go
// declaration
var a []int         // nil, no storage
var b = []int{1, 2} // initialised

// modification
a = append(a, 1)    // append to nil OK
b = append(b, 3)    // []int{1, 2, 3}

// a takes on b's descriptor
a = b               // overwrites a

d := make([]int, 5) // []int{0, 0, 0, 0, 0}

// copy of a's descriptor, changes to a means changes to e.
e := a              // same storage (alias) 

// comparison
e[0] == b[0]        // true
```

- slices are indexed like `[8:11]`, from 8 to 11 (excluding 11).

## Slice vs Array:
- Most Go APIs take slices as inputs, not arrays.
- Can't `const` arrays.

|Slice|Array|
|---|---|
|Variable length|Fixed length at compile time|
|Passed by reference|Passed by value (copied)|
|Not comparable|Comparable (==) as has fixed length|
|Cannot be used as map key|Can be used as map key|
|Has `copy` & `append` methods|N/A|
|Useful as function parameters|Useful as "pseudo" constants (DES encyption)|

### Examples
```go
var w = [...]int{1, 2, 3}   // array of len 3
var x = []int{0, 0, 0}      // slice of len 3

func do(a [3]int, b []int) []int {
    a = b                   // SYNTAX ERROR, can't assign slice to array
    a[0] = 4                // w unchanged, modifying local copy only
    b[0] = 3                // x changed, x passed by reference

    c := male([]int, 5)     // []int{0, 0, 0, 0, 0,}
    c[4] = 42
    copy(c, b)              // copies only 3 elems

    return c
}

y := do(w, x)
// w unchanged, x changed, new slice y
fmt.Println(w, x, y)        // [1 2 3] [3 0 0] [3 0 0 0 42]
```
## Maps
- Can read from a nil map, but inserting will panic.
- Maps are passed by reference; no copying, but updating is OK.
- Values used for keys must have `==` and `!=` defined.
### Examples
```go
// declaration
var m map[string]int        // nil, no storage
p := make(map[string]int)   // non-nil, but empty (make creates memory)

a := p["the"]   // nothing there, so returns logical default = 0
b := m["the"]   // same thing, has no keys, default value again.
m["and"] = 1    // PANIC - nil map, there's no spot for the key "and".
m = p           // OK, same map as p now.
m["and"]++      // both p and m will get modified
c := p["and"]   // returns 1
```

```go
// map literal
var m = map[string]int{
    "and": 1,
    "the": 1, 
    "or" : 2,
}

var n map[string]int    // nil
b := m == n             // SYNTAX ERROR, can't compare them
c := n == nil           // true
d := len(m)             // 3
e := cap(m)             // TYPE MISMATCH, works on slices not maps
```
- Maps have a special two-result lookup function.
- The second variable tells you if they key was there.
```go
p := map[string]int{}   // literal map, non-nil but empty

a := p["the"]           // returns 0 (is it not in the map or is the value 0?)
b, ok := p["and"]       // 0, false (the key was actually missing)

p["the"]++

c, ok := p["the"]       // 1, true (the key was in the map)

if w, ok := p["the"]; ok {
    // we know w is not the default value
    // 'ok' confirms that "the" was in the map.
    ...
}
```
## Make `nil` useful:
- Many built-ins are safe: `len`, `cap`, `range`
```go
var s []int             // slice of int is nil, len 0
var m map[string]int    // 

l := len(s)             // len of a nil slice is 0

i, ok := m["int"]       // 0, false for any missing key (default value)

for _, v := range s {   // skip if s is nil or empty, 0 iterations
    ...
}
```