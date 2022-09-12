import "../styles/globals.css";
import { AppProps } from "next/app";
import { createContext, useEffect, useState } from "react";
import { UserContext } from "../hooks/userContext";
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

  return (
    <UserContext.Provider value={appState}>
      <Component {...pageProps}/>;
    </UserContext.Provider>
  );
}

export default App;
