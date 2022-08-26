## Problem

`Resolve error: typescript with invalid interface loaded as resolver import react` or absolute path missing file extension

## Solution 1

`yarn add -D eslint-import-resolver-typescript`

```js
settings: {
  "import/resolver": {
    typescript: {
      alwaysTryTypes: true,
      project: "./tsconfig.json",
    },
  },
},
```

## Solution 2

```js
rules: {
  "import/extensions": "off",
}
```

## Reference

https://stackoverflow.com/questions/69446204/resolve-error-typescript-with-invalid-interface-loaded-as-resolver-eslint
