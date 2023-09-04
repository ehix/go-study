# Structs, Struct Tags & JSON

[Video](https://www.youtube.com/watch?v=0m6iFd9N_CY&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=13)

### Stucts:
A complex type, with fields and values, similar to a database record.
Fields are accessed with dot notation, can also be init with in the literal form.

Example:
```go
type Employee struct {
    Name    string
    Number  int
    Boss    *Employee      // types can ref to themselves via a pointer
    Hired   time.Time
}
```

Tip: if adding structs to a map, make a map of struct pointers not values.
Example:
```go
a := map[string]*Employee{}
...
a["Alex"] = &Employee{
    Name: "Alex"
    Number: 1
    Boss: a["someone"]
    Hired: time.Now()
}
```

Can be created anon, and literally:
```go
var album = struct {
    title   string
    artist  string
    year    int
    copies  int
}{
    "The White Album",
    "The Beatles",
    1968,
    10000000000000
}
```

### Struct compatibility:
Two `struct` types are compatible if:
1. the fields have the same types and names
2. they're in the same order
3. they have the same tags (*)

A `struct` may be copied or passed as a param in its entirety.
    - Pass as a copy and return a copy.
    - Pass a pointer and use reference symantics to modify.
A `struct` is comparable if all its fields are comparable, including anon structs.
The `zero` value for a `struct` is <i>"zero"</i> for each field in turn.


Structs with different names but same structure and behaviour (types and names) can be converted between one another.

### Struct Tags and JSON:
See `main.go`

### Struct tags have many uses:
Tags can also be used in conjunction with SQL queries
Example:
```go
import "github.com/jmoiron/sqlx"

type item struct {
    Name string `db:"name"`
    When string `db:"created"`
}

func putStats(db *sqlx.DB, item *item) error {
    stmt := `INSERT INTO items (name, created)
             VALUES (:name, :created);`
    _, err := db.NamedExec(stmt, item)
    return err
}
```