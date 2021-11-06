# Basic Hooks

- Don't put hooks inside conditional statement, React will throw an error, instead put conditional statement inside hooks.

## useState

- useState is used to preserve state of an application. Let's say we have a counter, everytime we click increase or decrease or reset button, the count value (stored in variable count) will change, but the weird thing is the UI won't change. That's happen because the way React works, **it doesn't notice there is a change unless the value is stored inside the state**, storing the value inside the state will ensure React remember the value and know when the value changes. When the state changes, React will run `render` method which will re-render the UI.

## useEffect

- If the depedency list (second argument) of useEffect is not specified, it will run after every re-render.
- If the depedency list (second argument) of useEffect is an empty array, it will only run on initial render.
- If the depedency list (second argument) of useEffect is an array of specified states, it will run every time the specified states changes.
