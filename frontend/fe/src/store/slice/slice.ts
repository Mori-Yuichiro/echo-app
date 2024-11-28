import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { UserType } from "@/app/types/user";

interface State {
    currentUser: UserType | undefined;
    reload: boolean;
    openModal: boolean;
    deleteTweetModal: boolean;
}

const initialState: State = {
    currentUser: undefined,
    reload: false,
    openModal: false,
    deleteTweetModal: false
};

const slice = createSlice({
    name: "state",
    initialState,
    reducers: {
        changeCurrentUser(state, action: PayloadAction<UserType | undefined>) {
            state.currentUser = action.payload;
        },
        toggleReload(state, action: PayloadAction<boolean>) {
            state.reload = action.payload;
        },
        toggleOpenModal(state, action: PayloadAction<boolean>) {
            state.openModal = action.payload;
        },
        toggleDeleteTweetModal(state, action: PayloadAction<boolean>) {
            state.deleteTweetModal = action.payload;
        }
    }
});

export const {
    toggleReload,
    changeCurrentUser,
    toggleOpenModal,
    toggleDeleteTweetModal
} = slice.actions;
export default slice.reducer;