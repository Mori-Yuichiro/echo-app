import { TweetType } from "@/app/types/tweet";
import { useTweetHook } from "@/hooks/tweet/useTweetHook";
import Link from "next/link";

export default function Tweet({ tweet }: { tweet: TweetType }) {
    const {
        pathName,
        currentUser,
        onClickCreateFavorite,
        onClickDeleteFavorite
    } = useTweetHook(tweet.id);

    return (
        <div className="border-b border-gray-200 px-4 py-3 space-y-3">
            <div className="flex gap-x-3">
                <div className="bg-slate-400 rounded-full w-8 h-8">
                    {tweet.user.image &&
                        <img
                            src={tweet.user.image}
                            alt="アイコン"
                            className="w-full h-full rounded-full"
                        />
                    }
                </div>
                <div className="w-full">
                    <Link href={`/profile/${tweet.user.id}`}>
                        {tweet.user.display_name ? (
                            <p>{tweet.user.display_name}</p>
                        ) : (
                            <p>{tweet.user.name}</p>
                        )}
                    </Link>
                    {(pathName.match("/tweets\/([0-9]+)")) ? (
                        <p>{tweet.content}</p>
                    ) : (
                        <Link
                            href={`/tweets/${tweet.id}`}
                            className="w-full"
                        >
                            <p>{tweet.content}</p>
                        </Link>
                    )}
                    {(tweet.image_urls && tweet.image_urls.length > 0) &&
                        <div className={`gap-x-4 gap-y-3 grid ${tweet.image_urls.length === 1 ? "grid-cols-1" : "grid-cols-2"}`}>
                            {tweet.image_urls.map(imageUrl => {
                                return (
                                    <div key={imageUrl} className={(tweet.image_urls.length > 0) ? "w-1/2" : undefined}>
                                        <Link
                                            href={imageUrl}
                                            target="_blank"
                                            rel="noreferrer"
                                        >
                                            <img src={imageUrl} alt="ツイート画像" />
                                        </Link>
                                    </div>
                                );
                            })}
                        </div>
                    }
                </div>
            </div>
            <div className="flex justify-between ml-8">
                <div>
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 1024 1024"><path fill="currentColor" d="M512 896q-66 0-134-16q-34 40-69.5 69.5t-60 43.5t-47.5 21.5t-30.5 8.5t-10.5 1q26-57 30-124.5T176 786Q94 723 47 635T0 448q0-91 40.5-174t109-143T313 35.5T512 0t199 35.5T874.5 131t109 143t40.5 174t-40.5 174t-109 143T711 860.5T512 896z" /></svg>
                </div>
                <div>
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 42 42"><path fill="currentColor" d="m24.5 30.5l7.96 7.371L40.5 30.5h-5V9.549c0-2.5-.561-3.049-3-3.049h-18l6.641 6H29.5v18h-5zm-8-18L8.52 5.16L.5 12.5h5v21.049c0 2.5.62 2.951 3 2.951h18.32l-6.32-6h-9v-18h5z" /></svg>
                </div>
                {(tweet.favorites && tweet.favorites.filter(fav => fav.userId === currentUser?.id)) ? (
                    <div className="cursor-pointer flex items-center gap-x-2">
                        <div
                            className="text-red-400"
                            onClick={onClickDeleteFavorite}
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 20 20"><path fill="currentColor" d="m10 3.22l-.61-.6a5.5 5.5 0 0 0-7.78 7.77L10 18.78l8.39-8.4a5.5 5.5 0 0 0-7.78-7.77l-.61.61z" /></svg>
                        </div>
                        <span>{tweet.favorites.length}</span>
                    </div>
                ) : (
                    <div className="cursor-pointer flex items-center gap-x-2">
                        <div
                            onClick={onClickCreateFavorite}
                        >
                            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"><path fill="currentColor" d="M12 4.528a6 6 0 0 0-8.243 8.715l6.829 6.828a2 2 0 0 0 2.828 0l6.829-6.828A6 6 0 0 0 12 4.528zm-1.172 1.644l.465.464a1 1 0 0 0 1.414 0l.465-.464a4 4 0 1 1 5.656 5.656L12 18.657l-6.828-6.829a4 4 0 0 1 5.656-5.656z" /></svg>
                        </div>
                        <span>0</span>
                    </div>
                )}
                <div>
                    <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24"><path fill="currentColor" d="M4 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v17a1 1 0 0 1-1.581.814L12 17.229l-6.419 4.585A1 1 0 0 1 4 21V4zm14 0H6v15.057l5.419-3.87a1 1 0 0 1 1.162 0L18 19.056V4z" /></svg>
                </div>
            </div>
        </div>
    );
}