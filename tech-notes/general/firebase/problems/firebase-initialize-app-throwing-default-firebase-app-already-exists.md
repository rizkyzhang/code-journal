---
tags:
  - firebase-problem
---
## Problem

`Firebase.initializeApp throwing default Firebase app already exists`
This issue start to happen in next.js 13.

## Solution 1 

```ts
// provide unique app name with random id

import { AppOptions, auth, credential } from "firebase-admin";
import { initializeApp } from "firebase-admin/app";
import { getAuth } from "firebase-admin/auth";
import { createId } from "@paralleldrive/cuid2";

const adminApp = initializeApp(firebaseAdminConfig, `admin-${createId()}`);
const beAuth = getAuth(adminApp);

export default beAuth;
```

## Solution 2

```ts
import { AppOptions, auth, credential } from "firebase-admin";
import { initializeApp, getApps } from "firebase-admin/app";
import { getAuth } from "firebase-admin/auth";
import { createId } from "@paralleldrive/cuid2";

const adminApp = getApps().length <= 0 ? initializeApp(firebaseAdminConfig) : getApps()[0];
const beAuth = getAuth(adminApp);

export default beAuth;
```

```ts
import { FirebaseOptions, initializeApp, getApps } from "firebase/app";
import { getAuth } from "firebase/auth";

const app = getApps().length > 0 ? getApp() : initializeApp(firebaseConfig);
const feAuth = getAuth(app);

export default feAuth
```

## Reference

- https://github.com/firebase/firebase-admin-node/issues/2111
- https://blog.stackademic.com/setting-up-firebase-authentication-with-next-13-app-router-using-server-components-03fbcab254e4