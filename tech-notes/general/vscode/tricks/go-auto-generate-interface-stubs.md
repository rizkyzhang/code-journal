---
tags:
  - vscode-trick
---

## Go auto generate interface stubs

For instance, if you have some code like

```go
type MyType interface {
Foo(num int) error
Bar() (int, error)
}

type myStruct struct {
a, b int
}
```

1. `Ctrl + Shift + p`
2. Go: Generate Interface Stubs
3. Type `s *myStruct MyType` on the input

## Reference

https://forum.golangbridge.org/t/vs-code-go-auto-generate-missing-method-func-stubs/25050/3
