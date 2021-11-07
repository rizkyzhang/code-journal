People think that computer science is the art of geniuses but the actual reality is the opposite, just many people doing things that build on each other, like a wall of mini stones.

Functions are the bread and butter of JavaScript programming. The concept of wrapping a piece of program in a value has many uses. It gives us a way to structure larger programs, to reduce repetition, to associate names with subprograms, and to isolate these subprograms from each other.

## Defining a function

```js
const makeNoise = function () {
  console.log("Pling!");
};

makeNoise();
// → Pling!

const power = function (base, exponent) {
  let result = 1;
  for (let count = 0; count < exponent; count++) {
    result *= base;
  }
  return result;
};

console.log(power(2, 10));
// → 1024
```

Some functions produce a value, such as `power` and `square`, and some don’t, such as `makeNoise`, whose only result is a side effect. A `return` statement determines the value the function returns. When control comes across such a statement, it immediately jumps out of the current function and gives the returned value to the code that called the function. A `return` keyword without an expression after it will cause the function to return undefined. Functions that don’t have a `return` statement at all, such as makeNoise, similarly return `undefined`.

## Bindings and scopes

Each binding has a scope, which is the part of the program in which the binding is visible. For bindings defined outside of any function or block, the scope is the whole program—you can refer to such bindings wherever you want. These are called global.

But bindings created for function parameters or declared inside a function can be referenced only in that function, so they are known as local bindings. Every time the function is called, new instances of these bindings are created. This provides some isolation between functions—each function call acts in its own little world (its local environment) and can often be understood without knowing a lot about what’s going on in the global environment.

Bindings declared with let and const are in fact local to the block that they are declared in, so if you create one of those inside of a loop, the code before and after the loop cannot “see” it. In pre-2015 JavaScript, only functions created new scopes, so old-style bindings, created with the var keyword, are visible throughout the whole function that they appear in—or throughout the global scope, if they are not in a function.

**Lexical scoping** is a concept in which inner function (inner scope) can access variables from outer function (outer scope). **Closure** is an inner function that can access variables from outer function.

## Optional arguments

JavaScript is extremely broad-minded about the number of arguments you pass to a function. If you pass too many, the extra ones are ignored. If you pass too few, the missing parameters get assigned the value undefined.

## Recursion

Worrying about efficiency can be a distraction. It’s yet another factor that complicates program design, and when you’re doing something that’s already difficult, that extra thing to worry about can be paralyzing.

Therefore, always start by writing something that’s correct and easy to understand. If you’re worried that it’s too slow—which it usually isn’t since most code simply isn’t executed often enough to take any significant amount of time—you can measure afterward and improve it if necessary.

Recursion is not always just an inefficient alternative to looping. Some problems really are easier to solve with recursion than with loops. Most often these are problems that require exploring or processing several “branches”, each of which might branch out again into even more branches.

## Growing functions

A useful principle is to not add cleverness unless you are absolutely sure you’re going to need it. It can be tempting to write general “frameworks” for every bit of functionality you come across. Resist that urge. You won’t get any real work done—you’ll just be writing code that you never use.

## Functions & side effects

A pure function is a specific kind of value-producing function that not only has no side effects but also doesn’t rely on side effects from other code—for example, it doesn’t read global bindings whose value might change. A pure function has the pleasant property that, when called with the same arguments, it always produces the same value (and doesn’t do anything else). A call to such a function can be substituted by its return value without changing the meaning of the code. When you are not sure that a pure function is working correctly, you can test it by simply calling it and know that if it works in that context, it will work in any context.
