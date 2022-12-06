## Problem

```go
file, err := os.Open("test")
if err != nil {
  return err
}
_, err = file.WriteString("This is a test")
if err != nil {
  return err
}
```

When running the code, got error `bad file descriptor`

## Cause

`file.WriteString` will only works if you want to write to the file for the first time.

```go
file, err := os.Create("test")
if err != nil {
  return err
}
_, err = file.WriteString("This is a test")
if err != nil {
  return err
}
```

By default `os.Open` only have a read-only `os.RDONLY` file descriptor set which mean you don't have permission to write like `os.Create`.

## Solution

You need to use `file.OpenFile` with permission flags: `os.O_CREATE|os.O_WRONLY|os.O_APPEND`

```go
file, err := os.OpenFile("test", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
if err != nil {
  return err
}
_, err = file.WriteString("This is a test")
if err != nil {
  return err
}
```

In DOS/WINDOWS you do not need to explicitly add a write flag with an append flag, as it is implied. Go in linux does need the extra flag added to work.

## Reference

https://stackoverflow.com/questions/33851692/golang-bad-file-descriptor
