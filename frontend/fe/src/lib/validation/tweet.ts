import { z } from "zod";

export const tweetPatchSchema = z.object({
    content: z.string()
        .min(1, { message: "ツイートは1文字以上にしてください" })
        .max(140, { message: "ツイートは140文字以内にしてください" })
});

export type TweetPatchSchemaType = z.infer<typeof tweetPatchSchema>;