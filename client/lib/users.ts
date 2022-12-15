import axios from "axios";
import getConfig from "next/config";

export const updateUser = async (fullName: string) => {
  try {
    //   const formData = new FormData();
    //   formData.append("full_name", fullName);

    const response = await axios.put(
      `${process.env.NEXT_PUBLIC_BACKEND_HOST}/user`,
      { full_name: fullName }
    );
    return { error: false, errors: response.data };
  } catch (e) {
    console.error(e);
    return {
      error: true,
      status: e.response.status,
      errors: e.response.data,
      message: e.message,
    };
  }
};
