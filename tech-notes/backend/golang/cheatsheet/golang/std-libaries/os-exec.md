## ⌨ A very basic os/exec function call

Lets start from here, this is a basic exec function call and the results are returned in "out" as []byte

```go
func main() {
    cmd := exec.Command("ls", "-lah")
    out, err := cmd.CombinedOutput()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    fmt.Printf("combined out:\n%s\n", string(out))
}
```

## 🌲 os/exec function call with extra Environment variable

```go
cmd := exec.Command("programToExecute")
additionalEnv := "FOO=bar"
newEnv := append(os.Environ(), additionalEnv)
cmd.Env = newEnv
out, err := cmd.CombinedOutput()
if err != nil {
log.Fatalf("cmd.Run() failed with %s\n", err)
}
fmt.Printf("%s", out)
```

## 📁 Set the cwd of the execution path

```go
cmd:= exec.Command("git", "log")
cmd.Dir = "your/intended/working/directory"
out, err := cmd.Output()
```

Notes: If you are calling a binary instead of a variable / program inside the ENV variable (or %PATH% in Window's case), use the absolute path of the binary. For example:

```go
abspath, _ := filepath.Abs("./mybinary")
cmd:= exec.Command(abspath, "log")
cmd.Dir = "your/intended/working/directory"
out, err := cmd.Output()
```

## ⌨️ Use os/exec with Linux pipe / bash commands

```go
rcmd := `iw dev | awk '$1=="Interface"{print $2}'`
cmd := exec.Command("bash", "-c", rcmd)
out, err := cmd.CombinedOutput()
if err != nil {
log.Println(err.Error())
}
log.Println(string(out))
```

## 🖥️ Use os/exec with Windows Batch commands

```go
cmd := exec.Command("cmd", "/c", "ffmpeg -i myfile.mp4 myfile.mp3 && del myfile.mp4")
out, err := cmd.CombinedOutput()
if err != nil {
log.Println(err.Error())
}
log.Println(string(out))
```

## 📑 Use string slice in os/exec function call

```go
args := []string{"hello", "world"}
cmd := exec.Command("echo", args...)
out, err := cmd.Output()
if err != nil {
fmt.Println(err)
}
fmt.Println(string(out))
```

## 🍆 Solving localization problem in Windows CMD for os/exec

This is suitable for developers that is using a non-Unicode supported version of Windows with default codepage of cmd localized in not English language.

```go
cmd := exec.Command("cmd", "/c", "chcp 65001 && netsh WLAN show drivers")
out, err := cmd.CombinedOutput()
if err != nil {
log.Println(err.Error())
}
log.Println(string(out))
```

Notes: Remember to remove the first line of the output which should be a confirmation of successfully changed codepage using chcp command

Reference: https://dev.to/tobychui/quick-notes-for-go-os-exec-3ejg
