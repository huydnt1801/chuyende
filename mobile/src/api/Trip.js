import axios from "axios";
import Utils from "../share/Utils";
import baseApi from "./baseApi";

const getList = () => {
    const url = "/trips";
    return baseApi.get(url, {
        headers: {
            cookie: Utils.data["cookie"]
        }
    });
}

/**
 * @typedef params
 * @property {Number} startX 
 * @property {Number} startY 
 * @property {String} startLocation 
 * @property {Number} endX 
 * @property {Number} endY 
 * @property {String} endLocation 
 * @property {Number} distance 
 * @property {Number} distance 
 * @property {String} type 
 * @param {params} params
 * @returns 
 */
const createTrip = async (params) => {
    const url = "/trips";
    const param = {
        ...params
    }
    return baseApi.post(url, param, {
        headers: {
            cookie: Utils.data["cookie"]
        }
    });
}
const tripApi = {
    getList,
    createTrip
}

export default tripApi