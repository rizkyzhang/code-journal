---
tags:
  - jetbrain-problem
---
## Problem

Env not loaded when using desktop shortcut

## Solution

Change shortcut Exec to `/usr/bin/zsh -l -i -c "/home/username/bin/PhpStorm-231.9161.47/bin/phpstorm.sh" %u`

## Reference

https://youtrack.jetbrains.com/issue/IDEABKL-7589/Load-interactive-shell-environment-variables-on-Linux#focus=Comments-27-5120666.0-0