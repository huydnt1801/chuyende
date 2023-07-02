import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import Api from "../api";

const initialState = {
    source: {},
    destination: {},
    listTrip: []
}

const thunkGetListTrip = createAsyncThunk(
    "thunkGetListTrip",
    async () => {
        return await Api.trip.getList();
    }
)

const tripSlice = createSlice({
    name: "trip",
    initialState: initialState,
    reducers: {
        setSource(state, { payload }) {
            state.source = payload
        },
        setDestination(state, { payload }) {
            state.destination = payload
        },
        clearTrip(state, { }) {
            state.destination = {}
            state.source = {}
        },
    },
    extraReducers: (builder) => {
        builder.addCase(thunkGetListTrip.fulfilled, (state, action) => {
            if (action.payload.result == Api.ResultCode.SUCCESS) {
                state.listTrip = action.payload.data.data
            }
        })
    }
});

const { reducer: tripReducer, actions } = tripSlice;
const { setSource, setDestination, clearTrip } = actions;

export default tripReducer
export { setSource, setDestination, clearTrip, thunkGetListTrip }