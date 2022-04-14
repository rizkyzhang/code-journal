- Since JSX is closer to JavaScript than to HTML, React DOM uses camelCase property naming convention instead of HTML attribute names. For example, class becomes `className` in JSX, and tabindex becomes `tabIndex`.

## JSX Represents Objects

Babel compiles JSX down to React.createElement() calls.

These two examples are identical:

```js
const element = <h1 className="greeting">Hello, world!</h1>;
```

```js
const element = React.createElement(
  "h1",
  { className: "greeting" },
  "Hello, world!"
);
```

## Element VS Component

- React element = an object that represent a HTML element.
- React component = a custom element that are made of multiple React elements.

React element:

```js
const element = <h1>I am a React element</h1>;
```

React component:

```js
const Profile = () => {
  return (
    <div>
      <h1>My name is John</h1>
      <p>I have 3 hobbies: </p>
      <ul>
        <li>Coding</li>
        <li>Hiking</li>
        <li>Swimming</li>
      </ul>
    </div>
  );
};
```

## Updating the Rendered Element

React elements are immutable. Once you create an element, you can’t change its children or attributes. An element is like a single frame in a movie: it represents the UI at a certain point in time.
