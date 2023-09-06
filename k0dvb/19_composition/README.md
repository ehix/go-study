# Composition, not Inheritance

[Video](https://www.youtube.com/watch?v=0X6AcnwocbM&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=20)

The fields and methods of an **embedded** `struct` are **promoted** to the level of the embedding structure.
```go
type Pair struct {
    Path string
    Hash string
}

// Pair is embedded in PairWithLength.
// Path and Hash are "promoted to PairWithLength".
type PairWithLength struct {
    Pair
    Length int
}

pwl := PairWithLength{Pair{"/usr", "0xfdfe"}, 121}
// Not.pwl.x.Path, because Path and Length are at the same 
// level in PairWithLength as they are in Pair.
fmt.Println(pwl.Path, pwl.Length)
```
Any type can be embedded in another type, not only structs.

A `struct` can embed a pointer to another type; promotion of its fields and methods works the same way. 

### Sorting

`sort.Interface` is defined as ...
```go
type Interface interface {
    // # elements in the collection
    Len() int
    // reports whether the element with index i
    // should sort before the lement with index j
    Less(i, j int) bool
    // swaps the elements with indices i and j
    Swap(i, j int)
}

func Sort(data Inferface) // <- takes interface
    ...

entries := []string{"baker", "dog", "able"}
sort.Sort(sort.StringSlice(entries)) // <- type cast into StringSlice
fmt.Println(entries) // [able baker dog]
```

`sort.Reverse` which is defined as ...
```go
// The methods of Interface are promoted into Reverse.
type reverse struct {
    // This embedded Interface permits Reverse to use the 
    // methods of another interface implementation.
    Interface
}

// Reverse redefines Less, to return the opposite of the embedded implementation's Less method.
func (r reverse) Less(i, j int) bool {
    return r.Interface.Less(j, i)
}

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
    return &reverse{data}
}

entries := []string{"baker", "dog", "able"}
sort.Sort(sort.StringSlice(entries)) // <- type cast into StringSlice
fmt.Println(entries) // [dog baker able]
```

### Make nil useful

```go
// FIFO
type StringStack struct {
	data []string   // "zero" value ready-to-use
    // data is lowercase/encap/not exported.
}

// Push and Pop are exported.
func (s *StringStack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StringStack) Pop() string {
    // Check len as can't pop of an empty stack.
    // This isn't needed but is added to control the panic message.
	if l := len(s.data); l > 0 {
		t := s.data[l-1]        // Tail
		s.data = s.data[:l-1]   // Head
		return t
	}
	panic("pop from empty stack")
}
```
#### Nothing prevents calling a method with a nil receiver
```go
// Linked list through pointers
type IntList struct {
	Value int
	Tail  *IntList
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil { // The list is done when the Tail pointer is nil.
		return 0
	}
    // Can recursively walk through list.
	return list.Value + list.Tail.Sum()
}
```
