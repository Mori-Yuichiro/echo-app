import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleReload } from "@/store/slice/slice";
import { usePathname } from "next/navigation"

export const useTweetHook = (id: number) => {
    const { instance } = axiosInstance();

    const pathName = usePathname();
    const currentUser = useAppSelector(state => state.slice.currentUser);
    const reload = useAppSelector(status => status.slice.reload);
    const dispatch = useAppDispatch();

    const onClickCreateFavorite = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post(
                `/favorite/${id}`,
                {},
                { withCredentials: true }
            );

            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err);
        }
    }

    const onClickDeleteFavorite = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.delete(
                `/favorite/${id}`,
                { withCredentials: true }
            );

            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err);
        }
    }

    return {
        pathName,
        currentUser,
        onClickCreateFavorite,
        onClickDeleteFavorite
    };
}