# Conventional Synchronisation

[Video](https://www.youtube.com/watch?v=DtXNSE3Yejg&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=29)

Go's CSP model is likely better than this. Mutexes are essentially embedded in the channel logic, where as mutex require more thought to achieve the same thing.

Package `sync`, `sync/atomic`.

## Mutual exclusion

#### Preventing race conditions and adding protections:
What if multiple goroutines must read/write some data?

We must make sure only **one** of them can do so at any instant (in the so-called *"critical section"*); where the variable is being accessed or modified.

We accomplish this with some type of lock:
- aquire the lock before accessing the data
- any other goroutine will **block** waiting to get the lock
- release the lock when done


## Mutex in action
Better to embed a mutex into some other type.

Original maps aren't goroutine safe, so this map type implements an embedded mutex to make it goroutine safe.
```go
type SafeMap struct {
    sync.Mutex  // not sage to copy
    m map[string]int
}

// so methods must take a pointer, not a value
func (s *SafeMap) Incr(key string) {
    s.Lock()
    defer s.Unlock

    // only one goroutine can execute this
    // code at the same time, guaranteed
    s.m[key]++
}
```
Using mutex has a performance overhead.
- has same effect/cost on readers and writers

## RWMutexes in action
Sometimes we need to prefer readers to (infrequent) writers, RWMutexes are an optimisation to allow this.

Not embedded in this example.

It prefers reader, and allows for multiple reads and the same time.
All writers are blocked until unlocked, but readers will be able to get the read lock.
```go
type InfoClient struct {
    mu          sync.RWMutex
    token       string
    tokenTime   time.Time
    TTL         time.Duration
}

// Reading the token
func(i *InfoClient) CheckToken() (string, time.Duration) {
    i.mu.RLock()
    defer i.mu.RUnlock()

    return i.token, i.TTL - time.Since(i.tokenTime)
}

// Writing a new token when invalid
func(i *InfoClient) ReplaceToken(ctx context.Context) (string, error) {
    // Don't need the lock yet to read the token
    token, ttl, err := i.getAccessToken(ctx)
    if err != nil {
        return "", err
    }

    // Need to lock it, as this is a writer.
    i.mu.Lock()
    defer i.mu.Unlock()

    // These things need to be changed, so lock.
    i.token = token
    i.tokenTime = time.Now()
    i.TTL = time.Duration(ttl) * time.Second

    return token, nil
}
```
## Atomic primitives

More tied to the hardware and limited, but sometimes more effecient.

## Only-once execution

A `sync.Once` object allows us to ensure a function runs only once
(only the first call to `Do` will call the function passed in).

Singleton, one thing in the program that does this function.
In concurrent program, we need to only create it once.

```go
var once sync.Once
var x

func initialise() {
    x = NewSingleton()
}

func handle(w http.ResponseWriter, r *http.Request) {
    // do this:
    once.Do(initialise)
    // not this:
    if x == nil {
        create and assign the singleton
    } 
    // it's possible for two handlers to run and create the singleton at the same time.
    ...
}
```

## Pool

In a long running server, there may be things we use a lot of, connections for example, e.g. gRCP.

A `Pool` provides for efficient and safe reuse of objects, but it's a container of `interface{}`.

The pool just sees whats in the pools as empty interfaces.

```go
var bufPool = sync.Pool {
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

// Get something from the pool, use it, and put it back.
func Log(w io.Writer, key, val string) {
    b := bufPool.Get().(*bytes.Buffer) // more reflection, bc they're interfaces in the pool. Convert into a concrete type.
    
    b.Reset()
    // write to it
    w.Write(b.Bytes())
    bufPool.Put(b)
}
```