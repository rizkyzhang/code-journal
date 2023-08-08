1. Create NEXT_PUBLIC_SITE_URL env to access the URL from server side

```
# development
NEXT_PUBLIC_SITE_URL=http://localhost:3000

# production
NEXT_PUBLIC_SITE_URL=https://yourdomain.com
```

2. Create the util

```ts
const IS_SERVER = typeof window === "undefined";
export default function getURL(path: string) {
  const baseURL = IS_SERVER
    ? process.env.NEXT_PUBLIC_SITE_URL
    : window.location.origin;
  return new URL(path, baseURL).toString();
}
```

## Reference

https://reacthustle.com/blog/how-to-get-the-current-domain-url-in-nextjs-server-side-and-client-side
