import axiosInstance from "@/lib/axiosInstance"
import { getCsrfToken } from "@/lib/csrf_lib";
import { messagePatchSchema, MessagePatchSchemaType } from "@/lib/validation/message";
import { useAppDispatch, useAppSelector } from "@/store/hooks";
import { toggleMessageModal, toggleReload } from "@/store/slice/slice";
import { zodResolver } from "@hookform/resolvers/zod";
import { useParams } from "next/navigation";
import { useForm } from "react-hook-form";

export const useMessageModalHook = () => {
    const { instance } = axiosInstance();

    const reload = useAppSelector(state => state.slice.reload);
    const messageModal = useAppSelector(state => state.slice.messageModal);
    const dispatch = useAppDispatch();

    const { id } = useParams<{ id: string }>();

    const {
        register,
        handleSubmit,
        formState: { errors },
        reset
    } = useForm<MessagePatchSchemaType>({
        resolver: zodResolver(messagePatchSchema)
    });

    const changeMessageModal = () => {
        dispatch(toggleMessageModal(!messageModal));
    }

    const onClickSendMessage = async (data: MessagePatchSchemaType) => {
        try {
            if (instance.defaults.headers.common["X-CSRF-Token"] === undefined) {
                const csrf = await getCsrfToken();
                instance.defaults.headers.common["X-CSRF-Token"] = csrf;
            }

            const { status } = await instance.post(
                `/room/${id}/message`,
                data,
                { withCredentials: true }
            );

            if (status === 200) {
                dispatch(toggleReload(!reload));
                dispatch(toggleMessageModal(!messageModal));
            }
        } catch (err) {
            reset({ message: "" });
            console.error(err);
        }
    }

    return {
        register,
        handleSubmit,
        errors,
        changeMessageModal,
        onClickSendMessage
    }
}