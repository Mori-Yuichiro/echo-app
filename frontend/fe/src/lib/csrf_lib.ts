import { CsrfToken } from "@/app/types/csrf";
import axios from "axios";

export const getCsrfToken = async () => {
    const { data } = await axios.get<CsrfToken>(
        "http://localhost:8080/csrf",
        { withCredentials: true }
    );
    return data.csrf_token;
}