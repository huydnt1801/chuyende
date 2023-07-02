import { useNavigation } from "@react-navigation/native";
import { View } from "react-native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faClock, faClockRotateLeft, faGear, faGift, faHeart, faHome, faPersonFalling, faUser } from "@fortawesome/free-solid-svg-icons";

import HomePage from "./pages/HomePage";
import Account from "./pages/Account";
import History from "./pages/History";
import Reward from "./pages/Reward";
import { useSelector } from "react-redux";
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
                        <Text>{"3 phút trước"}</Text>
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

    const { account, isDriver } = useSelector(state => state.account);

    const [trips, setTrips] = useState([]);
    const [isLoading, setIsLoading] = useState(false);

    const getTrips = async () => {
        const result = await Api.trip.getList();
        if (result.result = Api.ResultCode.SUCCESS) {
            setTrips(result.data.data);
        }
        else {

        }
    }

    const handleRefresh = async () => {
        setIsLoading(true);
        await getTrips();
        setIsLoading(false);
    }

    useEffect(() => {
        if (isDriver) {
            getTrips();
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
                    onPress={() => 1} />
                <Text className="ml-4 font-semibold">{t("ListTrip")}:</Text>
                <FlatList
                    data={trips}
                    contentContainerStyle={{
                        paddingHorizontal: 8,
                        paddingTop: 4,
                        paddingBottom: 20
                    }}
                    ListEmptyComponent={
                        <View>
                            <Text>Hien tai khong co chuyen di nao</Text>
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
                            time={item.c}
                            onPressAccept={() => {
                                Utils.showConfirmDialog({
                                    message: t("AreYouSureAcceptThisTrip"),
                                    onConfirm: async () => {

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