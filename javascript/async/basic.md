## Synchronous vs Asynchronous

- By default, JavaScript is a synchronous, blocking, single-threaded language, in which only one operation can be in progress at a time.
- Synchronous = operations are executed in sequence/order (only one operation at a time).
- Asynchronous = operations are executed simultaneously (multiple operations at a time).

### Real life examples

#### **Synchronous**

- Images that load in sequence, first image load then second image and then third image. Next image have to wait until previous image have completely loaded.
- Runners are in the same lane. One can't overtake the other. The race is finished one by one. If one stops, the following runners stops.

#### **Asynchronous**

- Images that load simultaneously.
- Runners are in different lanes. They'll finish the race on their own pace. Nobody stops for anybody.

### Code examples

#### **Synchronous**

```js
console.log("First");
console.log("Second");
console.log("Third");
```

```js
console.log("First");

// This will be shown after 2 seconds

setTimeout(() => {
  console.log("Second");
}, 2000);

console.log("Third");
```
