- TypeScript is an open source language and is a superset of JavaScript.
- **Using types is optional**.
- Compiled down to JavaScript.
- Types from 3rd parties can be added with type definitions.

## Dynamic VS Static Typing

- In dynamically typed languages, the types are associated with run-time values and not named explicitly.
- In statically typed languages, you explicitly assign types to variables, function parameters, return values, etc.
- Static languages: Java, C, C++, Rust, Go
- Dynamic langauges: Python, JavaScript, PHP, Ruby

## Compiling TypeScript

- TypeScript uses .ts and .tsx ectensions.
- TSC (TypeScript Compiler) is used to compile .ts files down to JS. It can also watch files and report errors at compile time.
- The tsconfig.json file is used to configure how TypeScript works.

## Setup

- Install tsc as devDepedency `npm i -D typescript`
- Create tsconfig.json `npx tsc --init`
- Run tsc `npx tsc`

## Setup React

- `npx create-react-app project --template typescript`
