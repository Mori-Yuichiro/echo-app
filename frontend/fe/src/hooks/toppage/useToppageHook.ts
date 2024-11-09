import { ImageType } from "@/app/types/image";
import { TweetType } from "@/app/types/tweet";
import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { fileRead, fileUpload, uploadImage } from "@/lib/file";
import { tweetPatchSchema, TweetPatchSchemaType } from "@/lib/validation/tweet";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
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

    const deleteDisplayImage = (mediaStr: string) => {
        const selectImage = images.find(image => image.mediaString === mediaStr);
        // 表示する画像を削除
        const updateImages = images.filter(image => image !== selectImage);
        setImages(updateImages);
        // 削除した画像を保存用データから削除
        const updateImageData = imageDatas.filter(imageData => imageData !== selectImage?.data);
        setImageDatas(updateImageData);
    }

    const resetImages = () => {
        setImages([]);
        setImageDatas([]);
    }

    const onSubmit = async (data: TweetPatchSchemaType) => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            if (images.length > 0) {
                const imageUrls: string[] = await Promise.all(imageDatas.map(async (imageData) => {
                    const imageUrl = await uploadImage(instance, imageData);
                    return imageUrl.data.ImageUrl;
                }));

                const imageString = JSON.stringify(imageUrls);

                const { status } = await instance.post<TweetType>(
                    "/tweets",
                    {
                        ...data,
                        image_urls: imageString
                    },
                    { withCredentials: true }
                );
                if (status === 200) {
                    dispatch(toggleReload(!reload));
                    resetImages();
                    reset({ content: "" });
                }
            } else {
                const { status } = await instance.post<TweetType>(
                    "/tweets",
                    data,
                    { withCredentials: true }
                );
                if (status === 200) {
                    dispatch(toggleReload(!reload));
                    reset({ content: "" });
                }
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
        images,
        deleteDisplayImage
    };
}