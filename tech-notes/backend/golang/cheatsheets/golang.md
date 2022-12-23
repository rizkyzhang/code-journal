## Number data types

- `uint8` 0 ↔ 255
- `uint16` 0 ↔ 65535
- `uint32` 0 ↔ 4294967295
- `uint64` 0 ↔ 18446744073709551615
- `uint`, same as uint32 or uint64 (depending on the value)
- `byte`, same as uint8
- `int8` -128 ↔ 127
- `int16` -32768 ↔ 32767
- `int32` -2147483648 ↔ 2147483647
- `int64` -9223372036854775808 ↔ 9223372036854775807
- `int`, same as int32 ot int64 (depending on the value)
- `rune`, same as int32

---

## Zero values

- 0 for all integer types,
- 0.0 for floating point numbers,
- false for booleans,
- "" for strings,
- nil for interfaces, slices, channels, maps, pointers and functions.

The elements of an array or struct will have its fields zeroed if no value is specified. This initialization is done recursively:

```go
type T struct {
	n int
	f float64
	next *T
}
fmt.Println([2]T{}) // [{0 0 <nil>} {0 0 <nil>}]
```

---

## Build

### Windows

```
# 64-bit
GOOS=windows GOARCH=amd64 go build -o bin/app-amd64.exe app.go

# 32-bit
GOOS=windows GOARCH=386 go build -o bin/app-386.exe app.go
```

### MacOS

```
# 64-bit
GOOS=darwin GOARCH=amd64 go build -o bin/app-amd64-darwin app.go

# 32-bit
GOOS=darwin GOARCH=386 go build -o bin/app-386-darwin app.go
```

### Linux

```
# 64-bit
GOOS=linux GOARCH=amd64 go build -o bin/app-amd64-linux app.go

# 32-bit
GOOS=linux GOARCH=386 go build -o bin/app-386-linux app.go
```

---

## JSON Decode

- Decode body into struct `json.NewDecoder(res.Body).Decode()`

---

## OK check

- Type casting operation will return the casted value and ok status (true if casting successful, vice versa) `metadataStr, ok := metadata.(string)`
- Accessing map property's value will return the value and ok status (true if property exist, vice versa) `metadata, ok := data["metadata"]`

---

## Print struct

- `%v` – It will print only values. Field name will not be printed. This is the default way of printing a struct. `fmt.Printf("%v", user)` - `{Sam 31 2000}`
- `%+v` – It will print both field and value. `fmt.Printf("%+v", emp)` - `{name:Sam age:31 salary:2000}`
- `%#v` – It will print the struct name, also both field and value. `main.employee` - `{name:"Sam", age:31, salary:2000}`
