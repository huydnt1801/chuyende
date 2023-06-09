import CheckNetwork, { checkNetworkRef } from "./components/CheckNetwork";
import ConfirmDialog, { confirmDialogRef } from "./components/ConfirmDialog";
import Loading, { loadingRef } from "./components/Loading";
import MessageDialog, { messageDialogRef } from "./components/MessageDialog";
import Toast, { toastRef } from "./components/Toast";

const wait = (ms) => new Promise(e => setTimeout(e, ms));

const Utils = {
    data: {},
    global: {
        cookie: "",
        accessToken: "",
        isDriver: false,
        token: null,
    },
    /**
     * Show a toast to screen
     * @param {String} message: message of toast
     * @param {Number | undefined} duration: time exist ms
     */
    toast: (message = "", duration = 2500) => {
        toastRef.current.hide();
        toastRef.current.show(message, duration);
    },
    /**
     * Show message dialog
     * @typedef config
     * @property {String | undefined} message
     * @property {() => void | undefined} onConfirm
     * @property {String | undefined} buttonText
     * @param {config} config 
     */
    showMessageDialog: (config) => {
        messageDialogRef.current.hide();
        messageDialogRef.current.show(config)
    },
    hideMessageDialog: () => {
        messageDialogRef.current.hide();
    },
    showCheckNetwork: () => {
        checkNetworkRef.current.hide();
        checkNetworkRef.current.show();
    },
    hideCheckNetwork: () => {
        checkNetworkRef.current.hide();
    },
    showLoading: () => {
        loadingRef.current.hide();
        loadingRef.current.show();
    },
    hideLoading: () => {
        loadingRef.current.hide();
    },
    /**
     * Show confirm dialog
     * @typedef configs
     * @property {String | undefined} message
     * @property {() => void | undefined} onConfirm
     * @property {() => void | undefined} onCancel
     * @property {String | undefined} buttonTextLeft
     * @property {String | undefined} buttonTextRight
     * @param {configs} config 
     */
    showConfirmDialog: (config = {}) => {
        confirmDialogRef.current.hide();
        confirmDialogRef.current.show(config);
    },
    hideConfirmDialog: () => {
        confirmDialogRef.current.hide();
    },
    /**
     * wait ms before next action
     * @param {Number} ms 
     */
    wait: async (ms) => {
        await wait(ms);
    },
}

const UtilComponents = () => {

    return (
        <>
            <Toast />
            <MessageDialog />
            <CheckNetwork />
            <Loading />
            <ConfirmDialog />
        </>
    );
}

export default Utils
export { UtilComponents }