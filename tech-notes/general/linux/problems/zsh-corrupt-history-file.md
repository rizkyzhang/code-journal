## Problem

zsh: corrupt history file

## Solution

```
mv .zsh_history .zsh_history_old
strings .zsh_history_old > .zsh_history
fc -R .zsh_history
```

## Reference

https://gggauravgandhi.medium.com/fix-for-zsh-corrupt-history-file-home-zsh-history-with-restore-history-to-new-file-f817a289d4aa
