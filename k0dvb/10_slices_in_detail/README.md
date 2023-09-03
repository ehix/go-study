# Slices in detail
[Video](https://www.youtube.com/watch?v=pHl9r3B2DFI&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=11)

### Empty vs `nil` slice:
- Example, nil vs empty slice (works the same for map).
- Take away, `if len(x) == 0` is prefered.
```go
package main

import {
    "encoding/json"
    "fmt"
}

func main() {
    var a []int
    
    j1, _ := json.Marshal(a)
    fmt.Println(string(j1))  // null

    b := []int{}
    
    j2, _ := json.Marshal(b)
    fmt.Printlin(string(j2)) // []
}
```

### Slice op:
- Slices are alias' to some underlying array.
- If using the [x:y] slice op, resulting slice takes on the cap of the original.
- Using [x:y:z] slice op, will control the len and cap of the resulting slice.