---
tags:
  - "#git-problem"
---
## Problem

`git: 'credential-libsecret' is not a git command. See 'git --help'.`

## Solution

1. `sudo apt install libsecret-1-0 libsecret-1-dev make gcc 
2. `sudo make --directory=/usr/share/doc/git/contrib/credential/libsecret`
3. Go to the repository and run `git config credential.helper /usr/share/doc/git/contrib/credential/libsecret/git-credential-libsecret`

## Reference

- https://publish.obsidian.md/git-doc/Authentication
- https://my-take-on.tech/2019/08/23/safely-storing-git-credentials/