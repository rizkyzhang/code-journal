# Webpack serve error

```
`[webpack-cli] D:\lifeRestart\my-project\node_modules\webpack-dev-server\lib\servers\WebsocketServer.js:10
static heartbeatInterval = 1000;
^

SyntaxError: Unexpected token =
at new Script (vm.js:83:7)
at NativeCompileCache._moduleCompile (D:\lifeRestart\my-project\node_modules\v8-compile-cache\v8-compile-cache.js:240:18)
at Module._compile (D:\lifeRestart\my-project\node_modules\v8-compile-cache\v8-compile-cache.js:184:36)
at Object.Module._extensions..js (internal/modules/cjs/loader.js:789:10)
at Module.load (internal/modules/cjs/loader.js:653:32)
at tryModuleLoad (internal/modules/cjs/loader.js:593:12)
at Function.Module._load (internal/modules/cjs/loader.js:585:3)
at Module.require (internal/modules/cjs/loader.js:692:17)
at require (D:\lifeRestart\my-project\node_modules\v8-compile-cache\v8-compile-cache.js:159:20)
at Server.getServerTransport (D:\lifeRestart\my-project\node_modules\webpack-dev-server\lib\Server.js:982:28)`
```

Causes: the static variables syntax is available in Node 6, but the static class variable that's used here is available in Node 12.6.

Solution: update nodejs to latest LTS version.

Reference: https://github.com/VickScarlet/lifeRestart/issues/176
