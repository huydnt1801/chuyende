import {
    FlatList,
    RefreshControl,
    Text,
    View
} from "react-native";
import { useTranslation } from "react-i18next";
import className from "./className";
import { useState, useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { thunkGetListTrip } from "../../../../slices/Trip";
import TripItem from "../../TripItem";

const History = () => {

    const { t } = useTranslation();
    const dispatch = useDispatch();
    const { listTrip } = useSelector(state => state.trip);
    const trips = listTrip.filter(i => i.status == "done")

    const [isLoading, setIsLoading] = useState(false);

    const handleRefresh = async () => {
        setIsLoading(true);
        dispatch(thunkGetListTrip());
        setIsLoading(false);
    }

    console.log(JSON.stringify(listTrip));
    useEffect(() => {
        dispatch(thunkGetListTrip());
    }, [])

    return (
        <View className={className.container}>
            <View>
                <Text style={{ fontWeight: "bold", fontSize: 20, margin: 10, color: "rgb(234 179 8)" }} > {t("TripHistory")} </Text>
            </View>
            <FlatList
                data={trips}
                contentContainerStyle={{
                    paddingHorizontal: 8,
                    paddingTop: 4,
                    paddingBottom: 20
                }}
                ListEmptyComponent={
                    <View className="items-center justify-center py-3">
                        <Text className="font-semibold">{t("NowThereAreNoTrip")}</Text>
                    </View>
                }
                refreshControl={
                    <RefreshControl
                        tintColor={"#C4C7C8"}
                        refreshing={isLoading}
                        onRefresh={handleRefresh} />
                }
                renderItem={({ item }) => (
                    <TripItem
                        vehicle={item.type}
                        startLocation={item.startLocation}
                        endLocation={item.endLocation}
                        distance={item.distance}
                        price={item.price}
                        time={item.createdAt}
                        onPressAccept={() => 1}
                        history={true}
                    />
                )} />
        </View>
    );
}

export default History