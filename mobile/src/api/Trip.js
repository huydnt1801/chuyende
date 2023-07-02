import axios from "axios";
import Utils from "../share/Utils";
import baseApi from "./baseApi";

const getList = () => {
    const url = "/trips";
    return baseApi.get(url, {
        headers: {
            cookie: Utils.global.cookie
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
const createTrip = (params) => {
    const url = "/trips";
    const param = {
        ...params
    }
    return baseApi.post(url, param, {
        headers: {
            cookie: Utils.global.cookie
        }
    });
}


const checkPrice = (distance) => {

    const url = "/trips/price";
    const params = {
        distance: distance
    }

    return baseApi.get(url, {
        params,
        headers: {
            cookie: Utils.global.cookie
        }
    });
}

const rate = (tripId, star) => {
    const url = `trips/${tripId}/rate`;
    const param = {
        rate: star
    }

    return baseApi.patch(url, param, {
        headers: {
            cookie: Utils.global.cookie
        }
    });
}

const updateStatus = (tripId, status) => {
    const url = `trips/${tripId}/status`;
    const param = {
        status: status
    }

    return baseApi.patch(url, param, {
        headers: {
            cookie: Utils.global.cookie
        }
    });
}

const tripApi = {
    getList,
    createTrip,
    checkPrice,
    rate,
    updateStatus
}

export default tripApi