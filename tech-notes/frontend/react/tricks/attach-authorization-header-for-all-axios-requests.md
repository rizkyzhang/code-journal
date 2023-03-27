Run this function in useEffect

```jsx
import axios from "axios";

export default function setAuthToken(token) {
  axios.defaults.headers.common["Authorization"] = "";
  delete axios.defaults.headers.common["Authorization"];

  if (token) {
    axios.defaults.headers.common["Authorization"] = `${token}`;
  }
}
```

Reference: https://stackoverflow.com/a/44798009
