import Button from "./Button";

export default function RightSidebar() {
    return (
        <div className="w-60 flex flex-col gap-5 border-l border-black px-5 py-4 max-md:hidden">
            <input
                type="text"
                placeholder="Search"
                className="bg-gray-300 px-4 py-2 rounded-full"
            />
            <div className="flex flex-col gap-2 border border-black rounded-2xl p-2">
                <h1 className="font-bold text-lg">Subscribe to Premium</h1>
                <p className="text-sm">Subscribe to unlock new features and if eligible, receive a share of ads revenue.</p>
                <Button
                    disabled
                    className="bg-blue-300 rounded-full p-2 text-sm w-3/5"
                >Subscribe</Button>
            </div>
            <div className="flex flex-col gap-3 border border-black rounded-lg p-2">
                <h1 className="text-xl">What's happening</h1>
                <p>Ruby</p>
                <p>Ruby</p>
                <p>Ruby</p>
                <p>Ruby</p>
                <p>Ruby</p>
            </div>
            <div className="flex flex-col gap-3 border border-black rounded-lg p-2">
                <p>Testさん</p>
                <p>Testさん</p>
                <p>Testさん</p>
                <p>Testさん</p>
            </div>
        </div>
    );
}