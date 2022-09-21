import { createContext } from "react";


export const UserContext = createContext({user: {}, authenticated: false, error: null});
