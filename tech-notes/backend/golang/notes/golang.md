- Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

- The functions in the encoding/json standard package will only encode and decode the exported struct fields.
- Two unnamed struct types are identical only if they have the same sequence of field declarations. Two field declarations are identical only if their respective names, their respective types and their respective tags are all identical. **Please note, two non-exported struct field names from different packages are always viewed as two different names**.
