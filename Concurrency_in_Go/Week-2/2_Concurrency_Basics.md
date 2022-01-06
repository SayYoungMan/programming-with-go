# Concurrency Basics

## Processes

- An instance of a running program
- Things unique to a process:
    1. Memory - Virtual address space, code, stack, heap, shared libraries
    2. Registers - fast memories storing program counter, data regs, stack pointer, etc.

### Operating Systems
- Allows many processes to execute concurrently
- Processes are switched quickly (20ms)
- User has the impression of parallelism due to scheduling
- Operating system must give processes fair access to resources

## Scheduling

### Scheduling Processes
- Operating system schedules processes for execution
- Gives the illusion of parallel execution

### Context Switch
- Control flow changes from one process to another

## Threads and Goroutines

### Threads vs. Processes
- Processes can be slow because context switch can take time. Take data, write to memory and read from memory back to register can take long.
- Thread share some context
- Many threads can exist in one process
- Less context so switching is faster
- OS Schedules threads rather than processes

### Goroutines
- Like a thread in Go
- Many Goroutines execute within a single OS thread

### Go Runtime Scheduler
- Schedules goroutines inside an OS thread
- Like a little OS inside a single OS thread
- `Logical processor` is mapped to a thread

## Interleavings
- Order of execution within a task is known
- Order of execution between concurrent tasks is unknown
- Interleaving of instructions between tasks is unknown
- Many interleavings are possible and must consider all possibilities
- Ordering is `non-deterministic`

## Race conditions
- Outcome depends on non-deterministic ordering
- Races occur due to communication

### Communication between Tasks
- Threads are largely independent but not completely independent
- Should aim to avoid communicating between goroutines as little as possible