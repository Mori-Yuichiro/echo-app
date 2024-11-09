import { UserType } from "./user";

export type TweetType = {
    id: number;
    content: string;
    image_urls: string[];
    user: UserType;
    createdAt: Datetime;
    updatedAt: Datetime;
}
