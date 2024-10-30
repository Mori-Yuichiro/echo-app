import { loginPatchSchema, LoginPatchSchemaType } from "@/lib/validation/login";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form"

export const useLoginModalHook = () => {
    const {
        register,
        handleSubmit,
        formState: { errors }
    } = useForm<LoginPatchSchemaType>({
        resolver: zodResolver(loginPatchSchema)
    });
}