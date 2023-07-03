import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { useTranslation } from "react-i18next";
import { Image, TouchableOpacity } from "react-native";
import { Pressable, View } from "react-native";
import { iconCarActive, iconMotorActive } from "../../../components/Icon";
import { Text } from "react-native";
import { faClock, faHeart, faPersonFalling } from "@fortawesome/free-solid-svg-icons";

const className = {
    container: `flex-1 bg-gray-50 flex-col`,
    wrapper: `flex-col bg-white rounded mb-3 px-3 active:bg-gray-200`,
    top: `flex-row items-center py-2 border-b border-gray-100`,
    image: `w-10 h-10 rounded-full bg-yellow-200`,
    right: `flex-col ml-3`,
    row: `flex-row items-center py-3`,
    border: `bg-black p-[6px] rounded-full`,
    border2: `bg-yellow-400 p-2 rounded-full`,
    place: `ml-3 flex-1`,
    time: `flex-row items-center mb-[2px]`,
    information: `flex-row items-center`,
    divide: `mx-2 w-[1px] h-[14px] rounded-full bg-black`,
    bold: `font-semibold`,
    button: `mt-2 w-full items-center justify-center py-3 bg-yellow-400 rounded mb-2`,
    accept: `font-semibold text-white`,
    setting: `absolute top-5 right-5 z-[1] p-2`
}

const TripItem = ({ vehicle, startLocation, endLocation, distance, price, time, onPressAccept, history = false }) => {

    const { t } = useTranslation();
    const today = new Date();
    const publishTime = new Date(time);
    const a = (today - publishTime);
    let b = '';
    if (a < 60000) {
        b = `${Math.floor(a / 1000)} ${t("SecondAgo")}`;
    } else if (a < 3600000) {
        b = `${Math.floor(a / 60000)} ${t("MinuteAgo")}`;
    } else if (a < 86400000) {
        b = `${Math.floor(a / 3600000)} ${t("HourAgo")}`;
    } else if (a < 2592000000) {
        b = `${Math.floor(a / 86400000)} ${t("DayAgo")}`;
    } else if (a < 31104000000) {
        b = `${Math.floor(a / 2592000000)} ${t("MonthAgo")}`;
    } else if (a > 31104000000) {
        b = `${Math.floor(a / 31104000000)} ${t("YearAgo")}`;
    }

    return (
        <Pressable className={className.wrapper}>
            <View className={className.top}>
                <Image
                    className={className.image}
                    source={vehicle == "motor" ? iconMotorActive : iconCarActive} />
                <View className={className.right}>
                    <View className={className.time}>
                        {history ?
                            (<>
                                <Text className={className.bold}>
                                    {`${publishTime.getDate() > 9 ? publishTime.getDate() : "0" + publishTime.getDate()}/${publishTime.getMonth() > 8 ? publishTime.getMonth() + 1 : ("0" + (publishTime.getMonth() + 1))}/${publishTime.getFullYear()}`}
                                </Text>
                                <View className={className.divide}></View>
                                <Text className={className.bold}>
                                    {`${publishTime.getHours() > 9 ? publishTime.getHours() : "0" + publishTime.getHours()}:${publishTime.getMinutes() > 9 ? publishTime.getMinutes() : "0" + publishTime.getMinutes()}`}
                                </Text>
                            </>
                            )
                            :
                            (
                                <>
                                    <FontAwesomeIcon
                                        icon={faClock}
                                        style={{ color: "rgb(156 163 175)", marginRight: 6 }} />
                                    <Text>{b}</Text>
                                </>

                            )}

                    </View>
                    <View className={className.information}>
                        <Text className={className.bold}>{t("Distance")}: </Text>
                        <Text className={className.bold}>{distance} KM</Text>
                        <View className={className.divide}></View>
                        <Text className={className.bold}>{t("Price")}: </Text>
                        <Text className={className.bold}>{price?.toLocaleString()} VND</Text>
                    </View>
                </View>
            </View>
            <View className={className.row}>
                <View className={className.border}>
                    <FontAwesomeIcon
                        icon={faPersonFalling}
                        transform={{ rotate: 45 }}
                        size={18}
                        style={{ color: "white" }} />
                </View>
                <Text
                    className={className.place}
                    numberOfLines={1}
                    lineBreakMode="tail">
                    {startLocation}
                </Text>
            </View>
            <View className={className.row}>
                <View className={className.border2}>
                    <FontAwesomeIcon
                        icon={faHeart}
                        size={14}
                        style={{ color: "white" }} />
                </View>
                <Text
                    className={className.place}
                    numberOfLines={1}
                    lineBreakMode="tail">
                    {endLocation}
                </Text>
            </View>
            {!history && (
                <TouchableOpacity
                    className={className.button}
                    activeOpacity={0.7}
                    onPress={onPressAccept}>
                    <Text className={className.accept}>{t("AcceptTrip")}</Text>
                </TouchableOpacity>
            )}
        </Pressable>
    )
}

export default TripItem