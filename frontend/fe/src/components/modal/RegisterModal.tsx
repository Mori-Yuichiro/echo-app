import { Dispatch, SetStateAction } from "react";
import Button from "../Button";

export default function LoginModal({
    openRegisterModal,
    setOpenRegisterModal
}: {
    openRegisterModal: boolean,
    setOpenRegisterModal: Dispatch<SetStateAction<boolean>>
}) {
    return (
        <div className="relative z-10" aria-labelledby="modal-title" role="dialog" aria-modal="true">
            <div className="fixed inset-0 bg-gray-500 bg-opacity-75 transition-opacity" aria-hidden="true"></div>
            <div className="fixed inset-0 z-10 w-screen overflow-y-auto max-sm:mx-auto">
                <div className="flex min-h-full items-end justify-center p-4 text-center sm:items-center sm:p-0">
                    <div className="relative transform overflow-hidden rounded-lg bg-white text-left shadow-xl transition-all sm:my-8 sm:w-full sm:max-w-lg max-sm:w-full">
                        <div className="bg-white space-y-2 px-4 pb-4 pt-5 sm:p-4 sm:pb-4">
                            <div className="flex justify-between">
                                <div
                                    className="cursor-pointer flex gap-x-5 items-center"
                                    onClick={() => setOpenRegisterModal(!openRegisterModal)}
                                >
                                    <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 16 16"><path fill="currentColor" d="M7.293 8L3.146 3.854a.5.5 0 1 1 .708-.708L8 7.293l4.146-4.147a.5.5 0 0 1 .708.708L8.707 8l4.147 4.146a.5.5 0 0 1-.708.708L8 8.707l-4.146 4.147a.5.5 0 0 1-.708-.708z" /></svg>
                                    <h1 className="text-lg">New Account</h1>
                                </div>
                            </div>
                            <div className="flex flex-col gap-y-3">
                                <input
                                    id="name"
                                    className="border border-slate-400 rounded-sm p-3"
                                    type="text"
                                    placeholder="Name"
                                />
                                <input
                                    id="email"
                                    className="border border-slate-400 rounded-sm p-3"
                                    type="email"
                                    placeholder="email"
                                />
                                <input
                                    id="password"
                                    className="border border-slate-400 rounded-sm p-3"
                                    type="password"
                                    placeholder="password"
                                />
                            </div>
                            <Button
                                className="border rounded-full px-3 w-full bg-blue-400 hover:bg-cyan-600 hover:text-white"
                            >Register</Button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}