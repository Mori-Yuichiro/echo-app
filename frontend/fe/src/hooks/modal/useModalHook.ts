import { ProfileType } from "@/app/types/profile";
import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { fileRead, fileUpload, uploadImage } from "@/lib/file";
import { profilePatchSchema, ProfilePatchSchemaType } from "@/lib/validation/profile";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleOpenModal, toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRef, useState } from "react";
import { useForm } from "react-hook-form";

type ProfileImagesType = {
    profile_image_url?: string;
    image?: string;
}

export const useModalHook = (profile: ProfileType) => {
    const { instance } = axiosInstance();

    const openModal = useAppSelector(state => state.slice.openModal);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();

    const [profileImageUrl, setProfileImageUrl] = useState<string>(profile.profile_image_url);
    const [avatarUrl, setAvatarUrl] = useState<string>(profile.image);

    const profileInputRef = useRef<HTMLInputElement | null>(null);
    const avatarInputRef = useRef<HTMLInputElement | null>(null);
    const fileOnClickProfile = fileUpload(profileInputRef);
    const fileOnClickAvatar = fileUpload(avatarInputRef);

    const fileInput = async (e: React.ChangeEvent<HTMLInputElement>, setImage: React.Dispatch<React.SetStateAction<string>>) => {
        const files = Array.from(e.target.files || []);

        if (files.length > 0) {
            const imageData = await fileRead(files[0]);
            setImage(imageData);
        }
    }

    const { display_name, bio, location, website } = profile;

    const checkFormData = (data: ProfilePatchSchemaType) => {
        if (
            data.display_name === display_name ||
            data.bio === bio ||
            data.location === location ||
            data.website === website
        ) {
            return true
        }
        return false;
    }

    const updateProfileImage = async () => {
        let newProfileImagaData: string;
        let newAvatarData: string;
        let profileImages: ProfileImagesType = {};

        if (profile.profile_image_url !== profileImageUrl) {
            newProfileImagaData = await (async () => {
                const imageUrl = await uploadImage(instance, profileImageUrl);
                return imageUrl.data.ImageUrl;
            })();
            profileImages.profile_image_url = newProfileImagaData;
        }

        if (profile.image !== avatarUrl) {
            newAvatarData = await (async () => {
                const imageUrl = await uploadImage(instance, avatarUrl);
                return imageUrl.data.ImageUrl;
            })();
            profileImages.image = newAvatarData;
        }

        return profileImages;
    }

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<ProfilePatchSchemaType>({
        resolver: zodResolver(profilePatchSchema),
        defaultValues: {
            display_name,
            bio,
            location,
            website
        }
    });

    const saveProfile = async (data: ProfilePatchSchemaType) => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const profileImages: ProfileImagesType = await updateProfileImage();

            if ("profile_image_url" in profileImages || "image" in profileImages) {
                if (checkFormData(data)) {
                    const { status } = await instance.put<ProfileType>(
                        "/users",
                        {
                            ...data,
                            ...profileImages
                        },
                        { withCredentials: true }
                    );

                    if (status === 200) {
                        dispatch(toggleOpenModal(!openModal));
                        dispatch(toggleReload(!reload));
                    }
                } else {
                    const { status } = await instance.put<ProfileType>(
                        "/users",
                        { ...profileImages },
                        { withCredentials: true }
                    );
                    if (status === 200) {
                        dispatch(toggleOpenModal(!openModal));
                        dispatch(toggleReload(!reload));
                    }
                }
            } else {
                if (checkFormData(data)) {
                    const { status } = await instance.put(
                        "/users",
                        data,
                        { withCredentials: true }
                    );
                    if (status === 200) {
                        dispatch(toggleOpenModal(!openModal));
                        dispatch(toggleReload(!reload));
                    }
                }
            }
        } catch (err) {
            console.error(err);
        }
    }

    return {
        profileImageUrl,
        avatarUrl,
        profileInputRef,
        avatarInputRef,
        fileOnClickProfile,
        fileOnClickAvatar,
        fileInputProfile: (e: React.ChangeEvent<HTMLInputElement>) => fileInput(e, setProfileImageUrl),
        fileInputAvatar: (e: React.ChangeEvent<HTMLInputElement>) => fileInput(e, setAvatarUrl),
        register,
        handleSubmit,
        errors,
        saveProfile
    };
}