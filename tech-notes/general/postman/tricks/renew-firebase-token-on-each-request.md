Put this in your collection's pre request section

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
  pm.request.headers.add({ key: "FirebaseUID", value: "Your UID" });
});
```
