import { useState } from "react"

export const useHomeHook = () => {
    const [openLoginModal, setOpenLoginModal] = useState<boolean>(false);

    return {
        openLoginModal,
        setOpenLoginModal
    };
}