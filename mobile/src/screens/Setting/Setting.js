import { Button } from "react-native";
import { Text, View } from "react-native";
import { useDispatch, useSelector } from "react-redux";
import { setAccount } from "../../slices/Account";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { CommonActions, StackActions, useNavigation } from "@react-navigation/native";
import {
    ScrollView,
} from "react-native";
import Header from "../../components/Header";
import { useTranslation } from "react-i18next";
import className from "./className";
import ButtonRow from "./ButtonRow";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faFingerprint, faKeyboard, faLanguage, faPowerOff, } from "@fortawesome/free-solid-svg-icons";
import Utils from "../../share/Utils";
import Api from "../../api";

const Setting = () => {

    const dispatch = useDispatch();
    const navigation = useNavigation();
    const { t } = useTranslation();

    const { account } = useSelector(state => state.account)
    return (
        <View className={className.container}>
            <View >
                <Header
                    title={t("Setting")}
                    onPressBack={() => navigation.goBack()}
                />
            </View>
            <ScrollView>

                <ButtonRow
                    classNames={`mt-2`}
                    title={t("SelectLanguage")}
                    onPress={() => navigation.navigate("Language")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faLanguage}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}

                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("SetUpUpdatePassword")}
                    onPress={() => navigation.navigate("ChangePassword")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faKeyboard}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("BiometricAuthentication")}
                    onPress={() => Utils.toast("Coming Soon!")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faFingerprint}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("Logout")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faPowerOff}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                    onPress={async () => {
                        dispatch(setAccount(null));
                        if (account) {
                            if (account.phoneNumber) {
                                await Api.account.logout(account.phoneNumber);
                            }
                        }
                        try {
                            await AsyncStorage.removeItem("account");
                            await AsyncStorage.removeItem("cookie");
                            await AsyncStorage.removeItem("isDriver");
                        } catch (error) {
                        }
                        navigation.dispatch(CommonActions.reset({
                            index: 0,
                            routes: [{ name: "Splash" }]
                        }));
                    }}
                />

            </ScrollView>
        </View>
    );
}


export default Setting