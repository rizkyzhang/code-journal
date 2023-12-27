---
tags:
  - linux-cheatsheet
---
- Get IP address `hostname -I | awk '{print $1}'`
- Download file `wget -P -O`
  - `-P` = output path
  - `-O` = file name

---

## SCP (Secure Copy Protocol)

`scp <options> source_path destination_path`

### Examples

```bash
scp file user@host:/path/to/file # copying a file to the remote system

scp user@host:/path/to/file file # copying a file from the remote system

scp file1 file2 user@host:/path/to/directory   # copying multiple files

scp -r directory user@host:/path/to/directory  # Copying an entire directory
```

### Options

- P(Caps) = specifies the port to establish connection with the remote host.
- p(lowercase) = preserves the times-tamp for ease of modification and access.
- r = copies the entire directory recursively
- q = copies files quietly, doesn't display the progress messages. Also known as quiet mode.
- C = for compression of data during transmission.
- v = verbose

---

## SFTP (Secure File Transfer Protocol)

`sftp user@host`

### Commands

- `put [local path] [remote path]` = Upload file
- `get [remote path] [local path]` = Download file
- `lcd` = Change the local directory
- `lmkdir` = Create a local directory
- `lpwd` = Display local working directory
- `lls` = Display the contents of the local working directory
