import { Image, Pressable, Text, TouchableOpacity, View } from "react-native"
import className from "./className"
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome"
import { faChevronDown, faChevronUp, faMoneyBillWave } from "@fortawesome/free-solid-svg-icons"
import { useEffect } from "react"
import { useState } from "react"
import { iconCarActive, iconCarInactive, iconMotorActive, iconMotorInactive } from "../../../../components/Icon"


const BottomSheet = ({ methodPayment, vehicle, onChangeVehicle }) => {

    const [show, setShow] = useState(false);

    return (
        <View className={className.wrapper}>
            <Pressable className={className.header}>
                <FontAwesomeIcon
                    icon={show ? faChevronDown : faChevronUp}
                    color="black" />
            </Pressable>
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
                    <Text className={className.money}>{'30K'}</Text>
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
                    <Text className={className.money}>{'50K'}</Text>
                </TouchableOpacity>
            </View>
            <View className={className.payment}>
                <Pressable
                    className={className.buttonPayment}>
                    <FontAwesomeIcon
                        icon={faMoneyBillWave}
                        size={24}
                        color="green" />
                    <View className={className.column}>
                        <Text className={className.money2}>
                            {'30K'}
                        </Text>
                        <Text className={className.description}>{"Tiền mặt"}</Text>
                    </View>
                </Pressable>
                <Pressable className={className.buttonPayment}>

                </Pressable>
            </View>
        </View >
    )
}

export default BottomSheet