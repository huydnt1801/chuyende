import { useEffect, useState } from "react";
import { View, Image } from "react-native";
import { useSpring, animated } from "@react-spring/native";
import { useDispatch } from "react-redux";
import { useNavigation, StackActions } from "@react-navigation/native";

import { docker } from "../../components/Image";
import { setAccount, setIsDriver } from "../../slices/Account";
import Utils from "../../share/Utils";
import AsyncStorage from "@react-native-async-storage/async-storage";

const Splash = () => {

    const navigation = useNavigation();
    const dispatch = useDispatch();

    const [style, spring] = useSpring(() => ({
        top: "100%",
        backgroundColor: "#d4ecf2"
    }));

    const rememberLogin = async () => {
        await Utils.wait(2000);
        Utils.showLoading();
        const account = await (async () => {
            try {
                const data = await AsyncStorage.getItem("account");
                return data;
            } catch (error) {
                return null;
            }
        })();
        const cookie = await (async () => {
            try {
                const data = await AsyncStorage.getItem("cookie");
                return data;
            } catch (error) {
                return null;
            }
        })();
        const _isDriver = await (async () => {
            try {
                const data = await AsyncStorage.getItem("isDriver");
                return data;
            } catch (error) {
                return null;
            }
        })();
        const isDriver = _isDriver ? (_isDriver == "1") : false;
        await Utils.wait(500);
        Utils.hideLoading();
        // console.log("cookie", cookie);
        // console.log("isDriver", isDriver);
        // console.log("account", account);
        Utils.global.isDriver = isDriver;
        if (cookie) {
            Utils.global.cookie = cookie;
        }
        if (account) {
            dispatch(setAccount(JSON.parse(account)));
            dispatch(setIsDriver(isDriver));
            navigation.dispatch(StackActions.replace("Home"));
        }
        else {
            navigation.dispatch(StackActions.replace("Login"));
        }
    }

    useEffect(() => {
        spring.start({
            top: "0%",
            config: {
                duration: 700
            }
        });
        rememberLogin();
    }, []);

    return (
        <View className="flex-1" style={{ backgroundColor: "#d4ecf2" }}>
            <animated.View className="items-center justify-center absolute w-full h-full" style={style}>
                <Image
                    source={docker}
                    resizeMode={"contain"}
                    style={{ width: "60%", marginBottom: 150 }}
                />
            </animated.View>
        </View>
    );
}

export default Splash