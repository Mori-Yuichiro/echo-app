"use client"

import { useDeleteTweetModal } from "@/hooks/modal/useDeleteTweetModal";
import Button from "../Button";

export default function DeleteTweetModal({
    id,
}: {
    id: number,
}) {
    const {
        onClickDeleteTweet,
        onClickDeleteTweetModal
    } = useDeleteTweetModal(id);

    return (
        <div className="relative z-10" aria-labelledby="modal-title" role="dialog" aria-modal="true">
            <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
            <div className="fixed inset-0 z-10 w-screen overflow-y-auto max-sm:mx-auto">
                <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <div className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg max-sm:w-full">
                        <div className="bg-white space-y-2 px-4 pb-4 pt-5 sm:p-4 sm:pb-4">
                            <div className="flex flex-col items-center gap-y-6">
                                <p>このツイートを消します。よろしいでしょうか?</p>
                                <div className="w-full flex justify-center gap-x-2">
                                    <Button
                                        className="bg-cyan-400 rounded-full w-1/5 py-1"
                                        onClick={onClickDeleteTweet}
                                    >Yes</Button>
                                    <Button
                                        className="bg-red-500 rounded-full w-1/5 py-1"
                                        onClick={onClickDeleteTweetModal}
                                    >No</Button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}