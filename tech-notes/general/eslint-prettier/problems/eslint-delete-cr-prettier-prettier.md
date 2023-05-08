## Problem

`[eslint] Delete 'CR' [prettier/prettier]`

## Solution

Add this into `.eslintrc` rules

```js
'prettier/prettier': [
  'error',
  {
    'endOfLine': 'auto',
  }
]
```

## Reference
