import { Image, Pressable, Text, View } from "react-native";

import className from "./className";
import { blitzcrank } from "../../../../../../components/Image";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faStar } from "@fortawesome/free-solid-svg-icons";

const Avatar = ({ name, phone, rate, onPress, classNames }) => {

    return (
        <Pressable
            className={className.wrapper + " " + classNames}
            onPress={onPress}>
            <View className={className.imageBorder}>
                <Image
                    className={className.avatar}
                    source={blitzcrank} />
            </View>
            <Text className={className.name}>{name}</Text>
            <View className={className.infor}>
                <View className={className.rate}>
                    <FontAwesomeIcon
                        icon={faStar}
                        style={{ color: "rgb(234 179 8)" }} />
                    <Text className={className.ratePoint}>{rate}</Text>
                </View>
                <Text className={className.devide}>{"・"}</Text>
                <Text className={className.phone}>{phone}</Text>
            </View>
        </Pressable>
    )
}

export default Avatar