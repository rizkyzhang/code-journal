## Problem

`Error [ERR_HTTP_HEADERS_SENT]: Cannot set headers after they are sent to the client`

## Solution

Use `return` statement for `res.json`, `res.status` and `res.send` provided the if else logic

## Reference

https://stackoverflow.com/questions/70912775/error-err-http-headers-sent-cannot-set-headers-after-they-are-sent-to-the-cli
