import axiosInstance from "@/lib/axiosInstance";
import { useEffect } from "react";
import { useError } from "./error/useError";
import { AxiosError } from "axios";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { changeCurrentUser } from "@/store/slice/slice";

type ItemListType = {
    label: string;
    svg: string;
    disabled: boolean;
    href: string;
};

export const useSidebarHook = () => {
    const { instance } = axiosInstance();
    const { switchErrorHandling } = useError();
    const currentUser = useAppSelector(state => state.slice.currentUser);
    const reload = useAppSelector(state => state.slice.reload);
    const dispatch = useAppDispatch();


    const ITEM_LIST: ItemListType[] = [
        {
            label: "Home",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 432 384"><path fill="currentColor" d="M171 363H64V192H0L213 0l214 192h-64v171H256V235h-85v128z"/></svg>',
            disabled: false,
            href: '/home'
        },
        {
            label: "Explore",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 14 14"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 11.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m7.5 2L10 10"/></svg>',
            disabled: true,
            href: '/explore'
        },
        {
            label: "Notifications",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><path fill="currentColor" d="M10.146 3.248a2 2 0 0 1 3.708 0A7.003 7.003 0 0 1 19 10v4.697l1.832 2.748A1 1 0 0 1 20 19h-4.535a3.501 3.501 0 0 1-6.93 0H4a1 1 0 0 1-.832-1.555L5 14.697V10c0-3.224 2.18-5.94 5.146-6.752zM10.586 19a1.5 1.5 0 0 0 2.829 0h-2.83zM12 5a5 5 0 0 0-5 5v5a1 1 0 0 1-.168.555L5.869 17H18.13l-.963-1.445A1 1 0 0 1 17 15v-5a5 5 0 0 0-5-5z"/></svg>',
            disabled: false,
            href: '/notifications'
        },
        {
            label: "Messages",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-width="2"><rect width="16" height="12" x="4" y="6" rx="2"/><path d="m4 9l7.106 3.553a2 2 0 0 0 1.788 0L20 9"/></g></svg>',
            disabled: false,
            href: '/rooms'
        },
        {
            label: "Grok",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.5 12c0-4.23 0-6.345 1.198-7.747q.256-.3.555-.555C5.655 2.5 7.77 2.5 12 2.5s6.345 0 7.747 1.198q.3.256.555.555C21.5 5.655 21.5 7.77 21.5 12s0 6.345-1.198 7.747q-.256.3-.555.555C18.345 21.5 16.23 21.5 12 21.5s-6.345 0-7.747-1.198q-.3-.256-.555-.555C2.5 18.345 2.5 16.23 2.5 12M8 17.5l8-10" color="currentColor"/></svg>',
            disabled: true,
            href: '/grok'
        },
        {
            label: "Bookmarks",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><path fill="currentColor" d="M5 21V3h14v18l-7-3z"/></svg>',
            disabled: false,
            href: '/bookmarks'
        },
        {
            label: "Jobs",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" color="currentColor"><path d="M11.007 21.5H9.605c-3.585 0-5.377 0-6.491-1.135S2 17.403 2 13.75s0-5.48 1.114-6.615S6.02 6 9.605 6h3.803c3.585 0 5.378 0 6.492 1.135c.857.873 1.054 2.156 1.1 4.365V13"/><path d="M19 18.5h-3m0 3a3 3 0 1 1 0-6m3 6a3 3 0 1 0 0-6M16 6l-.1-.31c-.495-1.54-.742-2.31-1.331-2.75c-.59-.44-1.372-.44-2.938-.44h-.263c-1.565 0-2.348 0-2.937.44c-.59.44-.837 1.21-1.332 2.75L7 6"/></g></svg>',
            disabled: true,
            href: '/jobs'
        },
        {
            label: "Communities",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 16 16"><path fill="currentColor" d="M2 5.5a3.5 3.5 0 1 1 5.898 2.549a5.51 5.51 0 0 1 3.034 4.084a.75.75 0 1 1-1.482.235a4 4 0 0 0-7.9 0a.75.75 0 0 1-1.482-.236A5.5 5.5 0 0 1 3.102 8.05A3.5 3.5 0 0 1 2 5.5M11 4a3.001 3.001 0 0 1 2.22 5.018a5 5 0 0 1 2.56 3.012a.749.749 0 0 1-.885.954a.75.75 0 0 1-.549-.514a3.51 3.51 0 0 0-2.522-2.372a.75.75 0 0 1-.574-.73v-.352a.75.75 0 0 1 .416-.672A1.5 1.5 0 0 0 11 5.5A.75.75 0 0 1 11 4m-5.5-.5a2 2 0 1 0-.001 3.999A2 2 0 0 0 5.5 3.5"/></svg>',
            disabled: true,
            href: '/communities'
        },
        {
            label: "Premium",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><path fill="currentColor" d="m20 7l-.95-2.05L17 4l2.05-.95L20 1l.95 2.05L23 4l-2.05.95L20 7Zm-4.45 5.7L11.3 8.45l2.875-2.875l4.25 4.25L15.55 12.7Zm4.25 9.9l-7.1-7.05l-6.875 6.875L1.6 18.15l6.85-6.85L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4Zm-13.95-3l5.45-5.5l-1.4-1.4l-5.5 5.45l1.45 1.45Zm0 0L4.4 18.15l1.45 1.45Z"/></svg>',
            disabled: true,
            href: '/premium'
        },
        {
            label: "Verified Orgs",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><g fill="none" fill-rule="evenodd"><path d="M24 0v24H0V0zM12.593 23.258l-.011.002l-.071.035l-.02.004l-.014-.004l-.071-.035q-.016-.005-.024.005l-.004.01l-.017.428l.005.02l.01.013l.104.074l.015.004l.012-.004l.104-.074l.012-.016l.004-.017l-.017-.427q-.004-.016-.017-.018m.265-.113l-.013.002l-.185.093l-.01.01l-.003.011l.018.43l.005.012l.008.007l.201.093q.019.005.029-.008l.004-.014l-.034-.614q-.005-.019-.02-.022m-.715.002a.02.02 0 0 0-.027.006l-.006.014l-.034.614q.001.018.017.024l.015-.002l.201-.093l.01-.008l.004-.011l.017-.43l-.003-.012l-.01-.01z"/><path fill="currentColor" d="m9.65 4l-3.111 7h3.447c.69 0 1.176.675.958 1.33l-1.656 4.967L16.586 10h-2.57a1.01 1.01 0 0 1-.903-1.462L15.382 4zM8.084 2.6c.162-.365.523-.6.923-.6h7.977c.75 0 1.239.79.903 1.462L15.618 8h3.358c.9 0 1.35 1.088.714 1.724L7.737 21.677c-.754.754-2.01-.022-1.672-1.033L8.613 13H5.015a1.01 1.01 0 0 1-.923-1.42z"/></g></svg>',
            disabled: true,
            href: '/orgs'
        },
        {
            label: "Profile",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><path fill="currentColor" d="M9.775 12q-.9 0-1.5-.675T7.8 9.75l.325-2.45q.2-1.425 1.3-2.363T12 4t2.575.938t1.3 2.362l.325 2.45q.125.9-.475 1.575t-1.5.675zm0-2h4.45L13.9 7.6q-.1-.7-.637-1.15T12 6t-1.263.45T10.1 7.6zM4 20v-2.8q0-.85.438-1.562T5.6 14.55q1.55-.775 3.15-1.162T12 13t3.25.388t3.15 1.162q.725.375 1.163 1.088T20 17.2V20zm2-2h12v-.8q0-.275-.137-.5t-.363-.35q-1.35-.675-2.725-1.012T12 15t-2.775.338T6.5 16.35q-.225.125-.363.35T6 17.2zm6 0"/></svg>',
            disabled: false,
            href: `/profile/${currentUser?.id}`
        },
        {
            label: "More",
            svg: '<svg xmlns="http://www.w3.org/2000/svg" width="28" height="28" viewBox="0 0 24 24"><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path fill="currentColor" d="M7 12.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m5 0a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1"/><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10"/></g></svg>',
            disabled: true,
            href: '/more'
        }
    ];

    useEffect(() => {
        const fetchData = async () => {
            try {
                const resUserId = await instance.get(
                    "/users",
                    { withCredentials: true }
                );
                const resCurrentUser = await instance.get(
                    `/users/${resUserId.data}`,
                    { withCredentials: true }
                );

                if (resCurrentUser.status === 200) dispatch(changeCurrentUser(resCurrentUser.data));
            } catch (err) {
                if (err instanceof AxiosError) {
                    switchErrorHandling(err.response?.data.message);
                }
                console.error(err);
            }
        }
        fetchData();
    }, [reload])

    return { ITEM_LIST, currentUser };
}