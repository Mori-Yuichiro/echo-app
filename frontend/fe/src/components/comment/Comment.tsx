import { CommentType } from "@/app/types/comment";

export default function Comment({ comment }: { comment: CommentType }) {
    return (
        <div className="border-b border-gray-200 px-4 py-3 space-y-3">
            <div className="flex gap-x-3 items-center">
                <div className="bg-slate-400 rounded-full w-8 h-8">
                    {comment.user.image &&
                        <img
                            src={comment.user.image}
                            alt="アイコン"
                            className="w-full h-full rounded-full"
                        />
                    }
                </div>
                <div className="w-full">
                    <p>{comment.comment}</p>
                </div>
            </div>
        </div>
    );
}