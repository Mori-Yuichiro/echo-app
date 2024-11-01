import { CsrfToken } from "@/app/types/csrf";
import axios from "axios";

export const useError = () => {
    const getCsrfToken = async () => {
        const { data } = await axios.get<CsrfToken>("http://localhost:8080/csrf");
        axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token;
    }

    const switchErrorHandling = (msg: string) => {
        switch (msg) {
            case "missing csrf token in request header":
                getCsrfToken();
                return "CSRFトークンがヘッダーに設定されていません";
            case "invalid csrf token":
                getCsrfToken();
                return "CSRFトークンが間違っています";
            case `ERROR: duplicate key value violates unique constraint "uni_users_email" (SQLSTATE 23505)`:
                return "そちらのEmailはすでに登録されています";
            case "crypto/bcrypt: hashedPassword is not the hash of the given password":
                return "Passwordが間違っています";
            default:
                return null;
        }
    }

    return { switchErrorHandling }
}