const downloadContent = async (src, name) => {
  const content = await fetch(src);
  const contentBlob = await content.blob();
  const contentURL = URL.createObjectURL(contentBlob);
  const link = document.createElement("a");

  link.href = contentURL;
  link.target = "_blank";
  link.download = name;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

// Reference: https://dev.to/sbodi10/download-images-using-javascript-51a9
