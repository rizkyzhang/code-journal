## Problem

Docker VS Code Error: connect EACCES /var/run/docker.sock

## Solution

Turns out the error message is somewhat expected using Ubuntu (and Linux in general) when running Docker as a non-root user. There are a set of post-installation steps under "Manage Docker as a non-root user" after installing Docker on Linux and connecting to it as a non-root user. In my case I only had to run a couple of commands and re-start the Virtual Machine to fix the error.

The first command is to create a docker group.

`sudo groupadd docker`

Turns out the docker group had already been created when installing Docker on Ubuntu 20.04 so this was unecessary.

The second command is to add my user to the docker group.

`sudo usermod -aG docker $USER`
At this point the documentation says one should logout and login for the settings to take affect, but if Linux is running in a virtual machine, it may be necessary to restart the virtual machine. I am running Ubuntu 20.04 in a virtual machine on my MacBook Pro, but for kicks I tried to just logout and login, and it didn't fix the VS Code error. However, when I restarted the virtual machine, the VS Code error went away and I was able to successfully connect to my Docker images and containers using the Docker VS Code extension.

## Reference

https://www.davidhayden.me/blog/fix-docker-vs-code-error-connect-eacces-var-run-docker-sock
