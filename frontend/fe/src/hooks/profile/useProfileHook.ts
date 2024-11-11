import { ProfileType } from "@/app/types/profile"
import axiosInstance from "@/lib/axiosInstance";
import { useAppSelector } from "@/store/hooks";
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
    const [openModal, setOpenModal] = useState<boolean>(false);
    const { id } = useParams<{ id: string }>();
    const router = useRouter();

    const currentUser = useAppSelector(state => state.slice.currentUser);

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
    }, [])

    return {
        profile,
        router,
        currentUser,
        tab,
        setTab,
        openModal,
        setOpenModal
    };
}