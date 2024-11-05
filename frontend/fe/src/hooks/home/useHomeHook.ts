import { getCsrfToken } from "@/lib/csrf_lib";
import axios from "axios";
import { useEffect, useState } from "react"


export const useHomeHook = () => {
    const [openRegisterModal, setOpenRegisterModal] = useState<boolean>(false);
    const [openLoginModal, setOpenLoginModal] = useState<boolean>(false);

    useEffect(() => {
        const fetchData = async () => {
            axios.defaults.withCredentials = true;
            const csrf = await getCsrfToken();
            axios.defaults.headers.common['X-CSRF-Token'] = csrf;
        }
        fetchData();
    }, [])

    return {
        openRegisterModal,
        setOpenRegisterModal,
        openLoginModal,
        setOpenLoginModal
    };
}