import { Text, View, Button, StyleSheet, TouchableOpacity, TextInput } from "react-native";
import MapView from "react-native-maps"
import className from "./className"
import { faCircleXmark, faLocationDot } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { useTranslation } from "react-i18next";
import { useState, useEffect } from "react";
import MarkerSelect from "./components/MarkerSelect";
import { useRoute } from "@react-navigation/native";
import { useDebounce } from "../../hooks";
import Api from "../../api";


const types = {
    SELECT_SOURCE: 1,
    SELECT_DESTINATION: 2
}

const SelectLocationOnMap = () => {

    const { type: type_ } = useRoute().params ?? { type: 1 };

    const [isChanging, setIsChanging] = useState(false);

    const defaultPosition = {
        latitude: 21.0285,
        longitude: 105.8542
    };

    const [location, setLocation] = useState({
        latitude: 0,
        longitude: 0
    });

    const debounce = useDebounce(location, 500);

    const { t } = useTranslation();
    const [searchInput, setSearchInput] = useState("");

    const getPlace = async () => {
        const result = await Api.google.getPositionByLatAndLong(location.latitude, location.longitude);
        console.log(JSON.stringify(result));
        if (result.status == "OK") {
            if (result.results.length > 0) {
                const tmp = result.results[0];
                setSearchInput(tmp.formatted_address ?? "");
            }
        }
    }

    useEffect(() => {
        if (debounce.latitude != 0 && debounce.longitude != 0) {
            getPlace();
        }
    }, [debounce])

    return (
        <View className={className.container}>
            <MapView
                style={StyleSheet.absoluteFillObject}
                initialRegion={{
                    latitude: defaultPosition.latitude,
                    longitude: defaultPosition.longitude,
                    latitudeDelta: 0.0522,
                    longitudeDelta: 0.0221
                }}
                showsBuildings={false}
                showsPointsOfInterest={false}
                showsUserLocation={false}
                onRegionChangeComplete={(region) => {
                    setIsChanging(false);
                    setLocation(region);
                }}
                onRegionChange={() => {
                    setIsChanging(true);
                }}
            />
            <View className={className.search}>
                <FontAwesomeIcon
                    icon={faLocationDot}
                    size={24}
                    style={{ color: "rgb(234 179 8)", marginRight: 8 }} />
                <TextInput
                    className={className.input}
                    value={searchInput}
                    placeholder={t("EnterDestination")}
                    onChangeText={text => setSearchInput(text)} />
                {searchInput.length > 0 && (
                    <TouchableOpacity
                        style={{ padding: 4 }}
                        activeOpacity={0.5}
                        onPress={() => setSearchInput("")}>
                        <FontAwesomeIcon
                            icon={faCircleXmark}
                            size={18}
                            style={{ color: "rgb(75 85 99)" }} />
                    </TouchableOpacity>
                )}
            </View>
            <MarkerSelect
                isChanging={isChanging} />
        </View>
    )
}

export default SelectLocationOnMap
export { types }