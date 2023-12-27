---
tags:
  - firebase-problem
---
## Problem

`sendSignInLinkToEmail throws auth/argument-error`

## Solution

`handleCodeInApp` should always be set to true when doing `signInWithEmailLink`.

```tsx
sendSignInLinkToEmail(auth, email, {
	url: "http://localhost:3000/signin-verify",
	handleCodeInApp: true,
})
```

## Reference

https://github.com/firebase/firebase-js-sdk/issues/7101