import { UserType } from "./user";

export type EntryType = {
    id: number;
    user_id: number;
    room_id: number;
    user: UserType;
    created_at: Datetime;
    updated_at: Datetime;
}