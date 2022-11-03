## Built-in Basic Types in Go

Go supports following built-in basic types:

- one boolean built-in boolean type: `bool`.
- 11 built-in integer numeric types: `int8`, `uint8`, `int16`, `uint16`, `int32`, `uint32`, `int64`, `uint64`, `int`, `uint`, and `uintptr`.
- two built-in floating-point numeric types: `float32` and `float64`.
- two built-in complex numeric types: `complex64` and `complex128`.
- one built-in string type: `string`.

Each of the 17 built-in basic types belongs to one different kind of type in Go. We can use the above built-in types in code without importing any packages, though all the names of these types are non-exported identifiers.

15 of the 17 built-in basic types are numeric types. Numeric types include integer types, floating-point types and complex types.

Go also support two built-in type aliases,

- `byte` is a built-in alias of `uint8`. We can view `byte` and `uint8` as the same type.
- `rune` is a built-in alias of `int32`. We can view `rune` and `int32` as the same type.

The integer types whose names starting with an `u` are unsigned types. Values of unsigned types are always non-negative. The number in the name of a type means how many binary bits a value of the type will occupy in memory at run time. For example, every value of the `uint8` occupies 8 bits in memory. So the largest `uint8` value is 255 (28-1), the largest `int8` value is 127 (27-1), and the smallest `int8` value is -128 (-27).

## Zero Values

Each type has a zero value. The zero value of a type can be viewed as the default value of the type.

- The zero value of a boolean type is false.
- The zero value of a numeric type is zero, though zeros of different numeric types may have different sizes in memory.
- The zero value of a string type is an empty string.

## Scopes

- There are 3 types of scope in GO:

1. Package scope
2. File scope = import statement
3. Block scope = code inside curly braces
