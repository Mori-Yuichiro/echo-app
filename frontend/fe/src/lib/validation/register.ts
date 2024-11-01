import { z } from "zod";

export const registerPatchSchema = z.object({
    name: z.string()
        .min(1, { message: "Nameは1文字以上にしてください" }),
    email: z.string()
        .min(1, { message: "Emailは1文字以上にしてください" })
        .email({ message: "Emailの形式が違います" }),
    password: z.string()
        .min(8, { message: "パスワードは8文字以上にしてください" })
        .regex(/(?=.*[0-9])(?=.*[a-z])(?=.*[A-Z])/, { message: "パスワードには大小英字と数字を含めてください" })
});

export type RegisterPatchSchemaType = z.infer<typeof registerPatchSchema>;