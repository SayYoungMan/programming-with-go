# Parallel Execution

- Two programs execute in `parallel` if they execute at exactly the same time
- At time t, an instruction is being performed for both P1 and P2
- Need replicated hardware

## Why Use Parallel Execution

- Tasks may complete more quickly
- Some tasks must be performed sequentially (so not always guaranteed faster execution)

# Von Neumann Bottleneck

## Speedup without Parallelism?

- Design faster processors
- Design processor with more memory
  - To reduce Von Neumann bottleneck where even if clock rate increases, the code execution is only slightly faster due to memory access speed
  - Cache access time = 1 clock cycle
  - Main memory access time = ~100 clock cycles
  - Thus, increasing on-chip cache improves performance

## Moore's law

- Predicted that transistor density would double every two years
- Smaller transistors switch faster
- Exponential increase in density would lead to exponential increase in speed

# Power Wall

## Power / Temperature Problem

- Transistors consume power when they switch
- Increasing transistor density leads to increased power consumption
  - Smaller transistors use less power, but density scaling is much faster
- High power leads to high temperature but air cooling can only remove so much heat

## Dynamic Power

- $P = \alpha \times CFV^2$
- $\alpha$ is percent of time switching
- C is capacitance which reduces with the size of transistor
- F is the clock frequency
- V is voltage swing (important)

## Dennard Scaling

- Voltage should scale with transistor size
- Keeps power consumption, and temperature, low
- Problem: Voltage can't go too low
  - must stay above `threshold voltage`
  - less noise tolerant
- Problem doesn't consider `leakage power` (insulator not thick enough so voltage leaking from conductor to conductor)
- Dennard scaling must stop

## Multi-core systems

- Cannot increase frequency
- Can still add processor cores, without increasing frequency (trend today)
- Code made to execute on multiple cores will exploit this architecture

# Concurrent vs Parallel

## Concurrent Execution

- Not necessarily the same as parallel execution
- `Concurrent`: start and end times overlap
- `Parallel`: execute at the exactly the same time

- Parallel tasks must be executed on different hardware where concurrent does not
- Mapping from tasks to hardware is not directly controlled by the programmer

## Concurrent Programming

- Programmer determines which tasks can be executed in parallel
- Mapping tasks to hardware
  - Operating system
  - Go runtime scheduler

## Hiding Latency

- Concurrency improves performance, even without parallelism
- Tasks must periodically wait for something (i.e. wait for memory)
- Other concurrent tasks can operate while one task is waiting

## Hardware Mapping in Go

- Programmer does not determine the hardware mapping
- Programmer makes parallelism possible
- Hardware mapping depends on many factors
