import { ButtonType } from "@/app/types/button";

export default function Button({ children, ...props }: ButtonType) {
    return (
        <button {...props}>{children}</button>
    );
}