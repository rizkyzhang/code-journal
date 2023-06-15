- User info stored in `/etc/passwd`
- User password stored in `/etc/shadow`
- Group info store in `/etc/group`
- There are 3 types of user: root, regular and system

## User

- Show user info `id <username>`
- Create user `useradd <username>`
- Set password to user `passwd <username>`
- Switch user `su - <username>`
- Switch to root user `sudo -i`
- Add user to group `usermod -aG <group-name> <username>`
- Delete user `userdel -r <username>`
- List files opened by user `lsof -u <username>`

## Group

- Create group `groupadd <name>`
- Delete group `groupdel <name>`
