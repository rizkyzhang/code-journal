---
tags:
  - next-js-trick
---
# Add border radius to Image by adding overflow hidden to parent component

```jsx
<div style={{ borderRadius: "5px", overflow: "hidden" }}>
  <Image src="..." layout="fill" objectFit="cover" />
</div>
```

## Reference

https://stackoverflow.com/questions/68920647/how-to-add-border-radius-to-next-js-image
