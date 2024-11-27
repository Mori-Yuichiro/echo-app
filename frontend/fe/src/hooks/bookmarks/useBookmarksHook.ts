import { BookmarkType } from "@/app/types/bookmark";
import axiosInstance from "@/lib/axiosInstance";
import { useAppSelector } from "@/store/hooks";
import { useRouter } from "next/navigation"
import { useEffect, useState } from "react";

export const useBookmarksHook = () => {
    const router = useRouter();
    const { instance } = axiosInstance();

    const [bookmarks, setBookmarks] = useState<BookmarkType[] | null>(null);

    const reload = useAppSelector(state => state.slice.reload);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data, status } = await instance.get<BookmarkType[]>(
                    "/bookmarks",
                    { withCredentials: true }
                );

                if (status === 200) {
                    setBookmarks(data);
                }
            } catch (err) {
                console.error(err)
            }
        }

        fetchData();
    }, [reload])

    return {
        router,
        bookmarks
    };
}