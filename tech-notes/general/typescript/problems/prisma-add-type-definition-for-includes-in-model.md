---
tags:
  - typescript-problem
  - typescript-problem-prisma
---

## Problem

Add type definitions for includes in a Prisma model

```ts
const getUser = await prisma.user.findUnique({
  where: {
    id: 1,
  },
  include: {
    posts: {
      select: {
        title: true,
      },
    },
  },
});
```

When I want to read the property `getUser.posts` I get the following error:

`TS2339: Property 'posts' does not exist on type 'User'.`

## Solution

The generated types do not include relations because queries don't return relations by default. To include the related models in your type, use the provided Prisma utility types like so:

```ts
import { Prisma } from "@prisma/client";

type UserWithPosts = Prisma.UserGetPayload<{
  include: { posts: true };
}>;
```

## Reference

https://stackoverflow.com/questions/71442989/how-to-add-type-definitions-for-includes-in-a-prisma-model
