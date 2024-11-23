import { CommentType } from "./comment";
import { FavoriteType } from "./favorite";
import { TweetType } from "./tweet";
import { UserType } from "./user";

export type ProfileType =
    UserType &
    { tweets: TweetType[] } &
    { favorites: FavoriteType[] } &
    { comments: CommentType[] }