---
tags:
  - next-js-problem
---
## Solution

In a Next.js 13 project, you can send a cookie to an API route by including it in the headers of your Axios request in your server-side function. In the API route, you can access this cookie using `request.cookies.get('cookieName')`.

```tsx
// middleware.ts

const url = `${process.env.NEXT_PUBLIC_API_BASE_URL}api/post/list`;
const { data } = await axios.get(url, {
    headers: {
        'Cookie': `jwt=${jwt};`
    },
});
```

## Reference

https://gist.github.com/mustafadalga/fde00871063edfa601dfa7497f671502