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
#### Nothing prevents calling a method with a nil receiver

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
For example, a linked list type.
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

### Linked Lists
[Link](https://www.prepbytes.com/blog/linked-list/advantage-and-disadvantage-of-linked-list-over-array/#:~:text=In%20short%2C%20there%20are%20several,efficient%20sorting%20in%20some%20cases.)

The linked list can be used to implement stacks and queues that grow dynamically. The linked list is also used in implementing graphs in which the adjacent vertices are stored in the nodes of the linked list popularly known as Adjacency list representation.

### Advantages of a Linked List over Arrays
Below are some advantages of Linked List over Arrays in programming, including:

1. Dynamic Size: One of the most significant advantages of linked list over arrays is that linked lists can grow or shrink dynamically during runtime. This means that the size of a linked list can be adjusted to accommodate new elements or remove existing elements without having to allocate or deallocate a fixed-size block of memory, as is the case with arrays.
2. Efficient Insertion and Deletion: Linked lists allow efficient insertion and deletion of elements at any position in the list, whereas arrays require shifting of elements when a new element is added or removed, which can be slow and inefficient for large arrays.
3. Memory Efficiency: Linked lists use memory more efficiently than arrays. In an array, all elements must be stored in contiguous memory locations, even if some of the elements are not used. In contrast, linked lists only allocate memory for the elements that are used, which can save memory in cases where the size of the data set is unknown or varies over time.
4. Easy Implementation of Abstract Data Types: Linked lists are easy to use and implement when implementing abstract data types such as stacks, queues, and trees. These data structures require frequent insertion and deletion of elements, which is a task in which linked lists are well-suited.
5. More Efficient Sorting: In some cases, linked lists can be more efficient for sorting algorithms than arrays. This is because linked lists do not require swapping elements like arrays, which can be time-consuming for large arrays.

In short, there are several advantages of linked list over arrays, such as dynamic size, efficient insertion and deletion, memory efficiency, easy implementation of abstract data types, and more efficient sorting in some cases.

### Disadvantages of a Linked List over Arrays
While there are several advantages of linked list over arrays, they also have some disadvantages that need to be considered. Here are some disadvantages of linked list over arrays:

1. Random Access: Linked lists do not provide random access to elements like arrays do. To access an element in a linked list, we have to start at the beginning of the list and traverse the list until we find the desired element. This makes accessing individual elements in a linked list slower than in an array.
2. Extra Memory Usage: Linked lists require extra memory compared to arrays. Each element in a linked list requires a reference to the next element, which takes up additional memory space. In contrast, arrays only need memory to store the elements themselves.
3. More Complex Implementation: Implementing a linked list is more complex than implementing an array because it requires managing pointers and dynamically allocating memory. This complexity can lead to more bugs and errors in the code.
