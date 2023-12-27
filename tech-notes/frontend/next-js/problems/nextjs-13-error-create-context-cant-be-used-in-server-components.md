---
tags:
  - next-js-problem
---
## Problem

`Error: createContext can't be used in Server Components`

## Solution

Instead of

```tsx
<AuthProvider auth={feAuth}>
	<html lang="en">
		<body className={inter.className}>{children}</body>
	</html>
</AuthProvider>
```

Create a function that wrap children with the provider

```tsx
export const AuthProviderWrapper = (props: { children: ReactNode }) => {
	return <AuthProvider auth={feAuth}>{props.children}</AuthProvider>;
};
```

```tsx
<AuthProviderWrapper>
	<html lang="en">
		<body className={inter.className}>{children}</body>
	</html>
</AuthProviderWrapper>
```

## Reference

https://dev.to/codingbrowny/using-context-providers-in-nextjs-server-components-2gk4