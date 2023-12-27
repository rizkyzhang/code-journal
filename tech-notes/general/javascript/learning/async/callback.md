## Callback

- Function you’re passing as an argument is called a **callback function** and the function you’re passing the callback function to is called a **higher order function**.

```js
function add(x, y) {
  return x + y;
}

function higherOrderFunction(x, callback) {
  return callback(x, 5);
}

higherOrderFunction(20, add);
```

- There are two popular uses for callback:

1. Abstraction over transforming one value into another. For example `map` and `filter` function.
2. Delaying execution of a function until a particular time. For example when a button is clicked, it will run a callback function.

- The main issue of callback is Callback hell. As humans, we naturally think sequentially. When you have nested callbacks inside of nested callbacks, it forces you out of your natural way of thinking. Bugs happen when there’s a disconnect between how your software is read and how you naturally think.

- The next issue of callbacks has to do with inversion of control. When you write a callback, you’re assuming that the program you’re giving the callback to is responsible and will call it when (and only when) it’s supposed to. You’re essentially inverting the control of your program over to another program. When you’re dealing with libraries like jQuery, lodash, or even vanilla JavaScript, it’s safe to assume that the callback function will be invoked at the correct time with the correct arguments. However, for many third-party libraries, callback functions are the interface for how you interact with them. It’s entirely plausible that a third party library could, whether on purpose or accidentally, break how they interact with your callback.

## Promise

- Promise can be in one of three states, pending, fulfilled or rejected. They represent the status of an asynchronous request. If the async request is still ongoing, the Promise will have a status of pending. If the async request was successfully completed, the Promise will change to a status of fulfilled (resolve callback executed). If the async request failed, the Promise will change to a status of rejected (reject callback executed).

- When the status of the promise changes to fulfilled, the function that was passed to `.then` will get invoked. When the status of a promise changes to rejected, the function that was passed to `.catch` will be invoked.

- Both `.then` and `.catch` will return a new promise. That seems like a small detail but it’s important because it means that promises can be chained.

## Async/await

- If the async function returns a value, that value will also get wrapped in a promise. That means you’ll have to use `.then` or `await` to access it.
