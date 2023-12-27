---
tags:
  - next-js-problem
---
## Problem

Component rendering twice on dev environment without any trigger

## Solution

Add `reactStrictMode: false` into `next.config.js`

## Reference

https://stackoverflow.com/questions/71847778/why-my-nextjs-component-is-rendering-twice