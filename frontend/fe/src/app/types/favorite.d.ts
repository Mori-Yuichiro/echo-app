import { TweetType } from "./tweet";

export type FavoriteType = {
    id: number;
    user_id: number;
    tweet_id: number;
    tweet: TweetType;
    created_at: Datetime;
    updated_at: Datetime;
}