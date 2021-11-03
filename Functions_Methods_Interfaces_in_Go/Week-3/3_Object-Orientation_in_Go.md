# Classes and Encapsulation

## Classes
- Collection of data fields and functions that share a well-defined responsibility
- Classes are a templacte that contain data fields not data

## Object
- Instance of a class
- Contains real data

## Encapsulation
- Data can be protected from the programmer
- Data can be accessed only using methods
- Maybe we don't trust the programmer to keep data consistent

# Support for classes

## No "Class" Keyword
- Most OO languages have a class keyword
- Datafields and methods are defined inside a class block

## Associating Methods with Data
- Method has a **receiver type** that it is associated with
- Use dot notation to call the method
```go
type MyInt int

func (mi MyInt) Double() int {
    return int(mi*2)
}
func main() {
    v := MyInt(3)
    fmt.Println(v.Double())
}
```
- Object v is an implicit argument to the method with `call by value`

## Structs
- Struct types compose data fields
- Traditional features of classes
- Structs and methods together allow arbitrary data and functions to be composed

# Encapsulation

## Controlling Access
- Can define public functions to allow access to hidden data

### Controlling access to structs
- Hide fields of structs by starting field name with a lower-case letter
- Define public methods which access hidden data

# Point Receivers

## Limitations of Methods
- Receiver is passed implicitly as an argument to the method
- Method cannot modify the data inside the receiver

### Large Receivers
- If receiver is large, lots of copying is required

```go
func (p *Point) OffsetX(v float64)
{
    p.x = p.x + v
}
```
- Receiver can be a pointer to a type
- Call by reference, pointer is passed to the method

# Pointer Receivers, Referencing, Deferencing

## No Need to Dereference
- Point is referenced as p, not *p
- Dereferencing is automatic with . operator

## No Need to Reference
- Do not need to reference when calling the method

## Using Pointer Receivers
- Good programming practice:
    1. All methods for a type have *pointer receivers*, or
    2. All methods for a type have *non-pointer receivers*
- Mixing pointer/non-pointer receivers for a type will get confusing