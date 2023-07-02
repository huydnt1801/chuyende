import { View } from "react-native";
import className from "./className";
import { useRef } from "react";
import { useEffect } from "react";
import { TextInput } from "react-native";
import { Pressable } from "react-native";
import { Text } from "react-native";
import { useState } from "react";
import { forwardRef } from "react";

/**
 * @typedef prop
 * @property {String | undefined} value
 * @property {String | undefined} ref
 * @property {((Number) => void) | undefined} onChangText
 * @param {prop} param
 * @returns 
 */
const InputPasswordOTP = ({ value, onChangText, classNames }, ref) => {

    const _value = value ?? "";
    const renderList = ["", "", "", "", "", ""];
    for (let i = 0; i < _value.length; i++) {
        renderList[i] = _value[i];
    }

    const _ref = ref ?? useRef();
    const [focus, setFocus] = useState(false);

    useEffect(() => {
        if (focus) {
            _ref.current?.blur();
            _ref.current?.focus();
        }
    }, []);

    return (
        <View className={className.inputWrapper + " " + classNames}>
            <TextInput
                ref={_ref}
                style={{ position: 'absolute', width: "20%", zIndex: -1, opacity: 0, height: 40 }}
                cursorColor={"rgba(0,0,0,0)"}
                value={value}
                onFocus={() => setFocus(true)}
                onBlur={() => setFocus(false)}
                maxLength={6}
                onChangeText={onChangText}
                keyboardType={"number-pad"} />
            {renderList.map((item, index) => (
                <Pressable
                    key={index}
                    className={className.block + (focus && ((index == value?.length || (index == value?.length - 1 && index == 5)) && " border-blue-400"))}
                    onPress={() => {
                        _ref.current.blur()
                        _ref.current.focus();
                    }}>
                    <Text
                        className={className.inputItem}>
                        {item}
                    </Text>
                </Pressable>
            ))}
        </View>
    );
}

export default forwardRef(InputPasswordOTP)