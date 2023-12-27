## Problem

Error response from daemon: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: error during container init: unable to apply apparmor profile: apparmor failed to apply profile: write /proc/self/attr/apparmor/exec: no such file or directory: unknown

## Solution

`sudo apt install apparmor`

## Reference

https://community.mailcow.email/d/2222-apparmor-problem-after-update-debian-11-container-not-starting/4
