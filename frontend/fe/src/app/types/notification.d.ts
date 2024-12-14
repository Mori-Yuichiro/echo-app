import { UserType } from "./user";

export type NotificationType = {
    id: number;
    visitorId: number;
    visitedId: number;
    tweetId: number;
    action: "comment" | "favorite" | "retweet";
    read: boolean;
    visitor: UserType;
}