"use client"

import { useBookmarksHook } from "@/hooks/bookmarks/useBookmarksHook";
import Loading from "@/components/Loading";
import Tweet from "@/components/tweet/Tweet";

export default function Bookmarks() {
    const {
        router,
        bookmarks
    } = useBookmarksHook();

    return (
        <div>
            <div className="flex items-center gap-x-3 px-3 py-2 border-b border-gray-200">
                <div
                    className="cursor-pointer"
                    onClick={() => router.back()}
                >
                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 256 256"><path fill="currentColor" d="M224 128a8 8 0 0 1-8 8H59.31l58.35 58.34a8 8 0 0 1-11.32 11.32l-72-72a8 8 0 0 1 0-11.32l72-72a8 8 0 0 1 11.32 11.32L59.31 120H216a8 8 0 0 1 8 8Z" /></svg>
                </div>
                <h1 className="font-bold text-lg">Bookmarks</h1>
            </div>
            {bookmarks ? (
                <>
                    {bookmarks.map(bookmark => (
                        <div key={bookmark.id}>
                            <Tweet tweet={bookmark.tweet} />
                        </div>
                    ))}
                </>
            ) : (
                <Loading />
            )}
        </div>
    );
}