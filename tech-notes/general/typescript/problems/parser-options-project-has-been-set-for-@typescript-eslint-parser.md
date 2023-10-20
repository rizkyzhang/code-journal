---
tags:
  - typescript-problem
  - typescript-problem-config
---

## Problem

`"ParserOptions.project" has been set for @typescript-eslint/parser`

## Solution

This error occurs when we try to lint files that aren't included in our TypeScript project. To solve the error, add the files in which you get the error to your `.eslintignore` file. Adding the files that cause the error to your `.eslintignore` file resolves it, because ESLint is no longer trying to parse them.

## Reference

https://bobbyhadz.com/blog/typescript-parseroptions-project-has-been-set-for
