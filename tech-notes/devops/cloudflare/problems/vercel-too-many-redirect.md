## Problem

Vercel too many redirect

## Solution

Update CloudFlare TLS settings. To make sure that you don’t get stuck in redirection loops, use either the Full or Full (Strict) setting so that all the web requests are only made using the HTTPS connection.

Reference: https://blog.runcloud.io/fixing-redirect-loop-on-cloudflare-ssl/
