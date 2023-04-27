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
