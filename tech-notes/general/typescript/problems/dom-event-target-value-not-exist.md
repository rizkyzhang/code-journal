---
tags:
  - typescript-problem
  - typescript-problem-jsx
---
## Problem

DOM `event.target.value` not exist

## Solution

The `EventTarget` type does not inherit from `HTMLElement` by default because HTML elements are not the only things that can be event targets. It’s left to you to determine what the proper type of the `target` object is before TypeScript can allow you to access any properties not found on `EventTarget`.

One way to fix this error is to type cast the `target` object with the `as` keyword.

```ts
const { value } = event.target as HTMLSelectElement;
```

## Reference

https://freshman.tech/snippets/typescript/fix-value-not-exist-eventtarget/
