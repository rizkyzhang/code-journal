const getImageSize = (file) => {
  return new Promise((resolve) => {
    const reader = new FileReader();

    reader.readAsDataURL(file);
    reader.addEventListener(
      "load",
      (e) => {
        if (e.target) {
          const image = new Image();
          image.src = e.target.result;
          image.onload = () => {
            const { height, width } = image;
            resolve({ height, width });
          };
        }
      },
      false
    );
  });
};
