## Terminology

- Compiling = The process of transforming a code written in a language into another language or different version of that language. For example code written in TypeScript, need to be compiled into JavaScript before being parsed by the browser so the browser understand it.

- Minifying = The process of removing unnecessary code formatting and comments without changing the code’s functionality. The goal is to improve the application’s performance by decreasing file sizes.

- Bundling = The process of resolving the web of dependencies and merging (or ‘packaging’) the files (or modules) into optimized bundles for the browser, with the goal of reducing the number of requests for files when a user visits a web page.

- Code splitting = The process of splitting the application’s bundle into smaller chunks required by each entry point. Next.js has built-in support for code splitting. Each file inside your `pages/` directory will be automatically code split into its own JavaScript bundle during the build step.

- Build time (or build step) is the name given to a series of steps that prepare your application code for production.

- Runtime (or request time) refers to the period of time when your application runs in response to a user’s request, after your application has been built and deployed.
