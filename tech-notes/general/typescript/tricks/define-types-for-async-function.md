## Define types for async function

To type an async function in TypeScript, set its return type to Promise<type>. Functions marked async are guaranteed to return a Promise even if you don't explicitly return a value, so the Promise generic should be used when specifying the function's return type.

```ts
type GetNumber = (num: number) => Promise<number>;

const getNumber: GetNumber = async (num) => {
  const result = await Promise.resolve(num);

  return result;
};
```

If the async function does not return a value, you should set the return type of the async function to be Promise<void>.

If the async function throws an error, you should set its return type to Promise<never>.

## Reference

https://bobbyhadz.com/blog/typescript-async-function-type
