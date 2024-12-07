import { UserType } from "./user";

export type MessageType = {
    id: number;
    user_id: number;
    room_id: number;
    message: string;
    user: UserType;
    created_at: Datetime;
    updated_at: Datetime;
}