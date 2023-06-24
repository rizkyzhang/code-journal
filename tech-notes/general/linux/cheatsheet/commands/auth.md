- User info stored in `/etc/passwd`
- User password stored in `/etc/shadow`
- Group info store in `/etc/group`
- There are 3 types of user: root, regular and system

## User

- Show user info `id [username]`
- Create user (centos) `useradd [username]`
- Create user (ubuntu) `adduser [username]`
- Set password to user `passwd [username]`
- Switch user `su - [username]`
- Switch to root user `sudo -i`
- Add user to group `usermod -aG [group-name] [username]`
- Delete user (centos) `userdel -r [username]`
- Delete user (ubuntu) `deluser -r [username]`
- List files opened by user `lsof -u [username]`

## Group

- Create group (centos) `groupadd [name]`
- Create group (ubuntu) `addgroup [name]`
- Delete group (centos) `groupdel [name]`
- Delete group (ubuntu) `delgroup [name]`
- Find all groups to which you belong `groups`

## Sudo

- Grant sudo privilege to user (ubuntu) `usermod -aG sudo [username]`
- Grant sudo privilege to user (centos) `usermod -aG wheel [username]`

## Permission

![](../permission-in-symbolic-notation.jpg)

| **PERMISSION TYPE** | **SYMBOL** | **FILE**                                      | **DIRECTORY**                                                                        |
| ------------------- | ---------- | --------------------------------------------- | ------------------------------------------------------------------------------------ |
| Read                | r          | Open and view file contents (cat, head, tail) | Read directory contents (ls, du)                                                     |
| Write               | w          | Edit, delete or rename file (vi)              | Edit, delete or rename directory and files within it; create files within it (touch) |
| Execute             | x          | Execute the file                              | Enter the directory (cd); without x, the directory’s r and w permissions are useless |
| None                | -          | Do nothing                                    | Do nothing                                                                           |

| **OCTAL DIGIT** | **PERMISSION(S) GRANTED** |
| --------------- | ------------------------- |
| 0               | None                      |
| 1               | x                         |
| 2               | w                         |
| 3               | wx 1 + 2 = 3              |
| 4               | r                         |
| 5               | rx 4 + 1 = 5              |
| 6               | rw 4 + 2 = 6              |
| 7               | rwx 4 + 2 + 1 = 7         |

- Change user and group ownership `chown [option] [username]:[group] [file]`
- Change group ownership `chgrp [group] [file]`
- Grant execute permission to user `chmod u+x [file]`
- Remove read permission from group `chmod g-r [file]`
- Remove write permission from other `chmod o-w [file]`
- Change multiple permissions `chmod u=rwx,g=rw,o= [file]`
- Full permission for user, group and other `chmod 777 [file]`
