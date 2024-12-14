import { NotificationType } from "@/app/types/notification";

export default function Notification({ notification }: { notification: NotificationType }) {
    return (
        <div className="border-b border-gray-200 px-4 py-3 space-y-3">
            <div className="flex flex-col gap-y-3">
                {(notification.action === "retweet") ? (
                    <>
                        <div className="flex gap-x-2">
                            <div
                                className="my-auto"
                            >
                                <svg className="text-green-300" xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 42 42"><path fill="currentColor" d="m24.5 30.5l7.96 7.371L40.5 30.5h-5V9.549c0-2.5-.561-3.049-3-3.049h-18l6.641 6H29.5v18h-5zm-8-18L8.52 5.16L.5 12.5h5v21.049c0 2.5.62 2.951 3 2.951h18.32l-6.32-6h-9v-18h5z" /></svg>
                            </div>
                            <div className="bg-slate-400 rounded-full w-8 h-8">
                                {notification.visitor.image &&
                                    <img
                                        src={notification.visitor.image}
                                        alt="アイコン"
                                        className="w-full h-full rounded-full"
                                    />
                                }
                            </div>
                        </div>
                        <div className="w-full ml-2">
                            <p>{notification.visitor.display_name ? notification.visitor.display_name : notification.visitor.name}があなたの投稿をリツイートしました。</p>
                        </div>
                    </>
                ) : (notification.action === "favorite") ? (
                    <>
                        <div className="flex gap-x-2">
                            <div
                                className="text-red-400 my-auto"
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 20 20"><path fill="currentColor" d="m10 3.22l-.61-.6a5.5 5.5 0 0 0-7.78 7.77L10 18.78l8.39-8.4a5.5 5.5 0 0 0-7.78-7.77l-.61.61z" /></svg>
                            </div>
                            <div className="bg-slate-400 rounded-full w-8 h-8">
                                {notification.visitor.image &&
                                    <img
                                        src={notification.visitor.image}
                                        alt="アイコン"
                                        className="w-full h-full rounded-full"
                                    />
                                }
                            </div>
                        </div>
                        <div className="w-full ml-2">
                            <p>{notification.visitor.display_name ? notification.visitor.display_name : notification.visitor.name}があなたの投稿にいいねをしました。</p>
                        </div>
                    </>
                ) : (
                    <>
                        <div className="flex gap-x-2">
                            <div
                                className="my-auto"
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 1024 1024"><path fill="currentColor" d="M512 896q-66 0-134-16q-34 40-69.5 69.5t-60 43.5t-47.5 21.5t-30.5 8.5t-10.5 1q26-57 30-124.5T176 786Q94 723 47 635T0 448q0-91 40.5-174t109-143T313 35.5T512 0t199 35.5T874.5 131t109 143t40.5 174t-40.5 174t-109 143T711 860.5T512 896z" /></svg>
                            </div>
                            <div className="bg-slate-400 rounded-full w-8 h-8">
                                {notification.visitor.image &&
                                    <img
                                        src={notification.visitor.image}
                                        alt="アイコン"
                                        className="w-full h-full rounded-full"
                                    />
                                }
                            </div>
                        </div>
                        <div className="w-full ml-2">
                            <p>{notification.visitor.display_name ? notification.visitor.display_name : notification.visitor.name}があなたの投稿にコメントをしました。</p>
                        </div>
                    </>
                )}
            </div>
        </div>
    );
}