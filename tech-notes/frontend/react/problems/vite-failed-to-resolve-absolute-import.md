## Problem

Failed to resolve absolute import with tsconfig already setup

```json
"compilerOptions": {
  "baseUrl": "./src",
  "paths": {
    "@/*": ["*"]
  },
}
```

## Solution

`yarn add -D vite-tsconfig-paths` to give vite the ability to resolve imports using TypeScript's path mapping.

```ts
import react from "@vitejs/plugin-react";
import tsconfigPaths from "vite-tsconfig-paths";

export default {
  plugins: [tsconfigPaths(), react()],
};
```

## Reference

https://stackoverflow.com/questions/75380295/error-in-vitest-config-with-nextjs13-failed-to-resolve-import/75391253#75391253
