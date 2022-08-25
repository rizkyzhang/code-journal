## Problem

Segmentation fault error when running test

## Solution 1

Install and setup `ts-jest`

## Solution 2

Add these properties to the jest.config.ts

```js
{
  transform: {},
  transformIgnorePatterns: [],
  extensionsToTreatAsEsm: [".ts", ".tsx"],
}
```

## Reference

https://jestjs.io/docs/ecmascript-modules
