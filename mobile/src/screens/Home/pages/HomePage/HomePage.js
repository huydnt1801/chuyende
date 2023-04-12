import { Text, View, StatusBar, Pressable, Button } from "react-native";
import { useNavigation } from "@react-navigation/native"

import className from "./className";

const HomePage = () => {

    const name = "Trương Quang Phú";
    const navigation = useNavigation();

    return (
        <View className={className.container}>
            <StatusBar />
            <View className={className.header}>
                <Text className={className.welcome}>
                    {`Chào ${name}!`}
                </Text>
                <Pressable>

                </Pressable>
            </View>
            <Button
                title="Hello"
                onPress={() => navigation.navigate("Mapp")} />
        </View>
    );
}

export default HomePage