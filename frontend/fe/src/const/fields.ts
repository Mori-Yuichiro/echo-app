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
    displayName: "",
    phoneNumber: "",
    bio: "",
    location: "",
    website: "",
    birthday: "",
    profileImageUrl: "",
};

export const fields: fieldType = {
    user: defaultUser
} as const;