## Setup Basic SSH

1. `ssh-keygen -t rsa -f ~/.ssh/id_rsa_test`
2. `ssh-add ~/.ssh/id_rsa_test`
3. Add host to `~/.ssh/config`

```
Host 11.111.111.11
	User ubuntu
	IdentityFile ~/.ssh/id_rsa_test
```

4. `ssh ubuntu@11.111.111.11`
