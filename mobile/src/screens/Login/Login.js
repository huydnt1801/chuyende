import {
    Text,
    TouchableOpacity,
    View,
    Pressable,
    Image,
    TextInput
} from "react-native";
import { useTranslation } from "react-i18next";
import { useRef, useState } from "react";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faFacebook, faGoogle } from "@fortawesome/free-brands-svg-icons";
import { StackActions, useNavigation } from "@react-navigation/native";

import { flagVietnam } from "../../components/Image";
import className from "./className";
import Utils from "../../share/Utils";
import Api from "../../api";
import { types } from "../PasswordOTP"
import { useEffect } from "react";
import InputPasswordOTP from "../../components/UI/InputPasswordOTP/InputPasswordOTP";
import { useDispatch } from "react-redux";
import { setAccount, setIsDriver } from "../../slices/Account";
import AsyncStorage from "@react-native-async-storage/async-storage";
/**
 * A text component can be clickable
 * @typedef props
 * @property {string | undefined} text
 * @property {number | undefined} opacity
 * @property {() => void | undefined} onPress
 * @param {props} 
 * @returns {JSX.Element} 
 */
const PressableText = ({ text, onPress, opacity = 0.5 }) => {

    if (!text) {
        return <></>
    }

    const [focus, setFocus] = useState(false);
    const renderList = text.split(" ");

    return (
        <>
            {renderList.map((item, index) => {
                return (
                    <Pressable
                        key={index}
                        style={{ opacity: focus ? opacity : 1.0 }}
                        onTouchStart={() => setFocus(true)}
                        onTouchEnd={() => setFocus(false)}
                        onPress={onPress}>
                        <Text className={"text-blue-600 font-semibold -mb-1"}>
                            {item}{" "}
                        </Text>
                    </Pressable>
                )
            })}
        </>
    )
}

const Login = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();
    const dispatch = useDispatch();

    const ref = useRef();
    const passwordRef = useRef();

    const [phone, setPhone] = useState("");
    const [password, setPassword] = useState("");
    const [isDriver, setIsDriver_] = useState(false);

    const handleClick = async () => {
        const _phone = phone.length == 10 ? phone : "0" + phone;
        if (isDriver) {
            passwordRef.current?.blur()
            Utils.showLoading();
            const result = await Api.account.loginDriver(_phone, password);
            await Utils.wait(500);
            Utils.hideLoading();
            console.log(result);
            if (result.result == Api.ResultCode.SUCCESS) {
                dispatch(setAccount(result.data.data));
                dispatch(setIsDriver(true));
                Utils.global.isDriver = true;
                const cookie = String(result.headers["set-cookie"] ?? "");
                Utils.global.cookie = cookie;
                try {
                    await AsyncStorage.setItem("account", JSON.stringify(result.data.data));
                    await AsyncStorage.setItem("cookie", cookie);
                    await AsyncStorage.setItem("isDriver", String(1));
                } catch (error) { };
                navigation.dispatch(StackActions.replace("Home"))
            }
            else {
                if (result.data && result.data.code == 404) {
                    Utils.showMessageDialog({
                        message: t("CanNotFindAccount")
                    });
                }
                else {
                    if (result.data && result.data.code == 401) {
                        Utils.showMessageDialog({
                            message: t("WrongPassword")
                        });
                    }
                }
            }
        }
        else {
            Utils.showLoading();
            const result = await Api.account.checkPhone(_phone);
            await Utils.wait(500);
            Utils.hideLoading();
            if (!result.data.data) {
                navigation.navigate("PasswordOTP", { type: types.PASSWORD_AND_NAME, userPhone: _phone })
            }
            else {
                navigation.navigate("PasswordOTP", { type: types.ONLY_PASSWORD, userPhone: _phone })
            }
        }
    }

    const isButtonActive = (() => {
        if (isDriver) {
            if (password.length == 6) {
                if (phone.length == 9) {
                    return true;
                }
                if (phone.length == 10) {
                    if (phone[0] == "0") {
                        return true;
                    }
                }
                return false;
            }
            else {
                return false;
            }
        }
        if (phone.length == 9) {
            return true;
        }
        if (phone.length == 10) {
            if (phone[0] == "0") {
                return true;
            }
        }
        return false;
    })();

    useEffect(() => {
        ref.current?.blur()
        ref.current?.focus()
    }, [isDriver])

    return (
        <View className={className.container}>
            <View className={className.main}>
                <Text className={className.hello}>{isDriver ? t("Login") : t("GoWithApp")}</Text>
                <Text className={className.login}>{isDriver ? t("LoginWithDriverAccount") : t("LoginRegisterNow")}</Text>
                <View className={className.wrapper}>
                    <Image
                        className={className.flag}
                        resizeMode={"center"}
                        source={flagVietnam} />
                    <Text className={className.prePhone}>
                        {"+84"}
                    </Text>
                    <TextInput
                        ref={ref}
                        placeholder={t("YourPhone")}
                        className={className.phone + (phone.length > -1 && " text-lg font-bold tracking-widest")}
                        value={phone}
                        maxLength={10}
                        onSubmitEditing={() => passwordRef.current?.focus()}
                        keyboardType={"numeric"}
                        onChangeText={phone => {
                            setPhone(
                                phone.replace(",", "")
                                    .replace(" ", "")
                                    .replace(".", "")
                                    .replace("-", "")
                            )
                        }} />
                </View>
                {isDriver && (
                    <>
                        <Text className={className.enterPassword}>{t("EnterPassword")}:</Text>
                        <InputPasswordOTP
                            ref={passwordRef}
                            value={password}
                            onChangText={password => {
                                setPassword(
                                    password.replace(",", "")
                                        .replace(" ", "")
                                        .replace(".", "")
                                        .replace("-", "")
                                )
                            }} />
                    </>
                )}
                {!isDriver && (
                    <Text className={className.agree}>
                        {t("IAgreeWith")}{" "}
                        <PressableText
                            text={t("ConditionsAndRules")}
                            onPress={() => 123} />
                        {t("And").toLowerCase()}{" "}
                        <PressableText
                            text={t("TransportContract")}
                        />
                        {t("OfApp")}
                    </Text>
                )}
                <Pressable
                    className={isButtonActive ? className.activate : className.deactivate}
                    disabled={!isButtonActive}
                    onPress={handleClick}>
                    <Text className={className.buttonText}>{t("Continue")}</Text>
                </Pressable>
                <View className={className.driver}>
                    <PressableText
                        text={isDriver ? t("IAmUser") : t("IAmDriver")}
                        onPress={() => setIsDriver_(!isDriver)} />
                </View>
            </View>
            <View className={className.bottom}>
                <Text>{t("OrLoginBy")}</Text>
                <View className={className.iconGroup}>
                    <TouchableOpacity
                        className={className.icon}
                        activeOpacity={0.5}
                        onPress={() => Utils.toast("Coming Soon!")}>
                        <FontAwesomeIcon
                            icon={faFacebook}
                            size={36} />
                    </TouchableOpacity>
                    <TouchableOpacity
                        className={className.icon}
                        activeOpacity={0.5}
                        onPress={() => Utils.toast("Coming Soon!", 5000)}>
                        <FontAwesomeIcon
                            icon={faGoogle}
                            size={36} />
                    </TouchableOpacity>
                </View>
            </View>
        </View>
    );
}

export default Login