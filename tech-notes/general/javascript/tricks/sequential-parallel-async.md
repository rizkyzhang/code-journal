---
tags:
  - js-trick
---
## Reading in sequence

If you want to read the files in sequence, you cannot use `forEach` indeed. Just use a modern `for-of` loop instead, in which await will work as expected:

```js
async function printFiles() {
  const files = await getFilePaths();

  for (const file of files) {
    const contents = await fs.readFile(file, "utf8");
    console.log(contents);
  }
}
```

## Reading in parallel

If you want to read the files in parallel, you cannot use forEach indeed. Each of the async callback function calls does return a promise, but you're throwing them away instead of awaiting them. Just use map instead, and you can await the array of promises that you'll get with Promise.all:

```js
async function printFiles() {
  const files = await getFilePaths();

  await Promise.all(
    files.map(async (file) => {
      const contents = await fs.readFile(file, "utf8");
      console.log(contents);
    })
  );
}
```

Reference: https://stackoverflow.com/a/37576787
