import { CommentType } from "./comment";
import { FavoriteType } from "./favorite";
import { UserType } from "./user";

export type TweetType = {
    id: number;
    content: string;
    image_urls: string[];
    user: UserType;
    createdAt: Datetime;
    updatedAt: Datetime;
    favorites: FavoriteType[];
    comments: CommentType[];
}
