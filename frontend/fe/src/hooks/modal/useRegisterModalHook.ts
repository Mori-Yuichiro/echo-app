import axiosInstance from "@/lib/axiosInstance";
import { registerPatchSchema, RegisterPatchSchemaType } from "@/lib/validation/register";
import { zodResolver } from "@hookform/resolvers/zod";
import { AxiosError } from "axios";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { useForm } from "react-hook-form"
import { useError } from "../error/useError";

export const useRegisterModalHook = () => {
    const [errorMsg, setErrorMsg] = useState<string | null>(null);
    const { instance } = axiosInstance();
    const router = useRouter();
    const { switchErrorHandling } = useError();

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<RegisterPatchSchemaType>({
        resolver: zodResolver(registerPatchSchema)
    });

    const onSubmit = async (data: RegisterPatchSchemaType) => {
        try {
            const resSignup = await instance.post(
                "/signup",
                data
            );

            if (resSignup.status === 201) {
                const { email, password } = data;

                const resLogin = await instance.post(
                    "/login",
                    {
                        email,
                        password
                    }
                );
                if (resLogin.status === 200) router.push("/home");
            } else {
                console.error(resSignup);
            }
        } catch (err) {
            if (err instanceof AxiosError) {
                setErrorMsg(
                    switchErrorHandling(err.response?.data)
                );
            }
            console.error(err);
        }
    }

    return {
        errorMsg,
        register,
        handleSubmit,
        errors,
        onSubmit
    };
}