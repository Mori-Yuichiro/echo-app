import { z } from "zod";

export const commentPatchSchema = z.object({
    comment: z.string()
        .min(1, { message: "コメントは1文字以上にしてください" })
        .max(140, { message: "コメントは140文字以内にしてください" })
});

export type CommentPatchSchemaType = z.infer<typeof commentPatchSchema>