# Strings

[Video](https://www.youtube.com/watch?v=nxWqANttAdA&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=5)

- Strings have two natures, logical and physical.
- Strings are all unicode, to help represent international characters.

- `byte`: `uint8`
- `rune`: `int32` (is the go equiv of a char, or wide char)
- `string`: an immutable sequence of `chars`
    physically a sequence of bytes (`utf-encoding`)
    logically a sequence of (unicode) `runes`

### Example
- The length of the string, is the length of the byte string that encodes it in utf-8.

```go
	s := "København"
	fmt.Printf("type: %8T value: %[1]v len: %d\n", s, len(s))
	a := []rune(s)
	fmt.Printf("type: %8T value: %[1]v len: %d\n", a, len(a))
	b := []byte(s)
	fmt.Printf("type: %8T value: %[1]v len: %d\n", b, len(b))
	// See, if ascii is 0-127, 'ø' is beyond range.
	// See, the 'ø' is represented by two bytes.
	// > type:   string value: København len: 10
	// > type:  []int32 value: [75 248 98 101 110 104 97 118 110] len: 9
	// > type:  []uint8 value: [75 195 184 98 101 110 104 97 118 110] len: 10
```

## In memory:

- Strings are <b>passed by reference</b>, so aren't copied.

```go 
s := "hello, world"
```
- `s` is like a "string descriptor", e.g a pointer with additional information, such as, the number of bytes it takes to make it. The length of the string is encoded in the descriptor.

```go
hello := s[:5]
world := s[7:]
```

- `hello` points to the same chars in memory, and the storage is reused.
This is because <b>strings don't have a null byte</b>, and are immutable. So, `s`, and the substings `hello`, and `world`, can reuse the memory.

```go
s[5] = 'a' // not allowed
s += "!"   // is allowed, creates another pointer to new memory with s + the update.
```
## Functions:

```go
s := "a string"
x := len(s) // build-in, = 8

strings.Contains(s, "g") // true
strings.Contains(s, "x") // false

strings.hasPrefix(s, "a") // true
strings.Index(s, "string") // = 2

s = strings.ToUpper(s) // "A STRING" (will be a copy in new memory block pointing to s, s doesn't change unless it's not being used and is garbage collected.)
```
