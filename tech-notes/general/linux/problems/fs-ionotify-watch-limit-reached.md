---
tags:
  - linux-problem
---
## Problem

`Error: ENOSPC: System limit for number of file watchers reached` or auto reload not working when running React server.

## Solution 1

Increase `max_user_watches` value temporarily.

`echo 1048576 > /proc/sys/fs/inotify/max_user_watches`

## Solution 2

Increase `max_user_watches` value permanently (persist after power off).

`echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p`

## Reference

https://stackoverflow.com/a/55543310
