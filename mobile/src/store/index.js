import { configureStore, applyMiddleware, compose } from "@reduxjs/toolkit";

const reducer = {

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