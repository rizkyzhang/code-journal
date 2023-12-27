---
tags:
  - aws-problem
  - typescript-problem
---
## Problem

`TypeError: this.ses.sendRawEmail is not a function`

## Solution

Change import from `import aws from '@aws-sdk/client-ses';` to `import * as aws from '@aws-sdk/client-ses';`

## Reference

https://github.com/nodemailer/nodemailer/issues/1293#issuecomment-878031206