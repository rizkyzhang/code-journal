Problem: The certificate for deb.nodesource expired.

```
## Confirming "bionic" is supported...

+ curl -sLf -o /dev/null 'https://deb.nodesource.com/node_14.x/dists/bionic/Release'

## Adding the NodeSource signing key to your keyring...

+ curl -s https://deb.nodesource.com/gpgkey/nodesource.gpg.key | gpg --dearmor | tee /usr/share/keyrings/nodesource.gpg >/dev/null

## Creating apt sources list file for the NodeSource Node.js 14.x repo...

+ echo 'deb [signed-by=/usr/share/keyrings/nodesource.gpg] https://deb.nodesource.com/node_14.x bionic main' > /etc/apt/sources.list.d/nodesource.list
+ echo 'deb-src [signed-by=/usr/share/keyrings/nodesource.gpg] https://deb.nodesource.com/node_14.x bionic main' >> /etc/apt/sources.list.d/nodesource.list

## Running `apt-get update` for you...

+ apt-get update
Ign:1 https://deb.nodesource.com/node_14.x bionic InRelease
Err:2 https://deb.nodesource.com/node_14.x bionic Release
  Certificate verification failed: The certificate is NOT trusted. The certificate chain uses expired certificate.  Could not handshake: Error in the certificate verification. [IP: 201.0.222.9 443]
Hit:3 http://security.ubuntu.com/ubuntu bionic-security InRelease
Hit:4 http://archive.ubuntu.com/ubuntu bionic InRelease
Hit:5 http://ppa.launchpad.net/brightbox/ruby-ng/ubuntu bionic InRelease
Hit:6 http://archive.ubuntu.com/ubuntu bionic-updates InRelease
Hit:7 http://archive.ubuntu.com/ubuntu bionic-backports InRelease
Hit:8 https://packagecloud.io/github/git-lfs/ubuntu bionic InRelease
Reading package lists...
E: The repository 'https://deb.nodesource.com/node_14.x bionic Release' does not have a Release file.
Error executing command, exiting
The command '/bin/bash -o pipefail -c curl -fsSL https://deb.nodesource.com/setup_14.x | bash - &&   apt-get install -y --no-install-recommends nodejs &&   npm i -g xunit-viewer &&   rm -rf /var/lib/apt/lists/*' returned a non-zero code: 1
```

Solution: `sudo apt install ca-certificates`

Reference: https://github.com/nodesource/distributions/issues/1266
