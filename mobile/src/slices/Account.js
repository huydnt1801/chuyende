import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    account: null,
    isDriver: false
}

const accountSlice = createSlice({
    name: "account",
    initialState: initialState,
    reducers: {
        setAccount(state, { payload }) {
            state.account = payload
        },
        setIsDriver(state, { payload }) {
            state.isDriver = payload
        }
    },
    extraReducers: (builder) => { }
});

const { reducer: accountReducer, actions } = accountSlice;
const { setAccount, setIsDriver } = actions;

export default accountReducer
export { setAccount, setIsDriver }