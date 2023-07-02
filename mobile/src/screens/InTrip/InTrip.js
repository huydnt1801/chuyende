import { useSelector } from "react-redux"
import DriverTrip from "./Driver/DriverTrip";
import UserTrip from "./User/UserTrip";

const InTrip = () => {
    const { isDriver } = useSelector(state => state.account);

    if (isDriver) {
        return <DriverTrip />
    }
    else {
        return <UserTrip />
    }
}

export default InTrip