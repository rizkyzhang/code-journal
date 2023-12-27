---
tags:
  - linux-trick
---
Add these into `~/.ssh/config`

```
ServerAliveInterval 60
ServerAliveCountMax 120
```

The `ServerAliveInterval` will send a keepalive every x seconds (default is 0, which disables this feature if not set to something else). This will be done `ServerAliveCountMax` times if no response is received. The default value of `ServerAliveCountMax` is 3.

## Reference

https://superuser.com/a/1285742
