import RightSidebar from "@/components/RightSidebar";
import Sidebar from "@/components/Sidebar";

export default function HomeLayout({ children }: { children: React.ReactNode }) {
    return (
        <div className="min-h-screen px-0 lg:px-32 md:px-20 sm:px-10 flex">
            <Sidebar />
            <main className="w-2/3 ml-60 md:border-l max-md:border-r max-md:border-black max-lg:ml-20 max-sm:border-black max-sm:border-l">{children}</main>
            <RightSidebar />
        </div>
    );
}