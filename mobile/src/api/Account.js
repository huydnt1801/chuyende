import Utils from "../share/Utils";
import baseApi from "./baseApi";

const login = (phone, password) => {
    const url = "/accounts/login";
    const param = {
        phoneNumber: phone,
        password
    }
    return baseApi.post(url, param);
}

const logout = (phone) => {
    const url = "/accounts/logout";
    const param = {
        phoneNumber: phone,
    }
    return baseApi.post(url, param, {
        headers: {
            cookie: Utils.global.cookie
        }
    });
}
const confirmOTP = (token, otp) => {
    const url = "/accounts/confirm";
    const param = {
        token,
        type: "submit-otp",
        otp
    }
    return baseApi.post(url, param);
}
const resendOTP = (phone) => {
    const url = "/accounts/resend";
    const param = {
        phoneNumber: phone
    }
    return baseApi.post(url, param);
}

/**
 * 
 * @param {String} phone
 * @param {String} username 
 * @param {String} password 
 * @returns {Promise<void>}
 */
const register = (phone, username, password) => {
    const url = "/accounts/register";
    const param = {
        phoneNumber: phone,
        fullName: username,
        password: password,
        password2: password
    }
    return baseApi.post(url, param);
}

/**
 * 
 * @param {String} phone 
 * @returns {Promise<void>}
 */
const checkPhone = (phone) => {
    const url = "/accounts/phone";
    const param = {
        phoneNumber: phone
    }
    return baseApi.post(url, param);
}

const loginDriver = async (phone, password) => {
    const url = "/accounts/login/driver";
    const param = {
        phoneNumber: phone,
        password: password
    }

    return baseApi.post(url, param)

}

const accountApi = {
    login,
    checkPhone,
    register,
    confirmOTP,
    loginDriver,
    resendOTP,
    logout
}

export default accountApi