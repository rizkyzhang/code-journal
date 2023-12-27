---
tags:
  - golang-util
---
```go
package main

func BoolAddr(b bool) *bool {
	boolVar := b
	return &boolVar
}
```
