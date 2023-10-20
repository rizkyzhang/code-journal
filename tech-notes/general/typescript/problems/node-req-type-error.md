---
tags:
  - typescript-problem
  - typescript-problem-config
  - typescript-problem-express-js
---

## Problem

`Property 'user' does not exist on type 'Request<ParamsDictionary, any, any, ParsedQs, Record<string, any>>'` when doing this: `req.user = user`.

## Solution steps

1. Create a types folder in your `src` directory

2. Create a folder inside the types folder with the name of the package you intend to extend. (In this case `express`).

3. Create an `index.d.ts` file in that folder

```
src/
  - types/
    - express/
      - index.d.ts
```

4. Add this code to the index file

```js
import express from "express";

declare global {
  namespace Express {
    interface Request {
      user?: Record<string,any>
    }
  }
}
```

5. Update your `tsconfig.json` file

```json
{
  "compilerOptions": {
    "typeRoots": ["./src/types", "./node_modules/@types"]
  }
}
```

## Important note

 `typeRoots` property cares about ordering. Make sure you specify your local types before the `node_modules` types! Use `["./types", "node_modules/@types"]` over `["node_modules/@types", "./types"]`.

## Reference

https://stackoverflow.com/questions/65848442/property-user-does-not-exist-on-type-requestparamsdictionary-any-any-pars
