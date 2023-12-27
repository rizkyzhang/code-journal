---
tags:
  - docker-problem
---
## Problem

`Error saving credentials: error storing credentials - err: exec: "docker-credential-desktop": executable file not found in $PATH, out: ``
`
## Solution

`rm ~/.docker/config.json`

## Reference

https://stackoverflow.com/questions/71770693/error-saving-credentials-error-storing-credentials-err-exit-status-1-out