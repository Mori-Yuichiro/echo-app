import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance";
import { useEffect, useState } from "react";

export const useToppageHook = () => {
    const [tab, setTab] = useState<"you" | "follow">("you");
    const { instance } = axiosInstance();
    const [tweets, setTweets] = useState<TweetType[] | null>(null);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const { data } = await instance.get<TweetType[]>(
                    "/tweets",
                    { withCredentials: true }
                );
                setTweets(data);
            } catch (err) {
                console.log(err);
            }
        }

        fetchData();
    }, [])

    return {
        tab,
        setTab,
        tweets
    };
}