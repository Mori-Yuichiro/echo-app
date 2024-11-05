import { ImageType } from "@/app/types/image";
import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { fileRead, fileUpload } from "@/lib/file";
import { tweetPatchSchema, TweetPatchSchemaType } from "@/lib/validation/tweet";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import { v4 as uuid } from "uuid";

export const useToppageHook = () => {
    const [tab, setTab] = useState<"you" | "follow">("you");
    const { instance } = axiosInstance();
    const [tweets, setTweets] = useState<TweetType[] | null>(null);

    const currentUser = useAppSelector(state => state.slice.currentUser);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();

    const {
        register,
        handleSubmit,
        reset,
        formState: { errors }
    } = useForm<TweetPatchSchemaType>({
        resolver: zodResolver(tweetPatchSchema)
    });

    const [images, setImages] = useState<ImageType[]>([]);
    const [imageDatas, setImageDatas] = useState<(string | ArrayBuffer | null)[]>([]);

    const fileInput = async (e: React.ChangeEvent<HTMLInputElement>) => {
        const selectedImages = [];
        const selectedImageDatas = [];

        const files = Array.from(e.target.files || []);
        for (const file of files) {
            const mediaString = uuid();
            selectedImages.push({
                data: await fileRead(file),
                fileName: file.name,
                mediaString
            });
            selectedImageDatas.push(await fileRead(file));
            setImageDatas([...imageDatas, ...selectedImageDatas]);
        }
        setImages([...images, ...selectedImages]);
    }
    const inputRef = useRef<HTMLInputElement | null>(null);
    const fileOnClick = fileUpload(inputRef);

    const onSubmit = async (data: TweetPatchSchemaType) => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post<TweetType>(
                "/tweets",
                data,
                { withCredentials: true }
            );
            if (status === 200) {
                dispatch(toggleReload(!reload));
                reset({ content: "" });
            }
        } catch (err) {
            reset({ content: "" });
            console.error(err);
        }
    }

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
    }, [reload])

    return {
        tab,
        setTab,
        tweets,
        currentUser,
        register,
        handleSubmit,
        errors,
        onSubmit,
        inputRef,
        fileOnClick,
        fileInput,
        images
    };
}