## Problem

Error: Channel closed by server: 406 (PRECONDITION-FAILED) with message "PRECONDITION_FAILED - unknown delivery tag 1"

## Solution

ACK rules:

- You have to ACK messages in same order as they arrive to your system
- You can't ACK messages on a different channel than that they arrive on

If you break any of these rules you will face 406 (PRECONDITION-FAILED) error message.

## Reference

https://www.grzegorowski.com/rabbitmq-406-channel-closed-precondition-failed
