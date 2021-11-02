# First-Class Values

## Functions are First-class
- Functions can be treated like other types
    - Variables can be declared with a function type
    - Can be created dynamically
    - Can be passed as argumetns and returned as values
    - Can be stored in data structures

## Variables as Functions
- Declare a variable as a func
```go
var funcVar func(int) int
    func incFn(x int) int {
        return x + 1
    }
func main() {
    funcVar = incFn
    fmt.Print(funcVar(1))
}
```
- Function is on right-hand side, without ()

## Functions as Arguments
- Functions can be passed to another function as an argument
```go
func applyIt(afunct func (int) int, val int) int {
    return afunct(val)
}
```

## Anonymous Functions
- Don't need to name a function
```go
func main() {
    v := applyIt(func (x int) int {return x+1}, 2)
    fmt.Println(v)
}
```

# Returning Functions

## Functions as Return Values
- Functions can return functions
- Might create a function with controllable parameters

## Environment of a Function
- Set of all names that are valid inside a function
- Names defined locally, in the function
- **Lexical Scoping**
- Environment includes names defined in block where the function is defined

## Closure
- Function + its environment
- When functions are passed/returned, their environment comes with them

# Variadic and Deferred

## Variable Argument Number
- Functions can take a variable number of arguments
- Use ellipsis ... to specify
- Treated as a slice inside function

## Variadic Slice Argument
- Can pass a slice to a variadic function
- Need the ... suffix

## Deferred Function Calls
- Call can be deferred until the surrounding function completes
- Typically used for cleanup activities
```go
func main() {
    defer fmt.Println("Bye!")

    fmt.Println("Hello!")
}
```

### Deferred Call arguments
- Arguments of a deferred call are evaluated immediately