import { EntryType } from "./entry";
import { MessageType } from "./message";

export type RoomType = {
    id: number;
    entries: EntryType[];
    messages: MessageType[];
    created_at: Datetime;
    updated_at: Datetime;
}