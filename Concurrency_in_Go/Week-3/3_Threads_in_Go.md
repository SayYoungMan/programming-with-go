# Threads in Go

## Goroutines

### Creating a Goroutine
- One goroutine is created automatically to execute the `main()``
- Other goroutines are created using the `go`keyword
- Goroutine does not block the other goroutine that it was called from

### Exiting a Goroutine
- A goroutine exits when its code is complete
- When the main goroutine is complete, all other goroutines exit
- A goroutine may not complete its execution because main completes early

## Exiting Goroutines

### Early Exit
```go
func main() {
    go fmt.Printf("New routine")
    fmt.Printf("Main routine")
}
```
- Only `Main routine`is printed
- Main finished before the new goroutine started
- We have no idea how scheduler is going to schedule goroutines

### Delayed Exit
```go
func main() {
    go fmt.Printf("New routine")
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("Main routine")
}
```
- Add a delay in the main routine to give the new routine a chance to complete
- This is a hacky attempt (should not do)

### Timing with Goroutines
- Adding a delay to wait for a goroutine is bad.
- Timing assumptions may be wrong
- Timing is Non-deterministic

## Basic Synchronization

### Synchronization
- Using global events whose execution is viewed by all threads, simultaneously
- It is used to restrict bad interleavings
- Can reduce the efficiency but used when necessary

## Wait Groups
### Sync WaitGroup
- Sync packages contains functions to synchronize between goroutines
- `sync.WaitGroup`forces a goroutine to wait for other goroutines
- Contains an internal counter
    - Increment counter for each goroutine to wait for
    - Decrement counter when each goroutine completes
    - Waiting goroutine cannot continue until counter is 0

### Using WaitGroup
- `Add()`increments the counter
- `Done()` decrements the counter
- `Wait()`blocks until counter == 0

## Communication
### Goroutine Communication
- Goroutines usually work together to perform a bigger task
- Often need to send data to collaborate

### Channels
- Transfer data between goroutines
- Channels are typed
- Use `make()`to create a channel
- Send and receive data using the <- operator
- Send data on a channel `c <- 3`
- Receive data from a channel `x := <- c`

## Blocking in Channels

### Unbuffered Channel
- Unbuffered channels cannot hold data in transit
    - Default is unbuffered
- Sending blocks until data is received
- Receiving blocks until data is sent

### Blocking and Synchronization
- Channel communication is synchronous
- Blocking is the same as waiting for communication
- Receiving and ignoring the result is the same as `Wait()`: `<- c`

## Buffered Channels
### Channel Capacity
- Channels can contain a limited number of objects
    - Default size 0 (unbuffered)
- Capacity is the number of objects it can hold in transit
- Optional argument to make() defines channel capacity: `c := make(chan int, 3)``
- Sending only blocks if buffer is full
- Receiving only blocks if buffer is empty

### Use of Buffering
- Sender and receiver do not need to operate at exactly the same speed
- Speed mismatch becomes acceptable