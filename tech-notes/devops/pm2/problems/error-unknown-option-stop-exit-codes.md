---
tags:
  - pm2-problem
---
## Problem

`error: unknown option '--stop-exit-codes'`

## Solution

pm2 as of now doesn't support `--stop-exit-code 0` parameter although it is written on the docs, this means all expected error will cause pm2 to restart when using express centralized/global error handling middleware, the solution is to catch error and handle it immediately inside `catch` block of router handler.

## Reference

[https://github.com/Unitech/pm2/issues/5208](https://github.com/Unitech/pm2/issues/5208).
