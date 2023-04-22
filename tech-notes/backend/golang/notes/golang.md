- Variable = a named memory address that stores a value
- Pointer = a variable that stores a memory address
- Reference = an alias of a variable, modifying it will also modify the original variable

```go
// Integer named x is set to 2
var x int = 2

// Integer pointer named pX is set to the memory address of x
var pX *int = &x

// Integer named y is set to the value pointed to by pX
var y int = *pX
```

- Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

- The functions in the encoding/json standard package will only encode and decode the exported struct fields.
- Two unnamed struct types are identical only if they have the same sequence of field declarations. Two field declarations are identical only if their respective names, their respective types and their respective tags are all identical. **Please note, two non-exported struct field names from different packages are always viewed as two different names**.
