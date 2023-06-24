import { Button } from "react-native";
import { Text, View } from "react-native";
import { useDispatch } from "react-redux";
import { setAccount } from "../../slices/Account";
import AsyncStorage from "@react-native-async-storage/async-storage";
import { StackActions, useNavigation } from "@react-navigation/native";

const Setting = () => {

    const dispatch = useDispatch();
    const navigation = useNavigation();

    return (
        <View style={{ flex: 1 }}>
            <Text>Setting</Text>
            <Button
                title="Log Out"
                onPress={async () => {
                    dispatch(setAccount(null));
                    try {
                        await AsyncStorage.removeItem("account");
                    } catch (error) {
                    }
                    navigation.dispatch(StackActions.replace("Splash"));
                }} />
        </View>
    );
}

export default Setting