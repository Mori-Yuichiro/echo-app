import { RoomType } from "@/app/types/room";
import { useAppSelector } from "@/store/hooks"

export const useRoomHook = () => {
    const currentUser = useAppSelector(state => state.slice.currentUser);

    return { currentUser };
}