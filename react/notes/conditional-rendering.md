# Conditional rendering

- When passing a event handler function with a parameter into an event attribute, we need to put it inside anonymous arrow function. Otherwise, React will throw us an error "too much re-render" because the event handler will be invoked infinitely which cause infinite re-render.
