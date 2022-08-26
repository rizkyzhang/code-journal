## Problem

```
(JSX attribute) HTMLAttributes<HTMLDivElement>.style?: React.CSSProperties
Type '{ backgroundColor: string; display: string; flex: string; flexDirection: string; }' is not assignable to type 'CSSProperties'.
  Types of property 'flexDirection' are incompatible.
    Type 'string' is not assignable to type 'FlexDirectionProperty'.ts(2322)
index.d.ts(1763, 9): The expected type comes from property 'style' which is declared here on type 'DetailedHTMLProps<HTMLAttributes<HTMLDivElement>, HTMLDivElement>'
```

## Solution

- `flexDirection: 'row' as 'row'`
- `<div style={{flexDirection: 'row'} as React.CSSProperties}>`

## Reference

https://github.com/cssinjs/jss/issues/1344
