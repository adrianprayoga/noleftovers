import "../styles/globals.css";
import { AppProps } from "next/app";
import { useEffect, useState } from "react";
import axios from "axios";

function App({ Component, pageProps }: AppProps) {
  const [appState, setAppState] = useState({
    user: {},
    error: null,
    authenticated: false,
  });

  useEffect(() => {
    const getUser = async () => {
      try {
        const res = await axios.get(`/auth/success`);
        if (res.status == 200) {
          setAppState({
            user: res.data.user,
            authenticated: true,
            error: null,
          });
        }
        console.log(res);
      } catch (err) {
        console.error("call error", err);
        setAppState({
          user: {},
          authenticated: false,
          error: "user is not authenticated",
        });
      }
    };

    !appState.authenticated && getUser();
  }, [appState.authenticated]);

  return <Component {...pageProps} appState={appState}/>;
}

export default App;
