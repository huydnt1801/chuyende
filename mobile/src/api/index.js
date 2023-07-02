import accountApi from "./Account"
import google from "./Google"
import tripApi from "./Trip"
import { Result } from "./baseApi"

const Api = {
    account: accountApi,
    trip: tripApi,
    google: google,
    ResultCode: Result
}

export default Api