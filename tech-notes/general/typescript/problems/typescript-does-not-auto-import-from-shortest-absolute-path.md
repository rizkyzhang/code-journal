---
tags:
  - typescript-problem
  - typescript-problem-import
---
## Problem

TypeScript does not auto import from shortest absolute path

## Solution

Change the paths order in `tsconfig.json` so it goes like this:

```json
"paths": {
  "@assets/*": ["src/assets/*"],
  "@components/*": ["src/components/*"],
  "@hooks/*": ["src/hooks/*"],
  "@src/*": ["src/*"]
}
```

The longest path should be the latest one.

## Reference

https://stackoverflow.com/questions/74093001/vscode-typescript-does-not-auto-import-from-shortest-path
