## Problem

`Internal watch failed: ENOSPC: System limit for number of file watchers reached, watch ...`


## Solution

`echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p`


## Reference

https://stackoverflow.com/a/55543310
