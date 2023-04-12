import { useNavigation } from "@react-navigation/native";
import { Pressable } from "react-native";
import { Text, View, Button, Image } from "react-native";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";

import HomePage from "./pages/HomePage";
import Account from "./pages/Account";
import iconHome from "../../assets/icon/home.png";
import iconAccount from "../../assets/icon/account.png"

const Tab = createBottomTabNavigator();

const Home = () => {

    const navigation = useNavigation();

    return (
        <Tab.Navigator
            screenOptions={{
                tabBarShowLabel: false,
                tabBarStyle: { height: 60 },
                tabBarHideOnKeyboard: true
            }}>
            <Tab.Screen
                name="HomePage"
                component={HomePage}
                options={{
                    headerShown: false,
                    tabBarIcon: ({ focused }) => (
                        <View>
                            <Image
                                style={[{
                                    width: 32,
                                    height: 32
                                }, { tintColor: focused ? "red" : "black" }]}
                                source={iconHome}
                            />
                        </View>
                    )
                }}
            />
            <Tab.Screen
                name="Account"
                component={Account}
                options={{
                    headerShown: false,
                    tabBarIcon: ({ focused }) => (
                        <View>
                            <Image
                                style={[{
                                    width: 32,
                                    height: 32
                                }, { tintColor: focused ? "red" : "black" }]}
                                source={iconAccount}
                            />
                        </View>
                    )
                }}
            />
        </Tab.Navigator>
    );
}

export default Home