---
tags:
  - typescript-problem
  - typescript-problem-enum
---
## Problem

`Type 'string' is not assignable to type "Orange" | "Apple" | "Banana"`

```js
export type Fruit = "Orange" | "Apple" | "Banana";
let banana: string = "Banana";
let myFruit: Fruit = banana; // error on this line
```

## Solution

```js
let banana = "Banana" as const
```

## Reference

https://stackoverflow.com/questions/37978528/typescript-type-string-is-not-assignable-to-type
