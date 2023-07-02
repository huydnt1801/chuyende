import { useNavigation } from "@react-navigation/native";
import { View } from "react-native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faClock, faClockRotateLeft, faGear, faGift, faHeart, faHome, faPersonFalling, faUser } from "@fortawesome/free-solid-svg-icons";

import HomePage from "./pages/HomePage";
import Account from "./pages/Account";
import History from "./pages/History";
import Reward from "./pages/Reward";
import { useDispatch, useSelector } from "react-redux";
import Utils from "../../share/Utils";
import { Text } from "react-native";
import Api from "../../api";
import Avatar from "./pages/Account/components/Avatar";
import { FlatList } from "react-native";
import { useState } from "react";
import className from "./className";
import { Image } from "react-native";
import { iconCarActive, iconMotorActive } from "../../components/Icon";
import { Pressable } from "react-native";
import { useTranslation } from "react-i18next";
import { RefreshControl } from "react-native";
import { TouchableOpacity } from "react-native";
import { useEffect } from "react";
import { Button } from "react-native";
import { thunkGetListTrip } from "../../slices/Trip";

const Tab = createBottomTabNavigator();

const pages = [
    {
        id: 0,
        name: "HomePage",
        component: HomePage,
        icon: faHome
    },
    {
        id: 1,
        name: "History",
        component: History,
        icon: faClockRotateLeft
    },
    {
        id: 2,
        name: "Reward",
        component: Reward,
        icon: faGift
    },
    {
        id: 3,
        name: "Account",
        component: Account,
        icon: faUser
    },
]

const TripItem = ({ vehicle, startLocation, endLocation, distance, price, time, onPressAccept }) => {

    const { t } = useTranslation();
    const today = new Date()
    const publishTime = new Date(time)
    const a = (today - publishTime)
    let b = ''
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
                        <FontAwesomeIcon
                            icon={faClock}
                            style={{ color: "rgb(156 163 175)", marginRight: 6 }} />
                        <Text>{b}</Text>
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
            <TouchableOpacity
                className={className.button}
                activeOpacity={0.7}
                onPress={onPressAccept}>
                <Text className={className.accept}>{t("AcceptTrip")}</Text>
            </TouchableOpacity>
        </Pressable>
    )
}

const Home = () => {

    const navigation = useNavigation();
    const { t } = useTranslation();
    const dispatch = useDispatch()

    const { account, isDriver } = useSelector(state => state.account);
    const { listTrip } = useSelector(state => state.trip);
    const trips = listTrip.filter(i => i.type == account?.license)

    const [isLoading, setIsLoading] = useState(false);

    const handleRefresh = async () => {
        setIsLoading(true);
        dispatch(thunkGetListTrip());
        setIsLoading(false);
    }

    console.log(account);

    useEffect(() => {
        if (isDriver) {
            dispatch(thunkGetListTrip());
        }
    }, [])

    if (isDriver) {
        return (
            <View className={className.container}>
                <TouchableOpacity
                    className={className.setting}
                    activeOpacity={0.5}
                    onPress={() => navigation.navigate("DriverSetting")}>
                    <FontAwesomeIcon
                        icon={faGear}
                        size={20} />
                </TouchableOpacity>
                <Avatar
                    name={account?.fullName}
                    phone={`+84 ${account?.phoneNumber?.slice(1)}`}
                    rate={4.8}
                    classNames={"pb-4"}
                    type={account?.license}
                    onPress={() => 1} />
                {/* <Button
                    title="hello"
                    onPress={() => navigation.navigate("InTrip", { tripId: 5 })} /> */}
                <Text className="ml-4 font-semibold">{t("ListTrip")}:</Text>
                <FlatList
                    data={trips}
                    contentContainerStyle={{
                        paddingHorizontal: 8,
                        paddingTop: 4,
                        paddingBottom: 20
                    }}
                    ListEmptyComponent={
                        <View className="items-center justify-center py-3">
                            <Text className="font-semibold">{t("NowThereAreNoTrip")}</Text>
                        </View>
                    }
                    refreshControl={
                        <RefreshControl
                            tintColor={"#C4C7C8"}
                            refreshing={isLoading}
                            onRefresh={handleRefresh} />
                    }
                    renderItem={({ item }) => (
                        <TripItem
                            vehicle={item.type}
                            startLocation={item.startLocation}
                            endLocation={item.endLocation}
                            distance={item.distance}
                            price={item.price}
                            time={item.createdAt}
                            onPressAccept={() => {
                                Utils.showConfirmDialog({
                                    message: t("AreYouSureAcceptThisTrip"),
                                    onConfirm: async () => {
                                        Utils.hideConfirmDialog();
                                        Utils.showLoading();
                                        const result = await Api.trip.updateStatus(item.id, "accept");
                                        await Utils.wait(500);
                                        Utils.hideLoading();
                                        if (result.result == Api.ResultCode.SUCCESS) {
                                            navigation.navigate("InTrip", { tripId: item.id })
                                        }
                                    }
                                })
                            }}
                        />
                    )} />
            </View>
        );
    }

    return (
        <Tab.Navigator
            screenOptions={{
                tabBarShowLabel: false,
                tabBarStyle: { height: 60 },
                tabBarHideOnKeyboard: true
            }}>
            {pages.map(item => (
                <Tab.Screen
                    key={item.id}
                    name={item.name}
                    component={item.component}
                    options={{
                        headerShown: false,
                        tabBarIcon: ({ focused }) => (
                            <View>
                                <FontAwesomeIcon
                                    className="bg-gray-300"
                                    icon={item.icon}
                                    size={24}
                                    style={{
                                        color: focused ? "rgb(234 179 8)" : " rgb(209 213 219)"
                                    }}
                                />
                            </View>
                        )
                    }}
                />
            ))}
        </Tab.Navigator>
    );
}

export default Home