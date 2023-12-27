---
tags:
  - js-util
---
```ts
function assertExists(key: string, value: string | undefined) {
	if (value !== undefined && value !== null) {
		return value;
	}
	
	throw new Error(`ENV ${key} not found or empty`);
}

export default {
  PORT: Number(assertExists("PORT", process.env.PORT)),
};
```

## Reference

https://blog.otaviocapila.dev/nodejs-with-typescript-environment-variables-type-safe