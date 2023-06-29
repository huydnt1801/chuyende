import axios from "axios";

import env from "../../env.json";
import suggest from "../../suggest.json";
import location1 from "../../location1.json";
import locationByLatAndLong from "../../locationbylatandlong1.json";

const getSuggestPlace = async (query) => {
    if (!query) {
        return {
            status: "FAIL"
        }
    }
    const url = "https://maps.googleapis.com/maps/api/place/autocomplete/json";

    const params = {
        key: env.SUGGEST_API_KEY,
        components: "country:vn",
        language: "vi",
        input: query
    }

    return suggest;
    return axios.get(url, { params });
}

const getPosition = async (placeId) => {
    if (!placeId) {
        return {
            status: "FAIL"
        }
    }
    const url = "https://maps.googleapis.com/maps/api/place/details/json";

    const params = {
        key: env.DETAIL_API_KEY,
        place_id: placeId,
        language: "vi"
    }

    return location1;

    return axios.get(url, { params });
}

const getPositionByLatAndLong = async (latitude, longitude) => {
    if (!latitude || !longitude || latitude == 0 || longitude == 0) {
        return {
            status: "FAIL"
        }
    }

    const url = "https://maps.googleapis.com/maps/api/geocode/json";

    const params = {
        key: env.GEO_CODE_API_KEY,
        language: "vi",
        latlng: `${latitude},${longitude}`
    }

    return locationByLatAndLong;

    return axios.get(url, { params });
}

const google = {
    getSuggestPlace,
    getPosition,
    getPositionByLatAndLong
}

export default google