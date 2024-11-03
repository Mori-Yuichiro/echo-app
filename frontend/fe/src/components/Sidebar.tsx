"use client"

import Link from "next/link";
import Button from "./Button";
import { useSidebarHook } from "@/hooks/useSidebarHook";

export default function Sidebar() {
    const { ITEM_LIST, currentUser } = useSidebarHook();

    return (
        <aside className="flex flex-col gap-3 h-screen max-lg:w-20 w-60 border-r border-black p-4 fixed max-sm:hidden">
            <Link href={'/home'}>
                <svg width="28" height="28" viewBox="0 0 1200 1227" fill="none" xmlns="http://www.w3.org/2000/svg">
                    <path d="M714.163 519.284L1160.89 0H1055.03L667.137 450.887L357.328 0H0L468.492 681.821L0 1226.37H105.866L515.491 750.218L842.672 1226.37H1200L714.137 519.284H714.163ZM569.165 687.828L521.697 619.934L144.011 79.6944H306.615L611.412 515.685L658.88 583.579L1055.08 1150.3H892.476L569.165 687.854V687.828Z" fill="white" />
                </svg>
            </Link>
            {ITEM_LIST.map((item, i) => {
                return (
                    <div key={`sidebar-${i}`}>
                        {item.disabled ?
                            <div className="flex items-center gap-5 hover:bg-slate-300">
                                <div dangerouslySetInnerHTML={{ __html: item.svg }} />
                                <h1 className="text-xl max-lg:hidden">{item.label}</h1>
                            </div> :
                            <Link
                                href={item.href}
                                className="flex items-center gap-5 hover:bg-slate-300"
                            >
                                <div dangerouslySetInnerHTML={{ __html: item.svg }} />
                                <h1 className="text-xl max-lg:hidden">{item.label}</h1>
                            </Link>
                        }
                    </div>
                );
            })}
            <Button disabled className="rounded-full bg-cyan-400 py-2">Post</Button>
            <div
                className="flex items-center gap-3 cursor-pointer"
            >
                <div className="bg-slate-400 w-8 h-8 rounded-full">
                    {currentUser?.image && <img className="w-full h-full rounded-full" src={currentUser.image!} alt="icon" />}
                </div>
                {currentUser?.displayName ? (
                    <p>{currentUser?.displayName}</p>
                ) : (
                    <p>{currentUser?.name}</p>
                )}
            </div>
        </aside >
    );
}