"use client"

import Button from "@/components/Button";
import Loading from "@/components/Loading";
import MessageModal from "@/components/modal/MessageModal";
import { useMessages } from "@/hooks/rooms/useMessages";

export default function MessagesPage() {
    const {
        router,
        currentUser,
        entry,
        messages,
        messageModal,
        changeMessageModal
    } = useMessages();

    return (
        <>
            {(entry && messages) ? (
                <>
                    <div className="min-h-screen">
                        <div className="flex justify-between border-b border-gray-200 p-2">
                            <div className="flex items-center gap-x-3 px-3">
                                <div
                                    className="cursor-pointer"
                                    onClick={() => router.back()}
                                >
                                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 256 256"><path fill="currentColor" d="M224 128a8 8 0 0 1-8 8H59.31l58.35 58.34a8 8 0 0 1-11.32 11.32l-72-72a8 8 0 0 1 0-11.32l72-72a8 8 0 0 1 11.32 11.32L59.31 120H216a8 8 0 0 1 8 8Z" /></svg>
                                </div>
                                <h1 className="font-bold text-lg">
                                    {entry.user.display_name ? entry.user.display_name : entry.user.name}
                                </h1>
                            </div>
                            <Button
                                className="bg-cyan-400 text-white rounded-full px-3 text-sm"
                                onClick={changeMessageModal}
                            >Message</Button>
                        </div>
                        {messages.map(message => (
                            <div
                                className="p-3"
                                key={message.id}>
                                {(message.user_id === currentUser?.id) ? (
                                    <div className="ml-auto w-1/3">
                                        <p className="p-3 w-full bg-cyan-400 text-white rounded-lg text-wrap">{message.message}</p>
                                    </div>
                                ) : (
                                    <div className="mr-auto w-1/3">
                                        <p className="p-3 w-full bg-slate-400 text-white rounded-lg text-wrap">{message.message}</p>
                                    </div>
                                )}
                            </div>
                        ))}
                    </div>
                    {messageModal && <MessageModal />}
                </>
            ) : (
                <Loading />
            )}
        </>
    );
}