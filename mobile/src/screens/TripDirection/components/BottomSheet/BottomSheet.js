import { Image, Pressable, Text, TouchableOpacity, View } from "react-native"
import className from "./className"
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome"
import { faChevronDown, faChevronUp, faCreditCard, faMoneyBillWave } from "@fortawesome/free-solid-svg-icons"
import { useEffect } from "react"
import { useState } from "react"
import { iconCarActive, iconCarInactive, iconMotorActive, iconMotorInactive } from "../../../../components/Icon"
import { useTranslation } from "react-i18next"


const BottomSheet = ({ methodPayment, distance, destination, source, vehicle, onChangeVehicle, price, onChangePayment }) => {

    const [show, setShow] = useState(false);

    const { t } = useTranslation()

    return (
        <View className={className.wrapper}>
            <TouchableOpacity
                className={className.header}
                activeOpacity={0.5}
                onPress={() => setShow(!show)}>
                <FontAwesomeIcon
                    icon={show ? faChevronDown : faChevronUp}
                    color="black" />
            </TouchableOpacity>
            <View className={className.vehicle}>
                <TouchableOpacity
                    activeOpacity={0.8}
                    className={className.buttonVehicle}
                    onPress={() => onChangeVehicle("motor")}>
                    <View className={vehicle == "motor" ? className.borderActive : className.border}>
                        <Image
                            source={vehicle == "motor" ? iconMotorActive : iconMotorInactive}
                            style={{ width: 60, height: 60, borderWidth: 4, borderColor: "white", borderRadius: 9999 }} />
                    </View>
                    <Text className={className.money}>{price?.motor?.toLocaleString()} VND</Text>
                </TouchableOpacity>
                <TouchableOpacity
                    activeOpacity={0.8}
                    className={className.buttonVehicle}
                    onPress={() => onChangeVehicle("car")}>
                    <View className={vehicle == "car" ? className.borderActive : className.border}>
                        <Image
                            source={vehicle == "car" ? iconCarActive : iconCarInactive}
                            style={{ width: 60, height: 60, borderWidth: 4, borderColor: "white", borderRadius: 9999 }} />
                    </View>
                    <Text className={className.money}>{price?.car?.toLocaleString()} VND</Text>
                </TouchableOpacity>
            </View>
            {show && (
                <View>
                    <View className="flex-row px-3 my-2">
                        <Text className="font-semibold flex-[2]">{t("Source")}:</Text>
                        <Text className="font-semibold flex-[3] text-right">{source.description}</Text>
                    </View>
                    <View className="flex-row px-3 my-2">
                        <Text className="font-semibold flex-[2]">{t("Destination")}:</Text>
                        <Text className="font-semibold flex-[3] text-right">{destination.description}</Text>
                    </View>
                    <View className="flex-row px-3 my-2">
                        <Text className="font-semibold flex-[2]">{t("Distance")}:</Text>
                        <Text className="font-semibold flex-[3] text-right">{distance} KM</Text>
                    </View>
                </View>
            )}
            <View className={className.payment}>
                <Pressable
                    className={methodPayment ? className.buttonPaymentActive : className.buttonPayment}
                    onPress={() => onChangePayment(true)}>
                    <FontAwesomeIcon
                        icon={faMoneyBillWave}
                        size={24}
                        color="green" />
                    <View className={className.column}>
                        <Text className={className.money2}>
                            {vehicle == "motor" ? price?.motor / 1000 : price?.car / 1000}K
                        </Text>
                        <Text className={className.description}>{t("Cash")}</Text>
                    </View>
                </Pressable>
                <Pressable
                    className={methodPayment ? className.buttonPayment : className.buttonPaymentActive}
                    onPress={() => onChangePayment(false)}>
                    <FontAwesomeIcon
                        icon={faCreditCard}
                        size={24}
                        color="blue" />
                    <View className={className.column}>
                        <Text className={className.money2}>
                            {vehicle == "motor" ? price?.motor / 1000 : price?.car / 1000}K
                        </Text>
                        <Text className={className.description}>{t("Card")}</Text>
                    </View>
                </Pressable>
            </View>
        </View >
    )
}

export default BottomSheet