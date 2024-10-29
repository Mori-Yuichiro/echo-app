import { useState } from "react"

export const useHomeHook = () => {
    const [openRegisterModal, setOpenRegisterModal] = useState<boolean>(false);

    return {
        openRegisterModal,
        setOpenRegisterModal
    };
}