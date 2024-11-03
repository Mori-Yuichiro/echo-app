import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance"
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

export const useTweetDetailHook = () => {
    const { instance } = axiosInstance();
    const [tweet, setTweet] = useState<TweetType | null>(null);
    const router = useRouter();
    const { id } = useParams<{ id: string }>();

    useEffect(() => {
        const fetchData = async () => {
            const { data } = await instance.get<TweetType>(
                `/tweets/${id}`,
                { withCredentials: true }
            );
            setTweet(data);
        }
        fetchData();
    }, [])

    return {
        tweet,
        router
    };
}