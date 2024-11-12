import { ProfileType } from "@/app/types/profile";
import axiosInstance from "@/lib/axiosInstance";
import { getCsrfToken } from "@/lib/csrf_lib";
import { fileRead, fileUpload } from "@/lib/file";
import { profilePatchSchema, ProfilePatchSchemaType } from "@/lib/validation/profile";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleOpenModal, toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";

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

    const saveProfile = async (newProfile: ProfilePatchSchemaType) => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.put<ProfileType>(
                "/users",
                newProfile,
                { withCredentials: true }
            );

            if (status === 200) {
                dispatch(toggleOpenModal(!openModal));
                dispatch(toggleReload(!reload));
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