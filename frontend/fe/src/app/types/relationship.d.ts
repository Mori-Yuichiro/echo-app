export type RelationshipType = {
    id: number;
    follower_id: number;
    followed_id: number;
    createdAt: Datetime;
    updatedAt: Datetime;
}