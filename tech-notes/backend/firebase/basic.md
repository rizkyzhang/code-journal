# Basic

## firebase vs firebase-tools vs firebase-admin

- The firebase NPM package contains the Firebase SDK for client-side JavaScript, so that you can use Firebase inside your app.
- The firebase-tools NPM package contains the Firebase command-line interface, so that you can call commands like firebase deploy or use the Firebase emulator suite on your machine.
- The firebase-admin NPM package contains the Firebase SDK for Node.js on trusted environments (such as your development machine, a server you control, or Cloud Functions). This allows you to run Node.js scripts that access you Firebase project with administrative privileges.

Reference: https://stackoverflow.com/questions/68324224/difference-between-firebase-vs-firebase-tools-vs-firebase-admin#:~:text=The%20firebase%2Dtools%20NPM%20package,the%20Firebase%20SDK%20for%20Node.

## Firebase auth token

Firebase Authentication tokens expire an hours after they are created. Firebase SDKs automatically refresh the token after about 55 minutes, which means you usually don't have to do anything yourself. The current user will only become null if the token can't be refreshed, for example if the account has been disabled on the server.

Reference: https://stackoverflow.com/a/62390404

## Firabase link multiple providers to one account

In some situations, Firebase will automatically link accounts when a user signs in with different providers using the same email address. This can only happen when specific criteria are met, however. To understand why, consider the following situation: a user signs in using Google with a @gmail.com account and a malicious actor creates an account using the same @gmail.com address, but signing in via Facebook. If these two accounts were automatically linked, the malicious actor would gain access to the user's account.

The following cases describe when we automatically link accounts and when we throw an error requiring user or developer action:

- User signs in with an untrusted provider, then signs in with another untrusted provider with the same email (for example, Facebook followed by GitHub). This throws an error requiring account linking.
- User signs in with a trusted provider, then signs in with untrusted provider with the same email (for example, Google followed by Facebook). This throws an error requiring account linking.
- User signs in with an untrusted provider, then signs in with a trusted provider with the same email (for example, Facebook followed by Google). The trusted provider overwrites the untrusted provider. If the user attempts to sign in again with Facebook, it will cause an error requiring account linking.
- User signs in with a trusted provider, then signs in with a different trusted provider with the same email (for example, Apple followed by Google). Both providers will be linked without errors.

You can manually set an email as verified by using the Admin SDK, but we recommend only doing this if you know the user really does own the email.

### Why not turn off One account per email address

By default, the setting One account per email address is active. But, unchecking this means that there are two different accounts (User UIDs) for the same email. One via Email/Password and one via Google Authentication. They will have separate User UIDs in Firebase Auth, so that may confuse you. Furthermore, if you manually link in your app two User UIDs, this creates a security hole: Someone can create an account without email verification to get access to an existing account. So don't do that.

Reference: https://stackoverflow.com/a/70195106
