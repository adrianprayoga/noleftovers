import { useState } from "react";
import axios from "axios";

const useUpload = () => {
  const [image, setImage] = useState(null);
  const [loading, setLoading] = useState(false);

  const [uploadedImage, setUploadedImage] = useState(null);
  const handleChangeImage = (e) => {
    setImage(e.target.files[0]);
  };

  const handleUploadImage = async () => {
    try {
      setLoading(true);

      const formData = new FormData();
      formData.append("image", image);
      const res = await axios.post(`${process.env.NEXT_PUBLIC_BACKEND_HOST}/images`, formData);
      if (res.data.data) {
        console.log(res.data);
        setUploadedImage(res.data.data);
      }

    } catch (error) {
      console.log(error);
    } finally {
      setImage(null);
      setLoading(false);
    }
  };


  return {
    image,
    uploadedImage,
    loading,
    handleChangeImage,
    handleUploadImage,
    handleRemoveImage,
  };
};

export default useUpload;
