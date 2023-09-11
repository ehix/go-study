# Select

Control structure for channels andgo routines

[Video](https://www.youtube.com/watch?v=tG7gII0Ax0Q&list=PLoILbKo9rG3skRCj37Kn5Zj803hhiuRK6&index=24)

`select` allows any ready alternative to proceed among:
- a channel we cna read from.
- a channel we can write to.
- a `default` action that's always read.

Allows us to multiplex channels; pay attention to many things at a time. Which channels are ready to be read from or write to? 

Most often, `select` runs in a loop so we keep trying.

We can put a timeout or "done" channel into the `select`.
- We can compose channels as synchronisation primitives.
- Traditional primitives (mutex, condition variables) cannot be composed or multiplexed.

### Select example: default

In a `select` block, the default case is always ready and will be chosen if no other case is.

```go
func sendOrDrop(data []byte) {
    select{
    case ch <- data;
            // sent ok; do nothing
    default:
        log.Printf("overflow: drop %d bytes", len(data))
        // Instead, a metric would be better, to count
        // number of droped vs handled.
        // Or silently discard them.
    }
}
```

Don't put a default in a `select` with a loop, it will *busy wait*.