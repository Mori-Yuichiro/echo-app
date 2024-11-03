"use client"

import Loading from "@/components/Loading";
import Tweet from "@/components/tweet/Tweet";
import { useToppageHook } from "@/hooks/toppage/useToppageHook";

export default function Toppage() {
    const {
        tab,
        setTab,
        tweets
    } = useToppageHook();

    return (
        <div>
            <ul className="flex flex-wrap text-sm font-medium text-center text-gray-500 border-b border-gray-200 dark:border-gray-700 dark:text-gray-400">
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