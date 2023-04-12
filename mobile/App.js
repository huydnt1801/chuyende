import { StyleSheet, Text, View } from "react-native";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { } from "nativewind"

import Home from "./src/screens/Home";
import Login from "./src/screens/Login";
import Mapp from "./src/screens/Mapp";

const Stack = createNativeStackNavigator();

const screens = [
    { name: "Home", component: Home },
    { name: "Login", component: Login },
    { name: "Mapp", component: Mapp },
]

const App = () => {
    return (
        <NavigationContainer>
            <Stack.Navigator>
                {screens.map(item => (
                    <Stack.Screen
                        key={item.name}
                        name={item.name}
                        component={item.component}
                        options={{ headerShown: false }} />
                ))}
            </Stack.Navigator>
        </NavigationContainer>
    );
}

export default App

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: "#fff",
        alignItems: "center",
        justifyContent: "center",
    },
});
