---
tags:
  - firebase-problem
---
## Problem

`TypeError: Cannot read property 'INTERNAL' of undefined at FirebaseNamespace.initializeApp`

## Solution

Change `import { initializeApp } from 'firebase-admin';` to `import { initializeApp } from 'firebase-admin/app';`

## Reference

https://github.com/firebase/firebase-admin-node/issues/593#issuecomment-1703977020
`