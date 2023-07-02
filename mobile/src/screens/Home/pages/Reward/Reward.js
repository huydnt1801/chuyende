import { ScrollView, Text, View, Pressable, Image, useWindowDimensions } from "react-native"
import Header from "../../../../components/Header";
import { useTranslation } from "react-i18next";
import { CommonActions, StackActions, useNavigation } from "@react-navigation/native";
import className from "./className";
import ButtonRow from "./ButtonRow/ButtonRow";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faCar } from "@fortawesome/free-solid-svg-icons";
import { bevoucher1 } from "../../../../components/Image";
import { bevoucher2 } from "../../../../components/Image";
const Reward = () => {
    const { t } = useTranslation();
    const navigation = useNavigation();
    const size = useWindowDimensions();
    const widthv = size.width;
    const heightv = size.height;
    return (
        <View className={className.container}>
            <View className={className.wrapper}>
                <View className={className.center}>
                    <Text className={className.title}>{t("Reward")}</Text>
                </View>
            </View>
            <ScrollView>
                <Pressable className={className.button}>

                    <Image
                        className={className.imageButton}
                        source={bevoucher1}
                        style={{
                            width:widthv,
                            resizeMode: 'contain',
                            height:200
                        }}
                    />

                </Pressable>
                <Pressable className={className.button}>

                    <Image
                        className={className.imageButton}
                        source={bevoucher2}
                        style={{
                            width: widthv,
                            resizeMode: 'contain',
                            height:200
                        }}
                    />
                    <Pressable className={className.button}>

                        <Image
                            className={className.imageButton}
                            source={bevoucher1}
                            style={{
                                width: widthv,
                                resizeMode: 'contain',
                                height: 200
                            }}
                        />

                    </Pressable>
                    
                </Pressable>
            </ScrollView>
        </View>

    )
}

export default Reward