---
tags:
  - typescript-problem
---
## Problem

Typeerror: do not know how to serialize a bigint

## Solution

Define this code in `index.ts`

```ts
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unreachable code error
// eslint-disable-next-line no-extend-native, func-names
BigInt.prototype.toJSON = function () {
  const int = Number.parseInt(this.toString(), 10);
  return int ?? this.toString();
};
```
