import { usePathname } from "next/navigation"

export const useTweetHook = () => {
    const pathName = usePathname();

    return { pathName };
}