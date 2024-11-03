import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { UserType } from "@/app/types/user";

interface State {
    currentUser: UserType | undefined;
    reload: boolean;
}

const initialState: State = {
    currentUser: undefined,
    reload: false
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
        }
    }
});

export const { toggleReload, changeCurrentUser } = slice.actions;
export default slice.reducer;