import { EntryType } from "@/app/types/entry";
import { MessageType } from "@/app/types/message";
import axiosInstance from "@/lib/axiosInstance"
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleMessageModal } from "@/store/slice/slice";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

export const useMessages = () => {
    const { instance } = axiosInstance();
    const router = useRouter();
    const [entry, setEntry] = useState<EntryType | null>(null);
    const [messages, setMessages] = useState<MessageType[] | null>(null);

    const { id } = useParams<{ id: string }>();

    const reload = useAppSelector(state => state.slice.reload);
    const currentUser = useAppSelector(state => state.slice.currentUser);
    const messageModal = useAppSelector(state => state.slice.messageModal);
    const dispatch = useAppDispatch();

    const changeMessageModal = () => {
        dispatch(toggleMessageModal(!messageModal));
    }

    useEffect(() => {
        const fetchData = async () => {
            const resEntry = await instance.get(
                `/entry/room/${id}`,
                { withCredentials: true }
            );
            setEntry(resEntry.data);

            const resMessagse = await instance.get(
                `/room/${id}/messages`,
                { withCredentials: true }
            );
            setMessages(resMessagse.data);
        }
        fetchData();
    }, [reload, currentUser])

    return {
        router,
        currentUser,
        entry,
        messages,
        messageModal,
        changeMessageModal
    };
}