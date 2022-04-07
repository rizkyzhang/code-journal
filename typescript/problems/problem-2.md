## Problem:

```ts
type User =
  | {
      name: string;
      email: string;
    }
  | {};
```

So essentially a user can be an object with name and email properties, or the empty object.

```js
if (user.name) {
  // ERROR
  console.log(user.name);
}
```

When trying to access name property from user object, TypeScript throw an error: `Property 'name' does not exist on type '{}'`.

## Solution:

Using in operator https://www.typescriptlang.org/docs/handbook/advanced-types.html#using-the-in-operator

```js
if (user.name) {
  // error
  user; // typeof user is User | {}
  user.name; // error
  user.email; // error
}
```

```js
if ("name" in user) {
  user; // ok, typeof user is User
  user.name; // string
  user.email; // string
}
```

Reference: https://stackoverflow.com/questions/52708347/typescript-discriminated-union-with-empty-object
