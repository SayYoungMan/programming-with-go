# Synchronized Communication

## Blocking on Channels

### Iterating Through a Channel
- Common to iteratively read from a channel
    ```go
    for i := range c {
        fmt.Println(i)
    }
    ```
- Continues to read from channel c
- One iteration each time a new value is received
- i is assigned to the read value
- Iterates when sender calls `close(c)`

### Receiving from Multiple Goroutines
- Multiple channels may be used to receive from multiple sources
- Data from both sources may be needed
- Read sequentially
    ```go
    a := <- c1
    b := <- c2
    fmt.Println(a*b)
    ```

### Select Statement
- May have a choice of which data to use
    - First-come first-served
- Use the `select` statement to wait on the first data from a set of channels
    ```go
    select {
        case a = <- c1:
            fmt.Println(a)
        case b = <- c2:
            fmt.Println(b)
    }
    ```

## Select

### Select Send or Receive
- May select either send or receive operations
    ```go
    select {
        case a = <- inchan:
            fmt.Println("Received a")
        case outchan <- b:
            fmt.Print("Sent b")
    }
    ```

### Select with an Abort Channel
- Use select with a `separate abort channel`
- May want to receive data until an abort signal is received
    ```go
    for {
        select {
            case a <- c:
                fmt.Println(a)
            case <-abort:
                return
        }
    }
    ```

### Default Select
- May want a default operation to avoid blocking
    ```go
    select {
        case a = <- c1:
            fmt.Println(a)
        case b = <- c2:
            fmt.Println(b)
        default:
            fmt.Println("nop")
    }
    ```

## Mutual Exclusion

### Goroutine Sharing Variables
- Sharing variables concurrently can cause problems
- Two goroutines writing to a shared variable can interfere with each other
- `Concurrency-Safe`: Function can be invoked concurrently without interfering with other goroutines

### Granularity of Concurrency
- Concurrency is at the machine code level not source code level
- Interleaving machine instructions causes unexpected results

## Mutex

### Correct Sharing
- Don't let 2 goroutines write to a shared variable at the same time!
- Need to restrict possible interleavings
- Access to shared variables cannot be interleaved
- `Mutual Exclusion`: Code segments in different goroutines which cannot execute concurrently
- Writing to shared variables should be mutually exclusive

### Sync.Mutex
- A Mutex ensures mutual exclusion
- Uses `binary semaphore`
    - Flag up - shared variable is in use
    - Flag down - shared variable is available

## Mutex Methods

### Sync.Mutex Methods
- `Lock()` method puts the flag up
    - If lock is already taken by a goroutine, `Lock()` blocks until the flag is down
- `Unlock()` method puts the flag down
    - When `Unlock()` is called, a blocked `Lock()` can proceed

## Once Synchronization

### Synchronous Initialization
- `Initialization` must happen once and must happen before everything else
- Could perform initialization before implementing goroutines

### Sync.Once
- Has one method, `once.Do(f)`
- Function f is executed only one time
    - Even if it is called in multiple goroutines
- All calls to `once.Do()` block until the first returns

## Deadlock

### Synchronization Dependencies
- Synchronization cases the execution of different goroutines to depend on each other

### Deadlock
- `Circular dependencies` cause all involved goroutines to block
- Can be caused by waiting on channels

### Deadlock Detection
- Golang runtime automatically detects when all goroutines are deadlocked
- Cannot detect when a subset of goroutines are deadlocked

## Dining Philosophers

### Dining Philosophers Problem
- Classic problem involving concurrency and synchronization
- Each chopstick is a mutex
- Each philosopher is associated with a goroutine and two chopsticks
- Deadlock happens if all the philosophers grab the left chopstick at the same time
- Solution: Each philosopher picks up lowest numbered chopstick first
    - But philosopher 4 is of low priority and this is called `starvation`