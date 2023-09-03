# Control Statements; Declarations & Types

[Video](https://www.youtube.com/watch?v=qpHLhmoV3BY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=8)

### Conditionals
- if-then-else can start with a short declaration or statement, like the standard error catching block.
```go
if err := doSomething(); err != nil {
    return err
}
```

### Loops and labels
- Theres no while loop.
- Maps are unordered.

- Can add a label to a loop:
```go
outer:
    for k := range testItemsMap {
        for _, v := range returnedData {
            if k == v.ID {
                continue outer // this label refers to the outer loop
            }
        }
        t.Errorf("key not found: %s", k)
    }
```

### Switch
- go has `switch`, cases break automatically:
```go
switch x := y.Get() x {
    case 0, 1, 2:
        ...
    case 3, 4, 5:
        ...
    default:
        ...
}
```
- another switch pattern, on logical comparison:
```go
switch {
    case x < y:
        ...
    case x > y:
        ...
    default:
        ...
}
```

### Packages
- Everything lives in a package, every source file starts with a package declaration.
- Therefore, two scope, package scope and function scope.
    ~ remember `:=` can only be used inside a function.
    ~ package scope variables are often declared with keywords `const`, `var`, `type`, etc.

- Packages aids encapsulation.
- <b>Every name that's capitalised is exported</b>, otherwise private to package.
- Generally, files of the same package live in the same directory.
- Can't have cyclic dependencies, so move common deps to a third package, or eliminate them.
- A package should embed deep functionality behind a simple API.

### Initialisation:
- Items within a package get init before `main` starts.
- Only the runtime can call `init()` implicitly, also before main.
    ~ typically used for drivers and plug-ins, e.g SQL.

### Named typing:
- Go used named typing for non-functional user-declared types.
- Vars might have the same underlying type, but as it's explicitly typed, it needs to be type-converted.
```go
type x int

func main() {
    var a x     // x is a defined type; base int
    
    b := 12     // b defaults to int
    
    a = b       // TYPE MISMATCH
    
    a = 12      // OK, untyped literal
    a = x(b)    // OK, type conversion
}
```
