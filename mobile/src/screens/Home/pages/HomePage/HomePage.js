import { Text, View, Pressable, Button, ScrollView, Image } from "react-native";
import { useNavigation } from "@react-navigation/native"

import className from "./className";
import Banner from "./components/Banner/Banner";
import { blitzcrank } from "../../../../components/Image";
import { useSelector } from "react-redux";

const HomePage = () => {

    const name = "Trương Quang Phú";
    const navigation = useNavigation();
    const { account } = useSelector(state => state.account)

    return (
        <View className={className.container}>
            <View className={className.header}>
                <Text className={className.welcome}>
                    {`Chào ${account?.full_name}!`}
                </Text>
                <Pressable>

                </Pressable>
            </View>
            <ScrollView>
                <Banner
                    onPress={() => navigation.navigate("SelectLocation")} />
                <View className={className.action}>
                    <Pressable className={className.button}>
                        <Image
                            className={className.imageButton}
                            source={blitzcrank} />
                        <Text className={className.textButton}>{"Tìm xe"}</Text>
                    </Pressable>
                    <Pressable className={className.button}>
                        <Image
                            className={className.imageButton}
                            source={blitzcrank} />
                        <Text className={className.textButton}>{"Thuê tài xế"}</Text>
                    </Pressable>
                </View>
            </ScrollView>
        </View>
    );
}

export default HomePage