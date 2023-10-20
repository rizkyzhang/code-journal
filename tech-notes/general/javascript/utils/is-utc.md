```ts
import { DateTime } from "luxon";

const isUTC = (input: string) => {
  const dt = DateTime.fromISO(input, { setZone: true });
  return dt.offset === 0;
};
```

Reference: https://stackoverflow.com/questions/72896281/check-if-iso-8601-string-is-in-utc
