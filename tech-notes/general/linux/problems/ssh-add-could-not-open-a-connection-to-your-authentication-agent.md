---
tags:
  - linux-problem
---
## Problem

ssh-add could not open a connection to your authentication agent

## Solution

`eval $(ssh-agent)`

ssh-agent outputs the environment variables you need to have to connect to it:

```
shadur@proteus:~$ ssh-agent
SSH_AUTH_SOCK=/tmp/ssh-492P67qzMeGA/agent.7948; export SSH_AUTH_SOCK;
SSH_AGENT_PID=7949; export SSH_AGENT_PID;
echo Agent pid 7949;
shadur@proteus:~$
By calling eval you immediately load those variables into your environment.
```

As to why ssh-agent can't do that itself... Note the word choice. Not "won't", "can't". In Unix, a process can only modify its own environment variables, and pass them on to children. It can not modify its parent process' environment because the system won't allow it. This is pretty basic security design.

You could get around the eval by using ssh-agent utility where utility is your login shell, your window manager or whatever other thing needs to have the SSH environment variables set. This is also mentioned in the manual.

## Reference

https://unix.stackexchange.com/a/351727
