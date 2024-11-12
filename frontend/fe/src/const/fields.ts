import { UserType } from "@/app/types/user"

type fieldType = {
    user: UserType;
}

const defaultUser: UserType = {
    id: 0,
    email: "",
    password: "",
    name: "",
    image: "",
    display_name: "",
    phone_number: "",
    bio: "",
    location: "",
    website: "",
    birthday: "",
    profile_image_url: "",
};

export const fields: fieldType = {
    user: defaultUser
} as const;