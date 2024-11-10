import { TweetType } from "./tweet";
import { UserType } from "./user";

export type ProfileType =
    UserType &
    { tweets: TweetType[] };