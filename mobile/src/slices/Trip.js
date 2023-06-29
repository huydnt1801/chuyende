import { createSlice } from "@reduxjs/toolkit";

const initialState = {
    source: {},
    destination: {}
}

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
        }
    }
});

const { reducer: tripReducer, actions } = tripSlice;
const { setSource, setDestination, clearTrip } = actions;

export default tripReducer
export { setSource, setDestination, clearTrip }