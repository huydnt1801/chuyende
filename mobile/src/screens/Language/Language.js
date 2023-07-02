import { useTranslation } from "react-i18next";
import { useNavigation } from "@react-navigation/native";
import {
    ScrollView, View
} from "react-native";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faCheck } from "@fortawesome/free-solid-svg-icons";

import Header from "../../components/Header";
import className from "./className";
import ButtonRow from "./ButtonRow";

const Language = () => {

    const navigation = useNavigation();
    const { i18n } = useTranslation();
    const language = i18n.language;
    const { t } = useTranslation();
    return (
        <View className={className.container}>
            <View >
                <Header
                    title={t("SelectLanguage")}
                    onPressBack={() => navigation.goBack()} />
            </View>
            <ScrollView>

                <ButtonRow
                    classNames={`mt-2`}
                    title={t("English")}
                    onPress={() => i18n.changeLanguage("en")}
                    iconRight={
                        <FontAwesomeIcon
                            icon={faCheck}
                            size={18}
                            style={{
                                color: language == "en" ? "rgb(234 179 8)" : "white",
                                marginRight: 8
                            }} />}

                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("Vietnamese")}
                    onPress={() => i18n.changeLanguage("vi")}
                    iconRight={
                        <FontAwesomeIcon
                            icon={faCheck}
                            size={18}
                            style={{
                                color: language == "vi" ? "rgb(234 179 8)" : "white",
                                marginRight: 8
                            }} />}
                />

            </ScrollView>
        </View>
    );
}


export default Language