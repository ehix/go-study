# Formatted & File I/O

[Video](https://www.youtube.com/watch?v=dqEtGT-dxoY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=8)

### Standard I/O:
- Unix has the notion of three standard I/O streams.
- They're open by default in every program.
    Standard input, output, error.
- These are normally mapped to the console but can be redirected.

```go
    fmt.Println("standard output")
    fmt.Fprintln(os.Stderr, "standard error")
```

### `fmt`:
```go
// always os.Stout:
fmt.Println(...interface{}) (int, error)
fmt.Printf(string, ...interface{}) (int, error)

// print to anything that has the correct Write() method
fmt.Fprintln(io.Writer, ...interface{}) (int, error)
fmt.Fprintf(io.Writer, string, ...interface{} (int, error))

// return a string
fmt.Sprintln(...interface{}) string
fmt.Sprintf(string, ...interface{}) string
```

See [docs for all format codes](https://pkg.go.dev/fmt)

### File I/O:
- `os`: open, create files, list dir, etc; Has `File` type.
- `io`: read/write; `bufio` buffered i/o scanners.
- `io/ioutil`: read/write entire file to memory.
- `strconv`: convert to/from strings.

