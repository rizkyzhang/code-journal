---
tags:
  - aws-problem
---
## Problem

`SignatureDoesNotMatch` error when sending email from AWS SDK.

## Solution

You might be using the access key generated from `AWS SES` dashboard, that is in fact a `SMTP` login/password, not an access key that you can use in the SDK (even if it appears under access keys in `IAM`).

The solution is create a new `IAM` user with [AmazonSESFullAccess](https://docs.aws.amazon.com/aws-managed-policy/latest/reference/AmazonSESFullAccess.html) permission, and obtain its `AccessKeyId` and `SecretAccessKey`.

## Reference

- https://stackoverflow.com/questions/68443435/signaturedoesnotmatch-error-when-sending-an-email-from-node-aws-sdk-simple-emai
- https://www.codu.co/articles/sending-emails-with-aws-ses-and-nodemailer-in-node-js-xfuucrri