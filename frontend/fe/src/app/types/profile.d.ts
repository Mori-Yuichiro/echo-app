import { CommentType } from "./comment";
import { FavoriteType } from "./favorite";
import { RelationshipType } from "./relationship";
import { RetweetType } from "./retweet";
import { TweetType } from "./tweet";
import { UserType } from "./user";

export type ProfileType =
    UserType &
    { tweets: TweetType[] } &
    { favorites: FavoriteType[] } &
    { comments: CommentType[] } &
    { retweets: RetweetType[] } &
    { followers: RelationshipType[] } &
    { followeds: RelationshipType[] }