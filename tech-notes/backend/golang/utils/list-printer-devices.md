---
tags:
  - golang-util
---
```go
package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	rcmd := `lpstat -p | awk '{print $2}'`
	cmd := exec.Command("bash", "-c", rcmd)
	_output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	output := strings.TrimSpace(string(_output))
	_printers := strings.Split(output, "\n")
	var printers []string

	for _, printer := range _printers {
		if printer != "to" {
			printers = append(printers, printer)
		}
	}

	fmt.Println(printers)
}
```
