# Basic Hooks

- Don't put hooks inside conditional statement, React will throw an error, instead put conditional statement inside hooks.

## useState

- useState is used to preserve state of an application. Let's say we have a counter, everytime we click increase or decrease or reset button, the count value (stored in variable count) will change, but the weird thing is the UI won't change. That's happen because the way React works, **it doesn't notice there is a change unless the value is stored inside the state**, storing the value inside the state will ensure React remember the value and know when the value changes. When the state changes, React will run `render` method which will re-render the UI.

## useState vs useReducer

- useState when managing only one state.
- useReducer when managing multiple state that depends on each other to reduce human error (accidently set wrong state).

## useState vs useRef vs normal variable

- useState is immutable, useRef is mutable via `current` property.
- useState cause re-render when state changes, useRef doesn't cause re-render when state changes.
- Both useState and useReef keep value between re-render, normal variable doesn't.

## useRef usecases

- Keep value without causing re-rerender (UI doesn't need to be updated)
- Access DOM of a an element/component.

## useEffect

- If the depedency list (second argument) of useEffect is not specified, it will run after every re-render.
- If the depedency list (second argument) of useEffect is an empty array, it will only run on initial render.
- If the depedency list (second argument) of useEffect is an array of specified states, it will run every time the specified states changes.
