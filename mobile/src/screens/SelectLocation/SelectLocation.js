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
import { useState } from "react";
import { set } from "lodash";


const SelectLocation = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();

    const [searchInput, setSearchInput] = useState("");

    return (
        <View className={className.container}>
            <View className={className.header}>
                <View className={className.left}>
                    <TouchableOpacity
                        activeOpacity={0.5}
                        className={className.back}
                        onPress={() => navigation.goBack()}>
                        <FontAwesomeIcon
                            icon={faArrowLeftLong}
                            size={24} />
                    </TouchableOpacity>
                    <Text className={className.destination}>{t("Destination")}</Text>
                </View>
                <TouchableOpacity
                    activeOpacity={0.5}>
                    <Text className={className.map}>{t("SelectFromMap")}</Text>
                </TouchableOpacity>
            </View>
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
        </View>
    );
}

export default SelectLocation