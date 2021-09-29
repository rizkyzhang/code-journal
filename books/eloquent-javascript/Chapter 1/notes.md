# Values, Types, and Operators

Inside the computer’s world, there is only data. You can read data, modify data, create new data—but that which isn’t Values, Types, and Operators data cannot be mentioned. All this data is stored as long sequences of bits and is thus fundamentally alike.

Bits are any kind of two-valued things, usually described as zeros and ones. Inside the computer, they take forms such as a high or low electrical charge, a strong or weak signal, or a shiny or dull spot on the surface of a CD. Any piece of discrete information can be reduced to a sequence of zeros and ones and thus represented in bits.

## Special numbers

There are three special values in JavaScript that are considered numbers but don’t behave like normal numbers.

The first two are Infinity and -Infinity, which represent the positive and negative infinities. Infinity - 1 is still Infinity, and so on. Don’t put too much trust in infinity-based computation, though. It isn’t mathematically sound, and it will quickly lead to the next special number: NaN.

NaN stands for “not a number”, even though it is a value of the number type. You’ll get this result when you, for example, try to calculate 0 / 0 (zero divided by zero), Infinity - Infinity, or any number of other numeric operations that don’t yield a meaningful result.

## Unary Operators

Operators that use two values are called **binary operators**, while those that take one are called **unary operators**. The minus and plus operator can be used both as a binary operator and as a unary operator.

There is only one value in JavaScript that is not equal to itself, and that is NaN (“not a number”).

## Comparison

```js
console.log(NaN == NaN);
// → false
```

NaN is supposed to denote the result of a nonsensical computation, and as such, it isn’t equal to the result of any other nonsensical computations
