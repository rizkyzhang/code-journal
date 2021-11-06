Problem: TypeError: Cannot read properties of undefined (reading ‘location‘)

Causes:

```js
import { Router } from "react-router-dom";
```

Solution:

```js
import { BrowserRouter as Router } from "react-router-dom";
```
