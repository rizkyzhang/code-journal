## Double click event with event.detail

```js
import React from "react";

export default function App() {
  const handleClick = (e) => {
    switch (e.detail) {
      case 1:
        console.log("click");
        break;
      case 2:
        console.log("double click");
        break;
      case 3:
        console.log("triple click");
        break;
      default:
        return;
    }
  };

  return <button onClick={handleClick}>Click me</button>;
}
```
