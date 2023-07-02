import { useTranslation } from "react-i18next";
import {
    Pressable,
    Text,
    TouchableOpacity,
    View
} from "react-native";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faArrowLeftLong, faCircleXmark, faLocationDot } from "@fortawesome/free-solid-svg-icons";

import className from "./className";
import { useNavigation, useRoute } from "@react-navigation/native";
import { TextInput } from "react-native";
import { useState, useEffect } from "react";
import { useDebounce } from "../../hooks";
import Header from "../../components/Header";
import Api from "../../api";
import { mapTypes } from "../SelectLocationOnMap";
import { useDispatch } from "react-redux";
import { setDestination, setSource } from "../../slices/Trip";

const locationTypes = {
    SELECT_SOURCE: 1,
    SELECT_DESTINATION: 2
}

const SelectLocation = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();
    const dispatch = useDispatch();

    const { type: type_ } = useRoute().params ?? { type: 1 };

    const [suggestPlaces, setSuggestPlaces] = useState([]);
    const [searchInput, setSearchInput] = useState("");
    const debounce = useDebounce(searchInput, 300);

    const getSuggestPlace = async () => {
        const result = await Api.google.getSuggestPlace(debounce);
        if (result.status == "OK") {
            setSuggestPlaces(result.predictions);
        }
    }

    useEffect(() => {
        if (debounce.trim()) {
            getSuggestPlace();
        }
        else {
            setSuggestPlaces([]);
        }
    }, [debounce]);

    return (
        <View className={className.container}>
            <Header
                title={type_ == locationTypes.SELECT_SOURCE ? t("Source") : t("Destination")}
                onPressBack={() => navigation.goBack()}
                componentRight={
                    <TouchableOpacity
                        activeOpacity={0.5}
                        onPress={() => {
                            if (type_ == locationTypes.SELECT_DESTINATION) {
                                navigation.navigate("SelectLocationOnMap", { type: mapTypes.SELECT_DESTINATION })
                            }
                            else {
                                navigation.navigate("SelectLocationOnMap", { type: mapTypes.SELECT_SOURCE })
                            }
                        }}>
                        <Text className={className.map}>{t("SelectFromMap")}</Text>
                    </TouchableOpacity>
                } />
            <View className={className.search}>
                <FontAwesomeIcon
                    icon={faLocationDot}
                    size={24}
                    style={{ color: "rgb(234 179 8)", marginRight: 8 }} />
                <TextInput
                    className={className.input}
                    value={searchInput}
                    placeholder={type_ == locationTypes.SELECT_SOURCE ? t("EnterSource") : t("EnterDestination")}
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
            <View className={className.placeSuggest}>
                {suggestPlaces.map((item, index) => (
                    <Pressable
                        key={item.id ?? index}
                        className={className.item}
                        onPress={() => {
                            if (type_ == locationTypes.SELECT_DESTINATION) {
                                setSearchInput(item.structured_formatting?.main_text ?? item.description ?? "");
                                dispatch(setDestination({
                                    latitude: item.latitude,
                                    longitude: item.longitude,
                                    id: item.id,
                                    place: item.structured_formatting?.main_text ?? "",
                                    description: item.description
                                }));
                                navigation.push("SelectLocation", { type: locationTypes.SELECT_SOURCE });
                            }
                            else {
                                setSearchInput(item.structured_formatting?.main_text ?? item.description ?? "");
                                dispatch(setSource({
                                    latitude: item.latitude,
                                    longitude: item.longitude,
                                    id: item.id,
                                    place: item.structured_formatting?.main_text ?? "",
                                    description: item.description
                                }));
                                navigation.push("TripDirection");
                            }
                        }}>
                        <View className={className.iconBorder}>
                            <FontAwesomeIcon
                                icon={faLocationDot} />
                        </View>
                        <View className={className.center}>
                            <Text className={className.textBold}>
                                {item.structured_formatting?.main_text}
                            </Text>
                            <Text
                                className={className.textLight}
                                numberOfLines={1}
                                lineBreakMode={"tail"}>
                                {item.description}
                            </Text>
                        </View>
                        <Text className={className.distance}>
                            {item.distance ?? `${(Math.random() * 10).toString().slice(0, 3)}KM`}
                        </Text>
                    </Pressable>
                ))}
            </View>
        </View>
    );
}

export default SelectLocation
export { locationTypes }