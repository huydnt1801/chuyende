import { Button } from "react-native";
import { Text, View } from "react-native";
import { useDispatch } from "react-redux";
import { setAccount } from "../../slices/Account";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { StackActions, useNavigation } from "@react-navigation/native";
import {
    ScrollView,
} from "react-native";
import Header from "../../components/Header";
import { useTranslation } from "react-i18next";
import className from "./className";
import ButtonRow from "./ButtonRow";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faFingerprint, faFlagUsa, faKeyboard, faLanguage, faPowerOff, } from "@fortawesome/free-solid-svg-icons";
import Utils from "../../share/Utils";

const Language = () => {

    const dispatch = useDispatch();
    const navigation = useNavigation();
    const { t } = useTranslation();
    return (
        <View className={className.container}>
            <View >
                <Header
                    title={t("SelectLanguage")} />
            </View>
            <ScrollView>

                <ButtonRow
                    classNames={`mt-2`}
                    title={t("English")}
                    onPress={() => Utils.toast("Coming Soon!")}
                    // iconLeft={
                    //     <FontAwesomeIcon
                    //         icon={faFlagUsa}
                    //         size={18}
                    //         style={{
                    //             color: "rgb(107 114 128)",
                    //             marginRight: 8
                    //         }} />}

                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("Vietnamese")}
                    onPress={() => Utils.toast("Coming Soon!")}
                    // iconLeft={
                    //     <FontAwesomeIcon
                    //         icon={faLanguage}
                    //         size={18}
                    //         style={{
                    //             color: "rgb(107 114 128)",
                    //             marginRight: 8
                    //         }} />}
                />

            </ScrollView>
        </View>
    );
}


export default Language