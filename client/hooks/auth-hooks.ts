import { useState, useEffect } from "react";
import axios from "axios";

axios.defaults.baseURL = `${process.env.NEXT_PUBLIC_BACKEND_HOST}`;
axios.defaults.withCredentials = true;

const useAuth = () => {
  const [state, setState] = useState({
    user: {},
    error: null,
    authenticated: false,
  });

  useEffect(() => {
    const getUser = async () => {
      try {
        const res = await axios.get(`/auth/success`);
        if (res.status == 200) {
          setState({
            user: res.data.user,
            authenticated: true,
            error: null,
          });
        }
      } catch (err) {
        console.error("call error", err);
        setState({
          user: {},
          authenticated: false,
          error: "user is not authenticated",
        });
      }
    };

    getUser();
  }, []);

  return {
    state,
  };
};

export default useAuth;
