import { ProfileType } from "@/app/types/profile";
import axiosInstance from "@/lib/axiosInstance";
import { fileRead, fileUpload } from "@/lib/file";
import { profilePatchSchema, ProfilePatchSchemaType } from "@/lib/validation/profile";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRef, useState } from "react";
import { useForm } from "react-hook-form";

export const useModalHook = (profile: ProfileType) => {
    const { instance } = axiosInstance();

    const [profileImageUrl, setProfileImageUrl] = useState<string>(profile.profileImageUrl);
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

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<ProfilePatchSchemaType>({
        resolver: zodResolver(profilePatchSchema)
    });

    const saveProfile = async (data: ProfilePatchSchemaType) => {
        try {
            console.log(data);
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