import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    account: null
}

const accountSlice = createSlice({
    name: "account",
    initialState: initialState,
    reducers: {},
    extraReducers: (builder) => { }
});

const accountReducer = accountSlice.reducer;
const { } = accountSlice.actions;

export default accountReducer