import axios from "axios";
import { useEffect, useState } from "react"

type CsrfToken = {
    csrf_token: string;
}

export const useHomeHook = () => {
    const [openLoginModal, setOpenLoginModal] = useState<boolean>(false);

    useEffect(() => {
        axios.defaults.withCredentials = true;
        const getCsrfToken = async () => {
            const { data } = await axios.get<CsrfToken>("http://localhost:8080/csrf");
            axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token;
        }
        getCsrfToken();
    }, [])

    return {
        openLoginModal,
        setOpenLoginModal
    };
}