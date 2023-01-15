- SSH client is used to connect to other computer
- SSH server is used to accept remote connection
- The computer you want to connect from need to have `openssh-client` installed and the remote computer need to have `openssh-server` installed
- If the remote computer doesn't have `openssh-server` installed, you will get `connection refused` error, to test if it is installed, run `ssh localhost` command on it, if it throws `connection refused` error then it is not installed
- There are two ways to authorize public key in a remote computer, copy the public key content manually to `authorized_keys` file or use `ssh-copy-id -i ~/.ssh/public-key-path remote-computer-address-or-alias-in-config-file` command to copy the public key automatically to the remote computer
- `tail -f /var/log/auth.log` view live log of ssh traffic

## Setup remote SSH connection

**Note: make sure openssh-server has been installed on the remote computer**

1. `ssh-keygen -t rsa -f ~/.ssh/id_rsa_test`
2. `ssh-add ~/.ssh/id_rsa_test`
3. Add host to `~/.ssh/config`

```
Host my-computer
	Hostname 11.111.111.11
	Port 22
	User ubuntu
	IdentityFile ~/.ssh/id_rsa_test
```

4. `ssh-copy-id -i ~/.ssh/id_rsa_test my-computer`
5. `ssh my-computer`
