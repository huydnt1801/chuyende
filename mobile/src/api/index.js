import accountApi from "./Account"
import google from "./Google"
import trip from "./Trip"
import { Result } from "./baseApi"

const Api = {
    account: accountApi,
    trip: trip,
    google: google,
    ResultCode: Result
}

export default Api