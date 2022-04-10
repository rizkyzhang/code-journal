## Problem

Error `ReferenceError: exports is not defined in ES module scope This file is being treated as an ES module because it has a '.js' file extension and '/home/test/package.json' contains "type": "module". To treat it as a CommonJS script, rename it to use the '.cjs' file extension.` when running `node dist/index.js` after `npm run build`.

## Solution

package.json

```json
{
  "type": "module" // 👈️ remove this
}
```

When type is set to module we are not able to use the exports and require syntax of CommonJS anymore and have to stick to ES Modules syntax for all of our imports and exports, which might be causing the error.

This happens when you have type set to module in your package.json file, but have module set to commonjs in your tsconfig.json file.

Open your tsconfig.json file and make sure it looks something like the following.

```json
{
  "compilerOptions": {
    "target": "es6",
    "module": "commonjs",
    "esModuleInterop": true,
    "moduleResolution": "node"
    // ... your other options
  }
}
```

## Reference

https://bobbyhadz.com/blog/typescript-uncaught-referenceerror-exports-is-not-defined
