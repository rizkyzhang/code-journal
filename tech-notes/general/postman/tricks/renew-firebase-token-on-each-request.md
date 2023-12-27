---
tags:
  - postman-trick
---
Put one of this snippet into your collection's pre-request script

```js
const postRequest = {
  url: "https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyPassword?key=your-api-key",
  method: "POST",
  header: {
    "Content-Type": "application/json",
  },
  body: {
    mode: "raw",
    raw: JSON.stringify({
      email: "your email",
      password: "your password",
      returnSecureToken: true,
    }),
  },
};

pm.sendRequest(postRequest, (error, response) => {
  var jsonData = response.json();
  pm.request.headers.add({
    key: "Authorization",
    value: `Bearer ${jsonData.idToken}`,
  });
});
```

```js
const postRequest = {
	url: 'https://securetoken.googleapis.com/v1/token?key=your-api-key',
	method: 'POST',
	header: {		
		'Content-Type': 'application/x-www-form-urlencoded',
	},
	body: {
		mode: 'urlencoded',
		urlencoded : [
			{key:'grant_type', value:'refresh_token'},
			{key:'refresh_token', value:'your-refresh-token'}
		]
	}
};

pm.sendRequest(postRequest, (error, response) => {
	var jsonData = response.json();
	pm.request.headers.add({
		key: 'Authorization', 
		value: `Bearer ${jsonData.id_token}`
	})
});
```