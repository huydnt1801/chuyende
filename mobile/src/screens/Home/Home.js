import { useState, useEffect } from "react";
import { useTranslation } from "react-i18next";
import { useNavigation } from "@react-navigation/native";
import { View, Image, Pressable, TouchableOpacity, RefreshControl, FlatList, Text } from "react-native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faClockRotateLeft, faGear, faGift, faHeart, faHome, faPersonFalling, faUser } from "@fortawesome/free-solid-svg-icons";
import { useDispatch, useSelector } from "react-redux";

import HomePage from "./pages/HomePage";
import Account from "./pages/Account";
import History from "./pages/History";
import Reward from "./pages/Reward";
import Utils from "../../share/Utils";
import Api from "../../api";
import Avatar from "./pages/Account/components/Avatar";
import className from "./className";
import { thunkGetListTrip } from "../../slices/Trip";
import TripItem from "./TripItem";

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


const Home = () => {

    const navigation = useNavigation();
    const { t } = useTranslation();
    const dispatch = useDispatch();

    const { account, isDriver } = useSelector(state => state.account);
    const { listTrip } = useSelector(state => state.trip);
    const trips = listTrip.filter(i => i.type == account?.license)

    const [isLoading, setIsLoading] = useState(false);

    const handleRefresh = async () => {
        setIsLoading(true);
        dispatch(thunkGetListTrip());
        setIsLoading(false);
    }

    useEffect(() => {
        if (isDriver) {
            dispatch(thunkGetListTrip());
        }
    }, []);

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