# event.currentTarget

The `currentTarget` read-only property of the `Event` interface identifies the current target for the event, as the event traverses the DOM. **It always refers to the element to which the event handler has been attached**, as opposed to `event.target`, which identifies the element on which the event occurred and which may be its descendant.

Note: **The value of `event.currentTarget` is only available while the event is being handled**. If you `console.log()` the `event` object, storing it in a variable, and then look for the `currentTarget` key in the console, its value will be null. Instead, you can either directly `console.log(event.currentTarget)` to be able to view it in the console or use the debugger statement, which will pause the execution of your code thus showing you the value of `event.currentTarget`. **So the value of `event.currentTarget` should be immediately used**.
