import { CsrfToken } from "@/app/types/csrf";
import axios from "axios";
import { useEffect, useState } from "react"


export const useHomeHook = () => {
    const [openRegisterModal, setOpenRegisterModal] = useState<boolean>(false);
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
        openRegisterModal,
        setOpenRegisterModal,
        openLoginModal,
        setOpenLoginModal
    };
}