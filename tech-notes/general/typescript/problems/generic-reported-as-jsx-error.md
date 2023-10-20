---
tags:
  - typescript-problem
  - typescript-problem-jsx
---
## Problem

TypeScript Generic Function Reported As JSX Error
## Solution

The usage of `<T>` prior to the function braces causes a JSX error within `.tsx` files: `"JSX element has no corresponding closing tag."`. 

The issue was claimed to be a limitation and there were a few workarounds:

- change from `.tsx `to `.ts`
- add a comma: `const f = <T,>(arg: T): T => {...}`
- extend this way: `const foo = <T extends unknown>(x: T) => x;`
- or extend this way: const foo = `<T extends {}>(x: T): T => x;`
## Reference

https://dev.to/tlylt/typescript-generic-function-reported-as-jsx-error-57nf
