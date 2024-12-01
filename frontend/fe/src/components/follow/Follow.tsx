import { UserType } from "@/app/types/user";
import Link from "next/link";

export default function Follow({ follow }: { follow: UserType }) {
    return (
        <Link href={`/profile/${follow.id}`}>
            <div className="border-b border-gray-200 px-4 py-3 space-y-3">
                <div className="flex gap-x-3 items-center">
                    <div className="bg-slate-400 rounded-full w-8 h-8">
                        {follow.image &&
                            <img
                                src={follow.image}
                                alt="アイコン"
                                className="w-full h-full rounded-full"
                            />
                        }
                    </div>
                    <div className="w-full">
                        {follow.display_name ? (
                            <p>{follow.display_name}</p>
                        ) : (
                            <p>{follow.name}</p>
                        )}
                        <p>{follow.bio}</p>
                    </div>
                </div>
            </div>
        </Link>
    );
}