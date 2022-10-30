## Problem

I generated a key pair on my phone in termux, added the pub to the server and connected fine BUT only after being prompted with "Enter passphrase for key..." and entering the key's passphrase. In the hope of stopping the prompt from cropping up I ran...

```
$ eval $(ssh-agent)
$ ssh-add
```

This seemed to solve the problem. I could exit from the ssh connection and reconnect...

`$ ssh <username>@0.0.0.0`
...without the annoying prompt. Great! UNTIL I closed termux and restarted it. Then the !#@! prompt was back again! And, i think I know why. After the termux restart, I ran...

`$ ssh-add`
...and the agent pid had changed. So, how can I register the ssh agent to keep the freaking prompt from coming back and making me reenter the passphrase if the pid changes every time I stop and start termux? I assume this is happening because, unlike stopping the terminal emulator on an actual Linux box, stopping and restarting termux is essentially shutting down and rebooting the whole Linux instance.

## Causes

Termux is not a Linux instance. That's just a terminal app.

On fresh start the $PREFIX/tmp directory will be wiped. The socket of ssh-agent will gone.

## Solution

Put below ssh commands on .zshrc so it will be run on termux startup

```
eval $(ssh-agent)
ssh-add
```
