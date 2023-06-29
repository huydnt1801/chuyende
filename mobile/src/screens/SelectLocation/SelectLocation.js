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
import { useNavigation } from "@react-navigation/native";
import { TextInput } from "react-native";
import { useState, useEffect } from "react";
import { useDebounce } from "../../hooks";
import Header from "../../components/Header";
import Api from "../../api";

const types = {
    SELECT_SOURCE: 1,
    SELECT_DESTINATION: 2
}

const data_ = [
    {
        id: 1,
        text: '121 Lê Lợi',
        value: "121 Lê Lợi, Nguyễn Trãi, Hà Đông, Hà Nội",
        distance: "5.9KM"
    },
    {
        id: 2,
        text: '121 Lê Lợi',
        value: "121 Lê Lợi, Nguyễn Trãi, Hà Đông, Hà Nội",
        distance: "5.9KM"
    },
    {
        id: 3,
        text: '121 Lê Lợi',
        value: "121 Lê Lợi, Nguyễn Trãi, Hà Đông, Hà Nội 12312  2123 2",
        distance: "5.9KM"
    },
    {
        id: 4,
        text: '121 Lê Lợi',
        value: "121 Lê Lợi, Nguyễn Trãi, Hà Đông, Hà Nội sadasdsadsad",
        distance: "5.9KM"
    },
    {
        id: 5,
        text: '121 Lê Lợi',
        value: "121 Lê Lợi, Nguyễn Trãi, Hà Đông, Hà Nội",
        distance: "5.9KM"
    },
]

const SelectLocation = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();

    const [data, setData] = useState([
        ...data_
    ]);

    const [suggestPlaces, setSuggestPlaces] = useState([]);

    const [searchInput, setSearchInput] = useState("");
    const debounce = useDebounce(searchInput, 500);

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
                title={t("Destination")}
                onPressBack={() => navigation.goBack()}
                componentRight={
                    <TouchableOpacity
                        activeOpacity={0.5}
                        onPress={() => navigation.navigate("SelectLocationOnMap")}>
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
            <View className={className.placeSuggest}>
                {suggestPlaces.slice(0, 5).map((item, index) => (
                    <Pressable
                        key={item.id ?? index}
                        className={className.item}>
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