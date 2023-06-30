import Utils from "../share/Utils";
import baseApi from "./baseApi";

const getList = () => {
    const url = "/trips";
    // const param = {
    //     phoneNumber: phone,
    //     password
    // }

    console.log("===>", Utils.data["cookie"]);
    return baseApi.get(url, {
        headers: {
            // cookie: Utils.data["cookie"] + "123"
        }
    });
}

const trip = {
    getList
}

export default trip