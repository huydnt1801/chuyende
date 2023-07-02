import { View } from "react-native"
import { useDispatch, useSelector } from "react-redux";
import { useTranslation } from "react-i18next"
import AsyncStorage from "@react-native-async-storage/async-storage";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faPowerOff } from "@fortawesome/free-solid-svg-icons";
import { CommonActions, useNavigation } from "@react-navigation/native";

import Header from "../../components/Header"
import { ButtonRow } from "../../components/UI";
import { setAccount } from "../../slices/Account";
import Api from "../../api";

const DriverSetting = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();
    const dispatch = useDispatch();
    const { account } = useSelector(state => state.account)

    return (
        <View className={`flex-1`}>
            <Header
                title={t("Setting")}
                onPressBack={() => navigation.goBack()} />
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
                    } catch (error) { };
                    navigation.dispatch(CommonActions.reset({
                        index: 0,
                        routes: [{ name: "Splash" }]
                    }));
                }} />
        </View>
    )
}

export default DriverSetting