import { ProfileType } from "@/app/types/profile"
import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleOpenModal, toggleReload } from "@/store/slice/slice";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react"

export const useProfileHook = () => {
    const [profile, setProfile] = useState<ProfileType | null>(null);
    const [tab, setTab] = useState<
        "posts" |
        "comments" |
        "retweets" |
        "articles" |
        "medias" |
        "likes"
    >("posts");
    const { id } = useParams<{ id: string }>();
    const router = useRouter();

    const currentUser = useAppSelector(state => state.slice.currentUser);
    const openModal = useAppSelector(state => state.slice.openModal);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();

    const onClickToggleModal = () => {
        dispatch(toggleOpenModal(!openModal));
    }

    const { instance } = axiosInstance();

    const onClickCreateRelationship = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post(
                `/users/${profile?.id}/follow`,
                {},
                { withCredentials: true }
            );

            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err)
        }
    }

    const onClickDeleteRelationship = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.delete(
                `/users/${profile?.id}/follow`,
                { withCredentials: true }
            );

            if (status === 200) dispatch(toggleReload(!reload));
        } catch (err) {
            console.error(err)
        }
    }

    useEffect(() => {
        const fetchData = async () => {
            const { data } = await instance.get<ProfileType>(
                `/users/${id}`,
                { withCredentials: true }
            );
            setProfile(data);
        }
        fetchData();
    }, [reload])

    return {
        profile,
        router,
        currentUser,
        tab,
        setTab,
        openModal,
        onClickToggleModal,
        onClickCreateRelationship,
        onClickDeleteRelationship
    };
}