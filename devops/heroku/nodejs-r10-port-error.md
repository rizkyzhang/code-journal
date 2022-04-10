## Problem

Node.js app crashing with an R10 error.

## Solution

Most Node.js apps bind to a specific port by default. For example, Express uses the following:

```js
app.listen(3000, function () {
  console.log("Example app listening on port 3000!");
});
```

However, Heroku dynos expose a dynamic port for your app to bind to. This value is exposed in the $PORT env var. You must change your code to bind to this port instead. For example:

```js
const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`Our app is running on port ${PORT}`);
});
```

This will use the $PORT env var if available, or fallback to a default port (useful for local development). Note that when browsing to your application on Heroku, you still use port 80 ([your-application].herokuapp.com) and not the port that your process binds on.

## Reference

https://help.heroku.com/P1AVPANS/why-is-my-node-js-app-crashing-with-an-r10-error
