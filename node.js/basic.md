# Basic

- Runtime = environment in whiich JavaScript run.
- Node.js is created by Ryan Dahl in 2009.
- Node.js allows us to run JavaScript outside of browser.
- REPL = CLI of Node.js

- Waiter = thread
- Customer = task
- Kitchen = node.js

## Multi thread synchronous

- Multiple waiter is assigned to one customer each.

## Single thread asynchronous

- Single waiter can managed multiple customer at a time.

- In browser, JavaScript variables and functions are tied to a global object named window.
- In node.js, JavaScript variables and functions are tied to a global object named global.
- In browser, 2 JavaScript files in a same html file can access each other codes, so name collision could happen as if it is a single file. This did not happen in node.js thought, because it uses a module system which only allow a JavaScript file to be accessed by another JavaScript file if there is a explicit exports statement in the file.

## Type of modules

- Core modules
- Local modules
- Third party modules

## Express

- Middleware is a function that were called between processing the request and sending the response in an application.
- Middleware is usually used for authentication, error checking and parsing request body.
- You **NEED** `express.json()` and `express.urlencoded()` for `POST` and `PUT` requests, because in both these requests you are sending data (in the form of json) to the server and it needs to be parsed first so server can read it.
