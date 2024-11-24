import { TweetType } from "./tweet";

export type RetweetType = {
    id: number;
    userId: number;
    tweetId: number;
    tweet: TweetType
    createdAt: Datetime;
    updatedAt: Datetime;
}