import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance"
import { getCsrfToken } from "@/lib/csrf_lib";
import { commentPatchSchema, CommentPatchSchemaType } from "@/lib/validation/comment";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useParams, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { useForm } from "react-hook-form";

export const useTweetDetailHook = () => {
    const { instance } = axiosInstance();
    const [tweet, setTweet] = useState<TweetType | null>(null);
    const router = useRouter();
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();
    const { id } = useParams<{ id: string }>();

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors }
    } = useForm<CommentPatchSchemaType>({
        resolver: zodResolver(commentPatchSchema)
    });

    const onClickSendComment = async (data: CommentPatchSchemaType) => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post(
                `/${tweet?.user.id}/comment`,
                {
                    ...data,
                    tweet_id: tweet?.id
                },
                { withCredentials: true }
            );

            if (status === 200) {
                reset({ comment: "" });
                dispatch(toggleReload(!reload));
            }
        } catch (err) {
            reset({ comment: "" });
            console.error(err);
        }
    }

    useEffect(() => {
        const fetchData = async () => {
            const { data } = await instance.get<TweetType>(
                `/tweets/${id}`,
                { withCredentials: true }
            );
            setTweet(data);
        }
        fetchData();
    }, [reload])

    return {
        tweet,
        router,
        register,
        handleSubmit,
        errors,
        onClickSendComment
    };
}