import { createContext } from "react";
import axios from "axios";

axios.defaults.baseURL = `${process.env.NEXT_PUBLIC_BACKEND_HOST}`;
axios.defaults.withCredentials = true;

interface User {
    id: number;
    auth_method: string;
    full_name: string
    email: string;
    last_login: string;
    picture: string
}

interface UserAuth {
    user: User;
    authenticated: boolean;
    error: object;
}

var UserAuthdefaultContext: UserAuth = {user: null, authenticated: false, error: null}

export const UserContext = createContext(UserAuthdefaultContext);
