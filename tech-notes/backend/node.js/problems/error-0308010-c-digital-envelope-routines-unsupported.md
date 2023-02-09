## Problem

Error message `"error:0308010C:digital envelope routines::unsupported"`

## Solution 1

You can try one of these:

1. Downgrade to Node.js v16.

You can reinstall the current LTS version from Node.js’ website.

You can also use nvm. For Windows, use nvm-windows.

2. Enable legacy OpenSSL provider.

On Unix-like (Linux, macOS, Git bash, etc.):
`export NODE_OPTIONS=--openssl-legacy-provider`

On Windows command prompt:
`set NODE_OPTIONS=--openssl-legacy-provider`

On PowerShell:
`$env:NODE_OPTIONS = "--openssl-legacy-provider"`

## Solution 2

### Danger

This question has more than 30 answers, most suggesting to either downgrade Node.js to pre v17 or to use the legacy SSL provider. Both of those solutions are hacks that leave your builds open to security threats.

### Reason For The Error

In Node.js v17, the Node.js developers closed a security hole in the SSL provider. This fix was a breaking change that corresponded with similar breaking changes in the SSL packages in NPM. When you attempt to use SSL in Node.js v17 or later without also upgrading those SSL packages in your package.json, then you will see this error.

### The Correct (safe) Solution (for npm users)

Use an up-to-date version of Node.js, and also use packages that are up-to-date with security fixes.

For many people, the following command will fix the issue:

`npm audit fix --force`
However, be aware that, for complex builds, the above command will pull in breaking security fixes that can potentially break your build.

### Note for Yarn users

As some commenter have pointed out, if you use Yarn rather than Npm, then the above npm based fix will likely not suit you or may yield incomplete results. I do not use Yarn, so can not help with that; however, the concept should be the same.

### A less heavy-handed (also correct) solution for Webpack

In your Webpack config, set either of the following: (See the output.hashFunction docs)

A. (Webpack v5) Set output.hashFunction = 'xxhash64'.
B. (Webpack v4) This will depend on what hash algorithms nodejs supports on your system. Some common options you can try are output.hashFunction = 'sha512' or output.hashFunction = 'sha256'.

## Reference

https://stackoverflow.com/questions/69692842/error-message-error0308010cdigital-envelope-routinesunsupported
