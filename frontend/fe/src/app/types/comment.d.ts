import { UserType } from "./user";

export type CommentType = {
    id: number;
    comment: string;
    user_id: number;
    tweet_id: number;
    user: UserType;
    created_at: Datetime;
    updated_at: Datetime;
}