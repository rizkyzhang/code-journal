## Problem

Redshift Fopen: Permission Denied Unable To Load Config File

## Solution

- `sudo vim /etc/apparmor.d/usr.bin.redshift `
- Append this to the file, below `owner @{HOME}/.config/redshift.conf r,`

```
owner @{HOME}/.config/redshift/** r,
owner @{HOME}/.config/redshift/hooks/** ux,
```

- `sudo systemctl reload apparmor.service`

## Reference

https://github.com/jonls/redshift/issues/850#issuecomment-1430106510
