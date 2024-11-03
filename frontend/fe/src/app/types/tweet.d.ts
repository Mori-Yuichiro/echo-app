import { UserType } from "./user";

export type TweetType = {
    id: number;
    content: string;
    user: UserType;
    createdAt: Datetime;
    updatedAt: Datetime;
}
