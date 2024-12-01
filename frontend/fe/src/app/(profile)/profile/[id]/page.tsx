"use client"

import Button from "@/components/Button";
import Comment from "@/components/comment/Comment";
import Loading from "@/components/Loading";
import Modal from "@/components/modal/Modal";
import Tweet from "@/components/tweet/Tweet";
import { useProfileHook } from "@/hooks/profile/useProfileHook";
import Link from "next/link";

export default function ProfilePage() {
    const {
        profile,
        router,
        currentUser,
        tab,
        setTab,
        openModal,
        onClickToggleModal,
        onClickCreateRelationship,
        onClickDeleteRelationship
    } = useProfileHook();

    return (
        <>
            {profile ? (
                <>
                    <div>
                        <div className="flex gap-x-4 items-center p-2 border-b border-black">
                            <div
                                className="cursor-pointer"
                                onClick={() => router.back()}
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" width="15" height="15" viewBox="0 0 512 512"><path fill="currentColor" d="M213.3 205.3v-128L0 248l213.3 170.7v-128H512v-85.4z" /></svg>
                            </div>
                            <h1 className="font-bold text-lg">{profile.display_name ? profile.display_name : profile.name}</h1>
                        </div>
                        <div>
                            <div className="h-48 bg-slate-400 relative">
                                {profile.profile_image_url && <img className="w-full h-full" src={profile.profile_image_url} alt="プロフィール画像" />}
                            </div>
                            <div className="ml-3 bg-slate-400 w-28 h-28 md:w-32 md:h-32 rounded-full absolute top-40">
                                {profile.image && <img className="w-full h-full rounded-full" src={profile.image} alt="プロフィール・アイコン" />}
                            </div>
                            <div className="flex justify-end p-4 items-center gap-x-3">
                                {(profile.id === currentUser?.id) ? (
                                    <Button
                                        className="rounded-full border border-black px-2 py-1"
                                        onClick={onClickToggleModal}
                                    >Edit Profile</Button>
                                ) : (
                                    <>
                                        {(profile.followers && profile.followers.filter(follower => follower.follower_id === currentUser?.id)) ? (
                                            <Button
                                                className="rounded-full bg-slate-400 text-white px-2 py-1"
                                                onClick={onClickDeleteRelationship}
                                            >Following</Button>
                                        ) : (
                                            <Button
                                                className="rounded-full border border-black px-2 py-1"
                                                onClick={onClickCreateRelationship}
                                            >Follow</Button>
                                        )}
                                    </>
                                )}
                            </div>
                        </div>
                        <div className="mb-8 px-4 space-y-10">
                            <h1 className="text-xl">{profile.display_name ? profile.display_name : profile.name}</h1>
                            <p>{profile.bio}</p>
                            <p>{profile.website}</p>
                            <div className="flex gap-x-3">
                                <p>{profile.followeds ? profile.followeds.length : 0} Followings</p>
                                <Link href={`/profile/${profile.id}/followers`}>
                                    <p>{profile.followers ? profile.followers.length : 0} Followers</p>
                                </Link>
                            </div>
                        </div>
                        <ul className="list-reset flex border-b border-black overflow-x-auto">
                            <li
                                className="-mb-px mr-1 w-1/2 mx-auto border-black text-center cursor-pointer hover:bg-slate-300"
                                onClick={() => setTab("posts")}
                            >
                                <span className={`inline-block rounded-t py-1 px-4 text-blue-dark font-semibold ${tab === "posts" && "border-b-4 border-blue-300"}`}
                                >Posts</span>
                            </li>
                            <li
                                className="-mb-px mr-1 w-1/2 text-center cursor-pointer hover:bg-slate-300"
                                onClick={() => setTab("comments")}
                            >
                                <span className={`inline-block py-1 px-4 text-blue hover:text-blue-darker font-semibold ${tab === "comments" && "border-b-4 border-blue-300"}`}
                                >Comments</span>
                            </li>
                            <li
                                className="-mb-px mr-1 w-1/2 text-center cursor-pointer hover:bg-slate-300"
                                onClick={() => setTab("retweets")}
                            >
                                <span className={`inline-block py-1 px-4 text-blue hover:text-blue-darker font-semibold ${tab === "retweets" && "border-b-4 border-blue-300"}`}>Retweets</span>
                            </li>
                            <li
                                className="-mb-px mr-1 w-1/2 text-center cursor-pointer hover:bg-slate-300"
                                onClick={() => setTab("articles")}
                            >
                                <span className={`inline-block py-1 px-4 text-blue hover:text-blue-darker font-semibold ${tab === "articles" && "border-b-4 border-blue-300"}`}>Articles</span>
                            </li>
                            <li
                                className="-mb-px mr-1 w-1/2 text-center cursor-pointer hover:bg-slate-300"
                                onClick={() => setTab("medias")}
                            >
                                <span className={`inline-block py-1 px-4 text-blue hover:text-blue-darker font-semibold ${tab === "medias" && "border-b-4 border-blue-300"}`}>Medias</span>
                            </li>
                            <li
                                className="-mb-px mr-1 w-1/2 text-center cursor-pointer hover:bg-slate-300"
                                onClick={() => setTab("likes")}
                            >
                                <span className={`inline-block py-1 px-4 text-blue hover:text-blue-darker font-semibold ${tab === "likes" && "border-b-4 border-blue-300"}`}>Likes</span>
                            </li>
                        </ul>
                        {tab === "posts" ? (
                            <>
                                {profile.tweets.map(tweet => (
                                    <div
                                        key={`profile-tweet-${tweet.id}`}
                                        className="border-black border-b">
                                        <Tweet tweet={tweet} />
                                    </div>
                                ))}
                            </>
                        ) : (tab === "comments") ? (
                            <>
                                {profile.comments.map(comment => (
                                    <div key={`profile-comment-${comment.id}`}>
                                        <Comment comment={comment} />
                                    </div>
                                ))}
                            </>
                        ) : (tab === "retweets") ? (
                            <>
                                {profile.retweets.map(retweet => (
                                    <div key={`profile-retweet-${retweet.id}`}>
                                        <Tweet tweet={retweet.tweet} />
                                    </div>
                                ))}
                            </>
                        ) : (tab === "likes") ? (
                            <>
                                {profile.favorites.map(favorite => (
                                    <div key={`profile-favorite-${favorite.id}`}>
                                        <Tweet tweet={favorite.tweet} />
                                    </div>
                                ))}
                            </>
                        ) : <></>
                        }
                    </div>
                    {openModal &&
                        <Modal
                            openModal={openModal}
                            setOpenModal={onClickToggleModal}
                            profile={profile}
                        />
                    }
                </>
            ) : (
                <Loading />
            )}
        </>
    );
}