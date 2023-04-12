import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";
import { } from "nativewind";
import { I18nextProvider } from "react-i18next";
import { Provider } from "react-redux"

import Home from "./src/screens/Home";
import Login from "./src/screens/Login";
import Mapp from "./src/screens/Mapp";
import store from "./src/store";

const Stack = createNativeStackNavigator();

const screens = [
    { name: "Home", component: Home },
    { name: "Login", component: Login },
    { name: "Mapp", component: Mapp },
]

const App = () => {
    return (
        <Provider store={store}>
            <I18nextProvider>
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