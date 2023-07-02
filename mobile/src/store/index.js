import { configureStore, applyMiddleware, compose } from "@reduxjs/toolkit";

import accountReducer from "../slices/Account";
import tripReducer from "../slices/Trip";

const reducer = {
    account: accountReducer,
    trip: tripReducer
}

const logger = (_store) => (next) => (action) => {
    return next(action);
}

const store = configureStore({
    reducer: reducer,
    devTools: true,
    enhancers: [compose(applyMiddleware(logger))],
    middleware: (getDefaultMiddleware) => getDefaultMiddleware({
        serializableCheck: false
    })
});

export default store