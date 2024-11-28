import axiosInstance from "@/lib/axiosInstance"
import { getCsrfToken } from "@/lib/csrf_lib";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleDeleteTweetModal, toggleReload } from "@/store/slice/slice";

export const useDeleteTweetModal = (id: number) => {
    const { instance } = axiosInstance();

    const reload = useAppSelector(state => state.slice.reload);
    const deleteTweetModal = useAppSelector(state => state.slice.deleteTweetModal);
    const dispatch = useAppDispatch();

    const onClickDeleteTweet = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.delete(
                `/tweets/${id}`,
                { withCredentials: true }
            );

            if (status === 200) {
                dispatch(toggleReload(!reload));
                dispatch(toggleDeleteTweetModal(!deleteTweetModal));
            }
        } catch (err) {
            console.error(err);
        }
    }

    const onClickDeleteTweetModal = () => {
        dispatch(toggleDeleteTweetModal(!deleteTweetModal));
    }

    return {
        onClickDeleteTweet,
        onClickDeleteTweetModal
    };
}