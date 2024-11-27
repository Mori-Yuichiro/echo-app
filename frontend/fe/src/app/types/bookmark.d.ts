import { TweetType } from "./tweet";

export type BookmarkType = {
    id: number;
    userId: number;
    tweetId: number;
    tweet: TweetType;
    createdAt: Datetime;
    updatedAt: Datetime;
}