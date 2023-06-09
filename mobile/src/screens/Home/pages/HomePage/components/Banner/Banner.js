import { ImageBackground, Pressable, TextInput, View } from "react-native"
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";

import className from "./className"
import { faLocationDot } from "@fortawesome/free-solid-svg-icons";
import bebanner from "../../../../../../assets/images/bebanner.png";
import { useTranslation } from "react-i18next";

const Banner = ({ onPress }) => {

    const { t } = useTranslation()

    return (
        <ImageBackground
            source={bebanner}
            className={className.wrapper}>
            <Pressable
                className={className.border}>
                <FontAwesomeIcon
                    icon={faLocationDot}
                    size={18}
                    style={{ color: "rgb(234 179 8)", marginRight: 8 }} />
                <TextInput
                    placeholder={t('WhereDoWantToGo')} />
                <Pressable
                    className={className.overlay}
                    onPress={onPress}>
                </Pressable>
            </Pressable>
            <View className={className.devide}></View>
        </ImageBackground>
    );
}

export default Banner