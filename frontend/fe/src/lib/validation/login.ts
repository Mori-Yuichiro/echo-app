import { z } from "zod";

const PW_REGEXP = new RegExp("^[a-z0-9]$");

export const loginPatchSchema = z.object({
    email: z.string()
        .min(1, { message: "Emailは1文字以上にしてください" })
        .email({ message: "Emailの形式が違います" }),
    password: z.string()
        .min(8, { message: "パスワードは8文字以上にしてください" })
        .regex(PW_REGEXP, { message: "パスワードには英字と数字を含めてください" })
});

export type LoginPatchSchemaType = z.infer<typeof loginPatchSchema>;