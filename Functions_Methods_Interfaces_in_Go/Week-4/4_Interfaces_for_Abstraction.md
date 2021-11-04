# Polymorphism
- Ability for an object to have different forms depending on the context
- *Identical* at a high level of abstraction
- *Different* at a low level of abstraction

## Inheritance
- Go does not support inheritance
- Subclass inherits the methods/data of the superclass

## Overriding
- Subclass redefines a method inheritied from the superclass
- Method can be polymorphic: different implementations for each class
    - Same signature (name, params, return)

# Interfaces
- Set of method signatures
    - Name, parameters, return values
    - Implementation is not defined
- Used to express conceptual similarity between types

## Satisfying an Interface
- Type satisfies an interface if type defines all methods specified in the interface
- Similar to inheritance with overriding

## Defining an Interface Type
```go
type Shape2D interface {
    Area() float64
    Perimeter() float64
}

type Triangle {...}
func (t Triangle) Area() float64 {...}
func (t Triangle) Perimeter() float64 {...}
```
- Triangle type satisfies the Shape2D interface
- No need to state it explicitly

# Interface vs Concrete Types

## Concrete Types
- Specify the exact representation of the data and methods
- Complete method implementation is included

## Interface Types
- Specifies some method signatures
- Implementations are abstracted

## Interface Values
- Can be treated like other values
    - Assigned to variables
    - Passed, returned
- Interface values have two components
    1. *Dynamic Type*: Concrete type which it is assinged to
    2. *Dynamic Value*: Value of the dynamic type
- Interface value is actually a pair
    - (dynamic type, dynamic value)

### Defining an interface type
```go
type Speaker interface {Speak ()}

type Dog struct {name string}
func (d Dog) Speak() {
    fmt.Println(d.name)
}
func main() {
    var s1 Speaker
    var d1 Dog{"Brian"}
    s1 = d1
    s1.Speak()
}
```
- Dynamic type is Dog, Dynamic value is d1

### Interface with Nil Dynamic Value
- An interface can have a nil dynamic value
```go
var s1 Speaker
var d1 * Dog
s1 = d1
```
- d1 has no concrete value yet
- s1 has a dynamic type but no dynamic value

## Nil Dynamic Value
- Can still call the `Speak()` method of s1
- Doesn't need a dynamic value to call
- Need to check inside the method

## Nil Interface Value
- Interface with `nil dynamic type`
- Very different from an interface with a nil dynamic value
` var s1 Speaker`
- Cannot call a method, runtime error

# Using Interfaces
 
## Ways to Use an Interface
- Need a function which takes multiple types of parameter
- Function `foo()` parameter
    - type X or type Y
- Define interface Z
- `foo()` parameter is interface Z
- Types X and Y satisfy Z
- Interface methods must be those needed by `foo()`

## Empty Interface
- Empty interface specifies no methods
- All types satisfy the empty interface
- Use it to have a function accept any type as a parameter

# Type Assertions

## Concealing Type Differences
- Interfaces hide the diffrerences between types
- Sometimes you need to treat different types in different ways

## Exposing Type Differences
- `DrawShape(s Shape2D) {}` will draw any shape
- Underlying API has different drawing functions for each shape
- Concrete type of shape s must be determined

## Type Assertions for Disambiguation
- Type assertions can be used to determine and extract the underlying concrete type
- Type assertion extracts Rectangle from Shape2D
    - Concrete type in parentheses
- If interface contains concrete type
    - `rect == concrete type, ok == true`
- If interface does not contain concrete type
    - `rect == zero, ok == false`

## Type Switch
- Switch statement used with a type assertion
```go
func DrawShape(s Shape2D) bool {
    switch := sh := s.(type) {
        case Rectangle:
            DrawRect(sh)
        case Triange:
            DrawRect(sh)
    }
}
```

# Error Handling

## Error Interface
- Many Go programs return error interface objects to indicate errors
```go
type error interface {
    Error() string
}
```
- Correct operation: `error == nil`
- Incorrect operation: `Error()` prints error message

## Handling Errors
- Check whether the error is nil
- If it is not nil, handle it