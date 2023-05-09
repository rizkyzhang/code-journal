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

https://stackoverflow.com/questions/53516594/why-do-i-keep-getting-eslint-delete-cr-prettier-prettier
