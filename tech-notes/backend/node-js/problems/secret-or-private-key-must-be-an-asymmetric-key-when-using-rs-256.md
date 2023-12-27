## Problem

`secretOrPrivateKey must be an asymmetric key when using RS256`

## Solution

If you are using `pem` string as a public key directly, use `jwk` public key instead and convert it to pem with `jwk-to-pem` library.

JWK object format example:

```json
{
  "kty": "RSA",
  "n": "9CjjhezNt2nxiLe1SjW7zQC03U6DU1
bdnL1HzvWHYf3VdXFzhUgvF2jjmmZq
jF88L2CzneqRp1JcyunpDBjbSb3qWK
UiER2cBiUipD1xtEXbcFWGVh2J66ZR
Mnnv2xYKBNjz3xCyqPuDqcbkaSEMyy
hwtFeLYjTVyMuDaU0E8hSERxGSXfXu
bWyG36JLVC3JvCx6YfivTXn3CCFy52
nhdp6pkekRHy275CrtYMPYAUXGXCAN
MNzFYLScwchuBr76mQgiRMMRmzcmGp
8KMHNVBtzhkUmeyHdq3UxvYAwz0Eni",
  "e": "AQAB"
}
```

Library usage if jwk stored as env:

```ts
const jwk = process.env.JWK_PUBLIC_KEY || "";
const pem = jwkToPem(JSON.parse(jwk));
```
