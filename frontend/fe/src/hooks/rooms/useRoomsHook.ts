import { RoomType } from "@/app/types/room";
import axiosInstance from "@/lib/axiosInstance";
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react";

export const useRoomsHook = () => {
    const { instance } = axiosInstance();
    const router = useRouter();

    const [rooms, setRooms] = useState<RoomType[] | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data } = await instance.get<RoomType[]>(
                    "/room",
                    { withCredentials: true }
                );
                setRooms(data);
            } catch (err) {
                console.error(err)
            }
        }
        fetchData();
    }, [])

    return {
        router,
        rooms
    }
}