import { TweetType } from "./tweet";

export type FavoriteType = {
    id: number;
    userId: number;
    tweetId: number;
    tweet: TweetType;
    createdAt: Datetime;
    updatedAt: Datetime;
}