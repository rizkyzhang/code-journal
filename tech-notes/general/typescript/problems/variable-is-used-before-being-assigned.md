## Problem

`Variable 'X' is used before being assigned`

## Cause

The error "Variable is used before being assigned" occurs when we declare a variable without assigning a value to it or only assign a value if a condition is met.

To solve the error, change the variable's type to possibly `undefined` or give it an initial value.

Here are 3 examples of how the error occurs.

```ts
type Employee = {
  name: string;
  country: string;
};

let employee: Employee; // 👈️ did not assign value to variable

// 1. ⛔️ Error: Variable 'employee' is used before being assigned.ts(2454)
employee.name = "bobbyhadz.com";

// ------------------------------------------

let arr: string[]; // 👈️ did not assign value to variable

// 2. ⛔️ Error: Variable 'arr' is used before being assigned.ts(2454)
arr[0] = "bobbyhadz.com";

// ------------------------------------------

// 3. Only assigns value to the variable if a condition is met 👇️
let salary: number;

if (Math.random() > 0.5) {
  salary = 100;
}

// ⛔️ Error: Variable 'salary' is used before being assigned.ts(2454)
if (salary === 100) {
  console.log("success");
}
```

The first two examples cause the error because:

1. We declare a variable and set its type.
2. We don't give an initial value to the variable.
3. We try to use the variable.

## Solution 1: Initialize the variable to solve the error

```ts
type Employee = {
  name: string;
  country: string;
};

const employee: Employee = {
  name: "",
  country: "",
};

employee.name = "bobby hadz";

// ------------------------------------

const arr: string[] = [];

arr[0] = "bobbyhadz.com";
```

## Solution 2: Update the variable's type to be possible undefined

```ts
type Employee = {
  name: string;
  country: string;
};

let employee: Employee | undefined; // 👈️ could be undefined

// 👉️ somewhere in your code this should be set to an object

if (employee !== undefined) {
  employee.name = "Bobby Hadz";
}

// ---------------------

let arr: string[] | undefined; // 👈️ could be undefined

// 👉️ somewhere in your code this should be set to an array

if (arr !== undefined) {
  arr[0] = "bobbyhadz.com";
}
```

## Solution 3: Using the Partial utility type to solve the error

```ts
type Employee = {
  name: string;
  country: string;
};

const employee: Partial<Employee> = {};

employee.name = "Bobby Hadz";
employee.country = "Brazil";

// 👇️ {name: 'Bobby Hadz', country: 'Brazil'}
console.log(employee);
```

Reference: https://bobbyhadz.com/blog/typescript-variable-is-used-before-being-assigned#:~:text=The%20error%20%22Variable%20is%20used,give%20it%20an%20initial%20value.
