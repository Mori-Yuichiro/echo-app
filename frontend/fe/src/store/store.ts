import { configureStore } from "@reduxjs/toolkit";
import sliceReducer from "./slice/slice";


export const store = configureStore({
    reducer: {
        slice: sliceReducer
    }
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;