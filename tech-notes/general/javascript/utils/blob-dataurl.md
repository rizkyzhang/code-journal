---
tags:
  - js-util
---
```js
const blobToDataURL = (blob) => {
  const reader = new FileReader();
  reader.readAsDataURL(blob);
  return new Promise((resolve) => {
    reader.onloadend = () => {
      resolve(reader.result);
    };
  });
};

const dataURLToBlob = async (dataURL, fileName, fileType) => {
  return await fetch(dataURL)
    .then((res) => res.blob())
    .then((blob) => new File([blob], fileName, { type: fileType }));
};
```

```ts
const getDataURLFromImageURL = (url: string): Promise<string> => {
  return new Promise((resolve) => {
    const img = new Image();
    img.setAttribute("crossOrigin", "anonymous");
    img.src = url;
    img.onload = () => {
      const canvas = document.createElement("canvas");
      canvas.width = img.width;
      canvas.height = img.height;
      const ctx = canvas.getContext("2d");
      ctx?.drawImage(img, 0, 0);
      const dataURL = canvas.toDataURL("image/jpeg");

      resolve(dataURL);
    };
  });
};
```

```js
// canvas to dataURL
const dataURL = canvas.toDataURL();
let blob;

// canvas to blob
canvas.getImageScaledToCanvas().toBlob((blb) => {
  if (blb) {
    blob = new File([blb], file.name, { type: file.type });
  }
});
```

// Reference 1: https://stackoverflow.com/questions/18650168/convert-blob-to-base64
// Reference 2: https://stackoverflow.com/questions/11876175/how-to-get-a-file-or-blob-from-an-object-url
// Reference 3: https://javascript.plainenglish.io/how-to-get-image-data-as-a-base64-url-in-javascript-223a0f2dc514
