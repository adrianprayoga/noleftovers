export const dataURLtoFile = (dataurl, filename) => {
  var arr = dataurl.split(","),
    mime = arr[0].match(/:(.*?);/)[1],
    bstr = atob(arr[1]),
    n = bstr.length,
    u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new File([u8arr], filename, { type: mime });
};

export const cropImage = (crop, imageUrl) => {
  var canvas = document.createElement("canvas");
  var ctx = canvas.getContext("2d");

  let img = new Image();
  img.src = imageUrl;

  const w = (crop.width * img.naturalWidth) / 100;
  const h = (crop.height * img.naturalHeight) / 100;

  // devicePixelRatio slightly increases sharpness on retina devices
  // at the expense of slightly slower render times and needing to
  // size the image back down if you want to download/upload and be
  // true to the images natural size.
  const pixelRatio = window.devicePixelRatio;

  canvas.width = Math.floor(w * pixelRatio);
  canvas.height = Math.floor(h * pixelRatio);

  ctx.scale(pixelRatio, pixelRatio);
  ctx.imageSmoothingQuality = "high";

  ctx.drawImage(
    img,
    (crop.x * img.naturalWidth * pixelRatio) / 100,
    (crop.y * img.naturalHeight * pixelRatio) / 100,
    img.naturalWidth,
    img.naturalHeight,
    0,
    0,
    img.naturalWidth,
    img.naturalHeight
  );

  return { canvas, w, h };
};

export class Rectangle {
  height: number;
  width: number;

  constructor(width, height) {
    this.width = width;
    this.height = height;
  }

  get isRectangle() {
    return Math.abs(this.height - this.width) < 10e-2
  }

}
