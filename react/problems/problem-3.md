Problem:

```
Error: Invalid hook call. Hooks can only be called inside of the body of a function component. This could happen for one of the following reasons:
1. You might have mismatching versions of React and the renderer (such as React DOM)
2. You might be breaking the Rules of Hooks
3. You might have more than one copy of React in the same app
See https://reactjs.org/link/invalid-hook-call for tips about how to debug and fix this problem.
```

Causes: Accidently passing component as function to `Route`

```html
<Route exact path="*">{Home}</Route>
```

Solution:

```html
<Route exact path="*"><Home /></Route>
```

```
<Route exact path="*" component={Home}></Route>
```
