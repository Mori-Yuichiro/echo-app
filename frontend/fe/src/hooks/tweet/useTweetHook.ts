import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleDeleteTweetModal, toggleReload } from "@/store/slice/slice";
import { usePathname } from "next/navigation"

export const useTweetHook = (tweet: TweetType) => {
    const { instance } = axiosInstance();

    const pathName = usePathname();
    const currentUser = useAppSelector(state => state.slice.currentUser);
    const reload = useAppSelector(state => state.slice.reload);
    const deleteTweetModal = useAppSelector(state => state.slice.deleteTweetModal);
    const dispatch = useAppDispatch();

    const onClickCreateFavorite = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post(
                `/tweets/${tweet.id}/${tweet.user.id}/favorite`,
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
                `/tweets/${tweet.id}/favorite`,
                { withCredentials: true }
            );

            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err);
        }
    }

    const onClickCreateRetweet = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post(
                `/tweets/${tweet.id}/${tweet.user.id}/retweet`,
                {},
                { withCredentials: true }
            );

            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err);
        }
    }

    const onClickDeleteRetweet = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.delete(
                `/tweets/${tweet.id}/retweet`,
                { withCredentials: true }
            );

            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err);
        }
    }

    const onClickCreateBookmark = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post(
                `/tweets/${tweet.id}/bookmark`,
                {},
                { withCredentials: true }
            )
            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err);
        }
    }

    const onClickDeleteBookmark = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.delete(
                `/tweets/${tweet.id}/bookmark`,
                { withCredentials: true }
            )
            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err);
        }
    }

    const onClickDeleteTweetModal = () => {
        dispatch(toggleDeleteTweetModal(!deleteTweetModal));
    }

    return {
        pathName,
        currentUser,
        onClickCreateFavorite,
        onClickDeleteFavorite,
        onClickCreateRetweet,
        onClickDeleteRetweet,
        onClickCreateBookmark,
        onClickDeleteBookmark,
        deleteTweetModal,
        onClickDeleteTweetModal
    };
}