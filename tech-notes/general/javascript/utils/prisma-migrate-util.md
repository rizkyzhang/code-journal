---
tags:
  - js-util
---
```ts
import { promisify } from "node:util";
import { exec as execCb } from "node:child_process";

const exec = promisify(execCb);

// TODO: re-write this when Prisma.io gets a programmatic migration API
// https://github.com/prisma/prisma/issues/4703
export async function prismaMigrate(databaseUrl: string): Promise<void> {
  // throws an error if migration fails
  const { stdout, stderr } = await exec("npx prisma migrate dev", {
    env: {
      ...process.env,
      DATABASE_URL: databaseUrl,
    },
  });
  // console.log(stdout)
  // console.log(stderr)
}
```

## Reference

https://github.com/prisma/prisma/issues/4703#issuecomment-1447354363