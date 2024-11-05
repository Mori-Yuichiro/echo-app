"use client"

import Button from "@/components/Button";
import Loading from "@/components/Loading";
import Tweet from "@/components/tweet/Tweet";
import { useToppageHook } from "@/hooks/toppage/useToppageHook";

export default function Toppage() {
    const {
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
    } = useToppageHook();

    return (
        <div>
            <ul className="flex flex-wrap text-sm font-medium text-center text-gray-500 dark:border-gray-700 dark:text-gray-400 border-b border-gray-200">
                <li
                    className="w-1/2 border-gray-200 border-r cursor-pointer hover:bg-gray-50"
                    onClick={() => setTab("you")}
                >
                    <span className={`inline-block p-4 rounded-t-lg dark:bg-gray-800 dark:text-blue-500 ${tab === "you" && "border-b-4 border-blue-400"}`}>For you</span>
                </li>
                <li
                    className="w-1/2 cursor-pointer hover:bg-gray-50"
                    onClick={() => setTab("follow")}
                >
                    <span className={`inline-block p-4 rounded-t-lg dark:bg-gray-800 dark:text-blue-500 ${tab === "follow" && "border-b-4 border-blue-400"}`}>Following</span>
                </li>
            </ul>
            <div className="border-b border-gray-200">
                {errors.content && <p className="text-red-500 p-2">{errors.content.message}</p>}
                <div className="flex gap-x-2 items-start px-3 py-2">
                    <div className="bg-slate-400 rounded-full w-8 h-8">
                        {currentUser?.image &&
                            <img
                                src={currentUser.image}
                                alt="アイコン"
                                className="w-full h-full rounded-full"
                            />
                        }
                    </div>
                    <div className="w-full flex flex-col gap-y-2">
                        <input
                            id="content"
                            type="text"
                            placeholder="What is happening"
                            className="w-full p-2"
                            {...register("content")}
                        />
                        <div className="flex justify-between">
                            <input
                                type="file"
                                multiple
                                className="hidden"
                                ref={inputRef}
                                onChange={fileInput}
                            />
                            <Button onClick={fileOnClick}>
                                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="20" viewBox="0 0 56 56"><path fill="currentColor" d="M7.715 49.574h40.57c4.899 0 7.36-2.437 7.36-7.265V13.69c0-4.828-2.461-7.265-7.36-7.265H7.715C2.84 6.426.355 8.84.355 13.69v28.62c0 4.851 2.485 7.265 7.36 7.265m31.57-21.633c-1.055-.937-2.25-1.43-3.515-1.43c-1.313 0-2.462.446-3.54 1.407l-10.593 9.469l-4.336-3.938c-.985-.867-2.04-1.336-3.164-1.336c-1.032 0-2.04.446-3 1.313L4.129 39.73V13.88c0-2.438 1.312-3.68 3.656-3.68h40.43c2.32 0 3.656 1.242 3.656 3.68v25.875Zm-21.469.258c3.024 0 5.508-2.484 5.508-5.531c0-3.023-2.484-5.531-5.508-5.531c-3.046 0-5.53 2.508-5.53 5.531a5.541 5.541 0 0 0 5.53 5.531" /></svg>
                            </Button>
                            <Button
                                className="bg-cyan-400 text-white w-1/10 rounded-full font-bold px-2 py-1"
                                onClick={handleSubmit(onSubmit)}
                            >Post</Button>
                        </div>
                    </div>
                </div>
            </div>
            {tweets ? (
                <>
                    {tweets.map(tweet => (
                        <div key={tweet.id}>
                            <Tweet tweet={tweet} />
                        </div>
                    ))}
                </>
            ) : (<Loading />)}
        </div>
    );
}