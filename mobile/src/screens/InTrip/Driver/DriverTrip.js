import { useEffect } from "react";
import { useState } from "react"
import { Text } from "react-native"
import { View } from "react-native"
import Api from "../../../api";
import { useTranslation } from "react-i18next";
import { TouchableOpacity } from "react-native";
import { ScrollView } from "react-native";
import Utils from "../../../share/Utils";
import { useNavigation } from "@react-navigation/native";
import { useDispatch } from "react-redux";
import { thunkGetListTrip } from "../../../slices/Trip";


const DriverTrip = ({ tripId }) => {

    const [trip, setTrip] = useState({});
    const { t } = useTranslation();
    const navigation = useNavigation();

    const getTripInformation = async () => {
        const result = await Api.trip.getTripInformation(tripId);
        if (result.result == Api.ResultCode.SUCCESS) {
            setTrip(result.data.data[0] ?? {})
        }
    }

    const dispatch = useDispatch();

    useEffect(() => {
        getTripInformation();
    }, []);

    return (
        <View className="flex-1 flex-col">
            <View className="flex-col items-center justify-center border-b border-gray-200 h-12">
                <Text className="font-semibold">{t("TripInformation")}</Text>
            </View>
            <ScrollView>
                <View className="px-3 py-3">
                    <Text className="font-semibold">{t("UserInformation")}:</Text>
                    <View className="flex-row mt-2 px-2">
                        <Text className="flex-[2]">{t("Name")}:</Text>
                        <Text className="font-semibold text-right">{trip?.user?.fullName}</Text>
                    </View>
                    <View className="flex-row px-2 mt-1 items-center">
                        <Text className="flex-[2]">{t("Phone")}:</Text>
                        <TouchableOpacity
                            className="bg-yellow-400 mr-2 py-1 px-2 rounded"
                            activeOpacity={0.7}
                            onPress={() => {
                                Utils.toast(t("Copied"))
                            }}>
                            <Text className="font-semibold text-white text-xs">{t("Copy")}</Text>
                        </TouchableOpacity>
                        <Text className="font-semibold text-right">{trip?.user?.phoneNumber}</Text>
                    </View>
                    <Text className="font-semibold mt-1">{t("TripInformation")}:</Text>
                    <View className="flex-row mt-2 px-2">
                        <Text className="flex-[2]">{t("Source")}:</Text>
                        <Text className="font-semibold flex-[4] text-right">{trip?.startLocation}</Text>
                    </View>
                    <View className="flex-row mt-2 px-2">
                        <Text className="flex-[2]">{t("Destination")}:</Text>
                        <Text className="font-semibold flex-[4] text-right">{trip?.endLocation}</Text>
                    </View>
                    <View className="flex-row mt-2 px-2">
                        <Text className="flex-[2]">{t("Vehicle")}:</Text>
                        <Text className="font-semibold text-right">{trip?.type == "motor" ? t("Bike") : t("Car")}</Text>
                    </View>
                    <View className="flex-row mt-2 px-2">
                        <Text className="flex-[2]">{t("Distance")}:</Text>
                        <Text className="font-semibold text-right">{trip?.distance} KM</Text>
                    </View>
                    <View className="flex-row mt-2 px-2">
                        <Text className="flex-[2]">{t("Price")}:</Text>
                        <Text className="font-semibold text-right">{trip?.price?.toLocaleString()} VND</Text>
                    </View>
                    <Text className="font-semibold my-2">{t("Detail")}:</Text>
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("AcceptTrip")}
                        isDone />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("ConfirmTrip")}
                        isDone />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("PickUser")}
                        isDone />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("TakeCar")} />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("CaptureCarAtStart")} />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("StartTrip")} />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("EndTrip")} />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("CaptureCarAtEnd")} />
                    <MessageRow
                        day={"20/06/2022"}
                        time={"11:20:58"}
                        message={t("BackCar")}
                        isEnd={true} />
                </View>
            </ScrollView>
            <TouchableOpacity
                className="items-center justify-end py-3 bg-yellow-400"
                activeOpacity={0.7}
                onPress={() => {
                    Utils.showConfirmDialog({
                        message: t("ConfirmCompleteTrip"),
                        onConfirm: async () => {
                            Utils.hideConfirmDialog();
                            Utils.showLoading();
                            const result = await Api.trip.updateStatus(tripId, "done");
                            await Utils.wait(500);
                            Utils.hideLoading();
                            if (result.result == Api.ResultCode.SUCCESS) {
                                dispatch(thunkGetListTrip())
                                navigation.goBack();
                            }
                            else {
                                dispatch(thunkGetListTrip())
                                navigation.goBack();
                            }
                        }
                    })
                }}>
                <Text className="font-semibold text-white">{t("CompleteTrip")}</Text>
            </TouchableOpacity>
        </View>
    );
}

const MessageRow = ({ day, time, message, isEnd = false, isDone }) => {

    return (
        <View className="flex-row px-2">
            <View className="flex-col">
                <Text>{day}</Text>
                <Text className="text-gray-500 text-right">{time}</Text>
            </View>
            <View className="flex-col mx-3 items-center">
                <View className="w-6 h-6 items-center justify-center rounded-full border-2 border-yellow-400">
                    {isDone && <View className="w-4 h-4 rounded-full bg-yellow-400"></View>}
                </View>
                {!isEnd && <>
                    <View className="w-[1px] h-2 bg-gray-500 my-1"></View>
                    <View className="w-[1px] h-2 bg-gray-500 my-1"></View>
                    <View className="w-[1px] h-2 bg-gray-500 my-1"></View>
                </>}
            </View>
            <Text className="font-semibold">{message}</Text>
        </View>
    )
}

export default DriverTrip