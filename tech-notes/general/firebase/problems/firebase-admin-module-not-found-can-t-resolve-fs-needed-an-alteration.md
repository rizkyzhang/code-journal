---
tags:
  - firebase-problem
---
## Problem

`Firebase: Module not found: Can't resolve 'fs' needed an alteration`

## Solution

Firebase Admin SDKs requires full Node.js runtime to work properly. If we are in next.js environment, we should only use it in next.js API route.

## Reference

https://github.com/firebase/firebase-admin-node/issues/2041#issuecomment-1385965747