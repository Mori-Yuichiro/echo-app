import { z } from "zod";

export const messagePatchSchema = z.object({
    message: z.string()
        .min(1, { message: "メッセージは1文字以上にしてください" })
        .max(140, { message: "メッセージは140文字以内にしてください" })
});

export type MessagePatchSchemaType = z.infer<typeof messagePatchSchema>;