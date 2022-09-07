## Define types for process.env

1. Create an environment.d.ts file inside `src/types/environment` and declare types in the global namespace.
2. Add properties to the ProcessEnv interface.
3. Make the file a module by using export {}.
4. Prepend `"./src/types"` to `typeRoots` property in tsconfig.json

```
src/
  - types/
    - environment/
      - environment.d.ts
```

```ts
declare global {
  namespace NodeJS {
    interface ProcessEnv {
      ACCESS_TOKEN_SECRET: string;
      REFRESH_TOKEN_SECRET: string;
    }
  }
}

export {};
```

We used the `export {}` line in our `environment.d.ts` file to mark it as an external module. A module is a file that contains at least 1 import or export statement. We are required to do that to be able to augment the global scope.

```json
{
  "compilerOptions": {
    "typeRoots": ["./src/types", "./node_modules/@types"]
  }
}
```

## Reference

https://bobbyhadz.com/blog/typescript-process-env-type
