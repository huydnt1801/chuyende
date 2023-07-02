import { useNavigation } from "@react-navigation/native";
import { Text, View, StatusBar, Pressable, Button, ScrollView, Image, StyleSheet } from "react-native";

import className from "./className";
import Banner from "./components/Banner/Banner";
import { bebanner, bevoucher, blitzcrank } from "../../../../components/Image";
import { useSelector } from "react-redux";
import { locationTypes } from "../../../SelectLocation";

import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faCheck, faDriversLicense, faMotorcycle, faPerson } from "@fortawesome/free-solid-svg-icons";
import { useTranslation } from "react-i18next";
const HomePage = () => {

    const navigation = useNavigation();
    const { t } = useTranslation()
    const { account } = useSelector(state => state.account);
    return (
        <View className={className.container}>
            <View className={className.header}>
                <Text className={className.welcome}>
                    {`${t("Wellcome")} ${account?.fullName}!`}
                </Text>
            </View>
            <ScrollView>
                <Banner
                    onPress={() => navigation.navigate("SelectLocation", { type: locationTypes.SELECT_DESTINATION })} />
                <View className={className.action}>
                    <Pressable className={className.button}>
                        <FontAwesomeIcon
                            icon={faMotorcycle}
                            size={40}
                            style={{
                                color: "rgb(234 179 8)",
                                marginRight: 8
                            }} />
                        {/* <Image
                            className={className.imageButton}
                            source={blitzcrank} /> */}
                        <Text className={className.textButton}>{t("FindCar")}</Text>
                    </Pressable>
                    <Pressable className={className.button}>
                        <FontAwesomeIcon
                            icon={faPerson}
                            size={40}
                            style={{
                                color: "rgb(234 179 8)",
                                marginRight: 8
                            }} />
                        {/* <Image
                            className={className.imageButton}
                            source={blitzcrank} /> */}
                        <Text className={className.textButton}>{t("HireDriver")}</Text>
                    </Pressable>
                </View>
                <View>
                    <Image
                        style={{
                            resizeMode: 'repeat',
                            // height: 100,
                            // width: 200,
                            // display: 'flex'
                        }}
                        source={bevoucher} />
                </View>

            </ScrollView>
            {/* <Image
                style={{
                    resizeMode: 'center',
                    // height: 100,
                    // width: 200,
                    // display: 'flex'
                }}
                source={bevoucher} /> */}
        </View>
    );
}

export default HomePage