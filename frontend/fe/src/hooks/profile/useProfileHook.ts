import { EntryType } from "@/app/types/entry";
import { ProfileType } from "@/app/types/profile"
import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleOpenModal, toggleReload } from "@/store/slice/slice";
import { useParams, useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react"

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
    const [entries, setEntries] = useState<EntryType[] | null>(null);
    const [anotherEntries, setAnotherEntries] = useState<EntryType[] | null>(null);
    const [commonRoomId, setCommonRoomId] = useState<number>(0);
    const [isRoom, setIsRoom] = useState<boolean>(false);
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

    const checkCommonRoomId = useCallback(() => {
        if (!entries || !anotherEntries) return;
        const currentUserRoomIds = entries.map(entry => entry.room_id);
        const anotherUserRoomIds = anotherEntries.map(entry => entry.room_id);
        const commonRoomIds = currentUserRoomIds.filter(roomId => anotherUserRoomIds.includes(roomId));

        if (commonRoomIds.length > 0) {
            setCommonRoomId(commonRoomIds[0]);
            setIsRoom(true);
        }
    }, [entries, anotherEntries])

    const onClickCreateRoom = async () => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const resRoom = await instance.post(
                "/room",
                {},
                { withCredentials: true }
            );

            if (resRoom.status === 200) {
                const resCurrentUserEntry = await instance.post(
                    "/entry",
                    {
                        user_id: currentUser?.id,
                        room_id: resRoom.data.id
                    },
                    { withCredentials: true }
                );
                const resAnotherUserEntry = await instance.post(
                    "/entry",
                    {
                        user_id: profile?.id,
                        room_id: resRoom.data.id
                    },
                    { withCredentials: true }
                );

                if (resCurrentUserEntry.status === 200 && resAnotherUserEntry.status === 200) router.push(`/rooms/${resRoom.data.id}`);
            }
        } catch (err) {
            console.error(err);
        }
    }

    useEffect(() => {
        const fetchData = async () => {
            if (!currentUser) return;

            const { data } = await instance.get<ProfileType>(
                `/users/${id}`,
                { withCredentials: true }
            );
            setProfile(data);

            const resEntries = await instance.get<EntryType[]>(
                `/entry/${currentUser?.id}`,
                { withCredentials: true }
            );
            if (resEntries.status === 200) setEntries(resEntries.data);

            const resAnotherEntries = await instance.get<EntryType[]>(
                `/entry/${id}`,
                { withCredentials: true }
            );
            if (resAnotherEntries.status === 200) setAnotherEntries(resAnotherEntries.data);
        }
        fetchData();
    }, [reload, currentUser, id])

    useEffect(() => {
        checkCommonRoomId();
    }, [entries, anotherEntries])

    return {
        profile,
        router,
        currentUser,
        tab,
        setTab,
        openModal,
        onClickToggleModal,
        onClickCreateRelationship,
        onClickDeleteRelationship,
        commonRoomId,
        isRoom,
        onClickCreateRoom
    };
}