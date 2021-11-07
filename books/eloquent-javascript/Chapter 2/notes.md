# Program Structure

## Expressions and statements

**A fragment of code that produces a value is called an expression**. Every value that is written literally (such as 22 or "psychoanalysis") is an expression. An expression between parentheses is also an expression, as is a binary operator applied to two expressions or a unary operator applied to one.

This shows part of the beauty of a language-based interface. Expressions can contain other expressions in a way similar to how subsentences in human languages are nested—a subsentence can contain its own subsentences, and so on. This allows us to build expressions that describe arbitrarily complex computations.

**If an expression corresponds to a sentence fragment, a JavaScript statement corresponds to a full sentence. A program is a list of statements**.

The simplest kind of statement is an expression with a semicolon after it. This is a program:

```js
1;
!false;
```

It is a useless program, though. An expression can be content to just produce a value, which can then be used by the enclosing code. A statement stands on its own, so it amounts to something only if it affects the world. It could display something on the screen—that counts as changing the world—or it could change the internal state of the machine in a way that will affect the statements that come after it. These changes are called side effects. The statements in the previous example just produce the values 1 and true and then immediately throw them away. This leaves no impression on the world at all. When you run this program, nothing observable happens.

## Bindings

How does a program keep an internal state? How does it remember things? We have seen how to produce new values from old values, but this does not change the old values, and the new value has to be immediately used or it will dissipate again. To catch and hold values, JavaScript provides a thing called a binding, or variable:

```js
let caught = 5 * 5;
```

That’s a second kind of statement. The special word (keyword) let indicates that this sentence is going to define a binding. It is followed by the name of the binding and, if we want to immediately give it a value, by an = operator and an expression.

The previous statement creates a binding called caught and uses it to grab hold of the number that is produced by multiplying 5 by 5.

When a binding points at a value, that does not mean it is tied to that value forever. The = operator can be used at any time on existing bindings to disconnect them from their current value and have them point to a new one.

```js
let mood = "light";
console.log(mood);
// → light
mood = "dark";
console.log(mood);
// → dark
```

**You should imagine bindings as tentacles, rather than boxes. They do not contain values; they grasp them—two bindings can refer to the same value**. A program can access only the values that it still has a reference to. When you need to remember something, you grow a tentacle to hold on to it or you reattach one of your existing tentacles to it.

## Functions

Executing a function is called _invoking_, _calling_, or _applying_ it. You can call a function by putting parentheses after an expression that produces a function value. Usually you’ll directly use the name of the binding that holds the function. The values between the parentheses are given to the program inside the function. In the example, the prompt function uses the string that we give it as the text to show in the dialog box. Values given to functions are called _arguments_.

## Return values

Showing a dialog box or writing text to the screen is a _side effect_. A lot of functions are useful because of the side effects they produce. Functions may also produce values, in which case they don’t need to have a side effect to be useful. For example, the function `Math.max` takes any amount of number arguments and gives back the greatest.

```js
console.log(Math.max(2, 4));
// → 4
```

When a function produces a value, it is said to return that value. Anything that produces a value is an expression in JavaScript, which means function calls can be used within larger expressions. Here a call to `Math.min`, which is the opposite of Math.max, is used as part of a plus expression:

```js
console.log(Math.min(2, 4) + 100);
// → 102
```
