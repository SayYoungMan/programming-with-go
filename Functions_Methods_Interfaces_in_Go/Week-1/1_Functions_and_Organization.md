# Why Use Functions?

## What is a Function
- A set of instructions with a name
```go
func PrintHello() {
    fmt.Printf("Hello, world.")
}
func main() {
    PrintHello()
}
```
- Function declaration, name, call

## Reusability
- You only need to declare a function once
- Good for commonly used operations
    - Graphics editing program might have `ThresholdImage()`
    - Database program might have `QueryDbase()`
    - Music program might have `ChangeKey()`

## Abstraction
- Details are hidden in the function
- Only need to understand operations
- Improves understandability
- Naming is important for clarity

# Function Parameters
- Functions may need input data to perform their operations
- **Parameters** are listed in parenthesis after function name
- **Arguments** are supplied in the call

### Parameter Options
- If no parameters are needed, put nothing in parentheses
- List arguments of same type `func foo(x, y int)`

### Return Values
- Functions can return a values as a result
- Type of return value after parameters in declaration
- Function call used on right-hand side of an assignment
- For multiple return values, multiple value types must be listed in the declaration

# Call by Value, Reference

## Call by Value
- Passed arguments are copied to parameters
- Modifying parameters has no effect outside the function

### Tradeoffs of call by value
- Advantage: **Data encapsulation**
    - Function variables only changed inside the function
- Disadvantage: **Copying time**
    - Large objects may take a long time to copy

## Call by Reference
- Programmer can pass a pointer as an argument
- Called function has direct access to caller variable in memory

### Tradeoffs of call by reference
- Advantage: **Copying time**
    - Don't need to copy arguments
- Disadvantage: **Data encapsulation**
    - Function variables may be changed in called functions
    - May be what you want

# Passing Arrays and Slices

### Passing Array Arguments
- Array arguments are copied
- Array can be big, so this can be a problem

### Passing array pointers
- Possible to pass array pointers but messy and unnecessary

### Pass Slices instead
- Slices contain a pointer to the array
- Passing a slice copies the pointer
```go
func foo (sli []int) []int {
    sli[0] = sli[0] + 1
}
func main(){
    a := []int{1, 2, 3}
    foo(a)
    fmt.Print(a)
}
```

# Well-written Functions

## Understandability
- Code is functions and data
- If you are asked to find a feature, you can find it quickly
- If you are asked about where data is used, you know

## Debugging Principles
- Code crashes inside a function
- Two options for the cause
    1. Function is written incorrectly
    2. Data that the function uses is incorrect

### Supporting Debugging
- Functions need to be understandable
    - Determine if actual behaviour matches desired behaviour
- Data needs to be traceable
    - Where did that data come from?
    - Global variables complicate this

# Guidelines for Functions

## Function Naming
- Give functions a good name
    - Behaviour can be understood at a glance
    - Parameter naming counts too

## Functional Cohesion
- Function should perform only one "operation"
- An operation depends on the context
- Merging behaviours makes code complicated

## Few Parameters
- Debugging requires tracing function input data
- More difficult with a large number of parameters
- Function may have bad functional cohesion

### Reducing Parameter Number
- May need to group related arguments into structures

## Function Complexity
- **Function length** is the most obvious measure
- Functions should be simple
- Short functions can be complicated too

### Function Length
- How do you write complicated code with simple functions?
- **Function Call Hierarchy**

## Control-flow Complexity
- Control-flow describes conditional paths

### Partitioning Conditionals
- Functional hierarchy can reduce control-flow complexity