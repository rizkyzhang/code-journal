## Problem

No GO files found

## Solution

Under build in the air.toml; change the cmd property. If the entry of your golang application is in a different directory, say cmd, you can change the property to something like this:
`cmd = "go build -o ./tmp/main ./cmd/main.go"`
