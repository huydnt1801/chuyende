import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { } from "nativewind";
import { I18nextProvider } from "react-i18next";
import { Provider } from "react-redux";

import Home from "./src/screens/Home";
import Login from "./src/screens/Login";
import PasswordOTP from "./src/screens/PasswordOTP";
import Payment from "./src/screens/Payment";
import Setting from "./src/screens/Setting";
import TripSetting from "./src/screens/TripSetting";
import Mapp from "./src/screens/Mapp";
import SelectLocation from "./src/screens/SelectLocation";
import store from "./src/store";
import i18next from "./src/share/Language";

import TestScreen from "./src/screens/TestScreen";

const Stack = createNativeStackNavigator();

const screens = [
    { name: "TestScreen", component: TestScreen },
    { name: "Home", component: Home },
    { name: "Login", component: Login },
    { name: "Payment", component: Payment },
    { name: "Setting", component: Setting },
    { name: "TripSetting", component: TripSetting },
    { name: "PasswordOTP", component: PasswordOTP },
    { name: "SelectLocation", component: SelectLocation },
    { name: "Mapp", component: Mapp },
]

const App = () => {
    return (
        <Provider store={store}>
            <I18nextProvider i18n={i18next}>
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
            </I18nextProvider>
        </Provider>
    );
}

export default App