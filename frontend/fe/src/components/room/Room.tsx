"use client"

import { RoomType } from "@/app/types/room";
import { useRoomHook } from "@/hooks/rooms/useRoomHook";
import Link from "next/link";

export default function Room({ room }: { room: RoomType }) {
    const { currentUser } = useRoomHook();

    return (
        <>
            {room.entries.map(entry => (
                <div key={entry.id}>
                    {(entry.user_id !== currentUser?.id) && (
                        <Link href={`/rooms/${room.id}`}>
                            <div className="border-b border-gray-200 px-4 py-3 space-y-3">
                                <div className="flex gap-x-3 items-center">
                                    <div className="bg-slate-400 rounded-full w-8 h-8">
                                        {entry.user.image &&
                                            <img
                                                src={entry.user.image}
                                                alt="アイコン"
                                                className="w-full h-full rounded-full"
                                            />
                                        }
                                    </div>
                                    <div className="w-full">
                                        {entry.user.display_name ? (
                                            <p>{entry.user.display_name}</p>
                                        ) : (
                                            <p>{entry.user.name}</p>
                                        )}
                                        <p className="text-slate-400">{room.messages[room.messages.length - 1].message}</p>
                                    </div>
                                </div>
                            </div>
                        </Link>
                    )}
                </div>
            ))}
        </>
    );
}