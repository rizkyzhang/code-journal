## Problem

`ERROR: Test failed: 403 (SignatureDoesNotMatch)` when running `s3cmd --configure`

## Solution

If CLI asked to retry, choose no and choose save to config file, run `s3cmd ls`, it should be fixed.
