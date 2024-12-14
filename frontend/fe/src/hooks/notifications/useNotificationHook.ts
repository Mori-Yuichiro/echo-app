import { NotificationType } from "@/app/types/notification";
import axiosInstance from "@/lib/axiosInstance";
import { useAppSelector } from "@/store/hooks";
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react";

export const useNotificationHook = () => {
    const { instance } = axiosInstance();
    const router = useRouter();

    const [notifications, setNotificatons] = useState<NotificationType[] | null>(null);

    const reload = useAppSelector(state => state.slice.reload);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<NotificationType[]>(
                    "/notifications",
                    { withCredentials: true }
                );
                if (status === 200) setNotificatons(data);
            } catch (err) {
                console.error(err)
            }
        }
        fetchData();
    }, [reload])

    return {
        router,
        notifications
    };
}