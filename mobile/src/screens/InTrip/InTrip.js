import { useRoute } from "@react-navigation/native";
import { useSelector } from "react-redux";

import DriverTrip from "./Driver/DriverTrip";
import UserTrip from "./User/UserTrip";

const InTrip = () => {

    const { tripId } = useRoute().params ?? {};
    const { isDriver } = useSelector(state => state.account);
    console.log("tripId", tripId);

    if (isDriver) {
        return <DriverTrip
            tripId={tripId} />
    }
    else {
        return <UserTrip
            tripId={tripId} />
    }
}

export default InTrip