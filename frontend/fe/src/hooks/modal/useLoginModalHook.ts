import axiosInstance from "@/lib/axiosInstance";
import { loginPatchSchema, LoginPatchSchemaType } from "@/lib/validation/login";
import { zodResolver } from "@hookform/resolvers/zod";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { useForm } from "react-hook-form"

export const useLoginModalHook = () => {
    const { instance } = axiosInstance();
    const router = useRouter();

    const [isLoading, setIsLoading] = useState<boolean>(false);

    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<LoginPatchSchemaType>({
        resolver: zodResolver(loginPatchSchema)
    });

    const onSubmit = async (data: LoginPatchSchemaType) => {
        try {
            const response = await instance.post(
                "/login",
                data
            );

            if (response.status === 200) {
                router.push("/home");
            } else {
                console.error(response);
            }
        } catch (err) {
            console.error(err);
        }
    }

    return {
        isLoading,
        setIsLoading,
        register,
        handleSubmit,
        errors,
        onSubmit
    };
}