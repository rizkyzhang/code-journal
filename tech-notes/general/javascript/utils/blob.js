const blobToBase64 = (blob) => {
  const reader = new FileReader();
  reader.readAsDataURL(blob);
  return new Promise((resolve) => {
    reader.onloadend = () => {
      resolve(reader.result);
    };
  });
};

const base64ToBlob = async (base64, fileName, fileType) => {
  return await fetch(base64)
    .then((res) => res.blob())
    .then((blob) => new File([blob], fileName, { type: fileType }));
};

// Reference 1: https://stackoverflow.com/questions/18650168/convert-blob-to-base64
// Reference 2: https://stackoverflow.com/questions/11876175/how-to-get-a-file-or-blob-from-an-object-url
