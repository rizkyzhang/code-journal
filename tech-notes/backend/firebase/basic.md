# Basic

## firebase vs firebase-tools vs firebase-admin

- The firebase NPM package contains the Firebase SDK for client-side JavaScript, so that you can use Firebase inside your app.
- The firebase-tools NPM package contains the Firebase command-line interface, so that you can call commands like firebase deploy or use the Firebase emulator suite on your machine.
- The firebase-admin NPM package contains the Firebase SDK for Node.js on trusted environments (such as your development machine, a server you control, or Cloud Functions). This allows you to run Node.js scripts that access you Firebase project with administrative privileges.

Reference: https://stackoverflow.com/questions/68324224/difference-between-firebase-vs-firebase-tools-vs-firebase-admin#:~:text=The%20firebase%2Dtools%20NPM%20package,the%20Firebase%20SDK%20for%20Node.
