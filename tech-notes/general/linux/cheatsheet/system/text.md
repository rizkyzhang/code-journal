---
tags:
  - linux-cheatsheet
---
- `cut -d: -f1 file` Get first column separated by :
- `grep -iR <pattern> file` Search text that match pattern in a file insensitively and recursively
- `grep -vR <pattern> file` Search text that don't match pattern recursively
- `sed 's/word1/word2/g' file` Preview result of replacing word1 with word2 globally
- `sed -i 's/word1/word2/g' file` Replace immediately word1 with word2 globally
