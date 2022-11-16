## Problem

Assignment entry nil map

```go
var m map[string]string
m["cat"] = "cool"
```

`panic: assignment to entry in nil map`

## Solution

You have to initialize the map using the make function or a map literal before you can add any property:

```go
m := make(map[string]string)
m["cat"] = "cool"
```

```go
m := map[string]string{}
m["cat"] = "cool"
```
