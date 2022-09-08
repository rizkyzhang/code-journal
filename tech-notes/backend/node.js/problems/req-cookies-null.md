## Problem

req.cookies null

## Solution

If you are in localhost and don't have a ssl/tls setup currently, maybe try putting the `secure: false`, since `secure` is only used to make sure the data which is being transferred is actually encrypted using the HTTPS protocol.

or if you want to you can also set it dynamically as:
`secure: process.env.NODE_ENV !== "development"` // note that it will only happen if you have the environment variable `NODE_ENV `set properly.

## Reference

https://github.com/expressjs/express/issues/4924#issuecomment-1185274752
