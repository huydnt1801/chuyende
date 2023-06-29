import accountApi from "./Account"
import google from "./Google"
import { Result } from "./baseApi"

const Api = {
    account: accountApi,
    google: google,
    ResultCode: Result
}

export default Api