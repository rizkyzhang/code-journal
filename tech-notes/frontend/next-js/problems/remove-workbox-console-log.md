## Problem

Remove workbox console.log

## Solution

`disable: process.env.NODE_ENV === "development"`

```js
const withPWA = require("next-pwa");
const runtimeCaching = require("next-pwa/cache");

module.exports = withPWA({
  pwa: {
    disable: process.env.NODE_ENV === "development",
    dest: "public",
    runtimeCaching,
  },
});
```

## Reference

https://stackoverflow.com/questions/65952421/nextjs-remove-workbox-console-log-messages
