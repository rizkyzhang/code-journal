- Get refresh token from `user.stsTokenManager.refreshToken`

## Expiration

  
- Firebase ID tokens are short lived and last for an hour; the refresh token can be used to retrieve new ID tokens. Refresh tokens expire only when one of the following occurs: **The user is deleted**. **The user is disabled**.
- Password reset `generatePasswordResetLink`: 1 hour  
- Email verification `generateEmailVerificationLink`: 3 days  
- Email link sign-in `generateSignInWithEmailLink`: 6 hours

Reference: https://medium.com/@beard0/firebase-admin-auth-link-expiration-times-9b5b090eae37