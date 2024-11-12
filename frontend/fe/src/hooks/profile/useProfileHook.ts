import { ProfileType } from "@/app/types/profile"
import axiosInstance from "@/lib/axiosInstance";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleOpenModal } from "@/store/slice/slice";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react"

export const useProfileHook = () => {
    const [profile, setProfile] = useState<ProfileType | null>(null);
    const [tab, setTab] = useState<
        "posts" |
        "comments" |
        "highlights" |
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
        onClickToggleModal
    };
}