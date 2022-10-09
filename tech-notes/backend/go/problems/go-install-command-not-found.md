## Problem

After installing global package with `go install`, error command not found.

## Solution

Add GOPATH binary path (usually $HOME/go/bin) to exported PATH in `.zshrc`

```zsh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```
