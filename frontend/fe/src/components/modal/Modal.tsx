"use client"

import { Dispatch, SetStateAction } from "react";
import { ProfileType } from "@/app/types/profile";
import { useModalHook } from "@/hooks/modal/useModalHook";
import Button from "../Button";


export default function Modal({
    openModal,
    setOpenModal,
    profile
}: {
    openModal: boolean,
    setOpenModal: Dispatch<SetStateAction<boolean>>,
    profile: ProfileType
}) {
    const {
        profileImageUrl,
        avatarUrl,
        profileInputRef,
        fileOnClickProfile,
        avatarInputRef,
        fileOnClickAvatar,
        fileInputProfile,
        fileInputAvatar,
        register,
        handleSubmit,
        errors,
        saveProfile
    } = useModalHook(profile);

    return (
        <div className="relative z-10" aria-labelledby="modal-title" role="dialog" aria-modal="true">
            <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
            <div className="fixed inset-0 z-10 w-screen overflow-y-auto max-sm:mx-auto">
                <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <div className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg max-sm:w-full">
                        <div className="bg-white space-y-2 px-4 pb-4 pt-5 sm:p-4 sm:pb-4">
                            <div className="flex justify-between">
                                <div
                                    className="cursor-pointer flex gap-x-5 items-center"
                                    onClick={() => setOpenModal(!openModal)}
                                >
                                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 16 16"><path fill="currentColor" d="M7.293 8L3.146 3.854a.5.5 0 1 1 .708-.708L8 7.293l4.146-4.147a.5.5 0 0 1 .708.708L8.707 8l4.147 4.146a.5.5 0 0 1-.708.708L8 8.707l-4.146 4.147a.5.5 0 0 1-.708-.708z" /></svg>
                                    <h1 className="text-lg">Edit Profile</h1>
                                </div>
                                <Button
                                    className="border rounded-full px-3"
                                    onClick={handleSubmit(saveProfile)}
                                >Save</Button>
                            </div>
                            {errors && <p className="text-red-500">{errors.displayName?.message}</p>}
                            <div className="h-48 bg-slate-400 w-full relative">
                                <div className="h-full w-full">
                                    {profileImageUrl && <img
                                        className="w-full h-full object-cover"
                                        src={profileImageUrl}
                                        alt="プロフィール画像"
                                    />}
                                </div>
                                <div>
                                    <input
                                        type="file"
                                        className="hidden"
                                        ref={profileInputRef}
                                        onChange={fileInputProfile}
                                    />
                                    <Button
                                        className="absolute top-1/2 right-1/2"
                                        onClick={fileOnClickProfile}
                                    >
                                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 42 42"><path fill="currentColor" d="M14.41 22.58c0 3.359 2.73 6.09 6.09 6.09c3.359 0 6.09-2.73 6.09-6.09s-2.73-6.09-6.09-6.09a6.095 6.095 0 0 0-6.09 6.09zM3.5 36.5h34c2.63 0 3-.37 3-3v-23c0-2.462-.38-3-3-3h-10c0-2.57-.42-3-3-3h-8c-2.55 0-3 .48-3 3h-10c-2.58 0-3 .692-3 3v23c0 2.6.38 3 3 3zm7.64-13.92c0-5.17 4.19-9.359 9.36-9.359s9.359 4.189 9.359 9.359s-4.189 9.359-9.359 9.359s-9.36-4.189-9.36-9.359z" /></svg>
                                    </Button>
                                </div>
                            </div>
                            <div className="bg-slate-400 rounded-full max-h-48 max-w-48 h-24 w-24 relative">
                                {avatarUrl && <img
                                    src={avatarUrl}
                                    alt="アイコン"
                                    className="rounded-full w-full h-full object-cover"
                                />}
                                <div>
                                    <input
                                        type="file"
                                        multiple
                                        className="hidden"
                                        ref={avatarInputRef}
                                        onChange={fileInputAvatar}
                                    />
                                    <Button
                                        className="absolute top-1/2 right-1/2 cursor-pointer translate-x-1/2 translate-y-[-50%]"
                                        onClick={fileOnClickAvatar}
                                    >
                                        <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 42 42"><path fill="currentColor" d="M14.41 22.58c0 3.359 2.73 6.09 6.09 6.09c3.359 0 6.09-2.73 6.09-6.09s-2.73-6.09-6.09-6.09a6.095 6.095 0 0 0-6.09 6.09zM3.5 36.5h34c2.63 0 3-.37 3-3v-23c0-2.462-.38-3-3-3h-10c0-2.57-.42-3-3-3h-8c-2.55 0-3 .48-3 3h-10c-2.58 0-3 .692-3 3v23c0 2.6.38 3 3 3zm7.64-13.92c0-5.17 4.19-9.359 9.36-9.359s9.359 4.189 9.359 9.359s-4.189 9.359-9.359 9.359s-9.36-4.189-9.36-9.359z" /></svg>
                                    </Button>
                                </div>
                            </div>
                            <div className="flex flex-col gap-y-3">
                                <input
                                    id="displayName"
                                    className="border border-slate-400 rounded-sm p-3"
                                    type="text"
                                    placeholder="Name"
                                    {...register("displayName")}
                                />
                                <textarea
                                    id="bio"
                                    className="border border-slate-400 rounded-sm p-3"
                                    placeholder="Bio"
                                    {...register("bio")}
                                />
                                <input
                                    id="location"
                                    className="border border-slate-400 rounded-sm p-3"
                                    type="text"
                                    placeholder="location"
                                    {...register("location")}
                                />
                                <input
                                    id="website"
                                    className="border border-slate-400 rounded-sm p-3"
                                    type="text"
                                    placeholder="website"
                                    {...register("website")}
                                />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}