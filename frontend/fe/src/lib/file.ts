import { AxiosInstance } from "axios";
import { MutableRefObject } from "react";

export const fileRead = async (file: File) => {
    const fileReader = new FileReader();
    fileReader.readAsDataURL(file);
    await new Promise<void>((resolve) => (fileReader.onload = () => resolve()));
    return fileReader.result as string;
};

export const fileUpload = (inputRef: MutableRefObject<HTMLInputElement | null>) => {
    return () => {
        if (inputRef.current) {
            inputRef.current.click();
        }
    }
}

export const uploadImage = async (instance: AxiosInstance, imageData: string | ArrayBuffer | null) => {
    return await instance.post('api/image-upload', {
        imageData
    });
}