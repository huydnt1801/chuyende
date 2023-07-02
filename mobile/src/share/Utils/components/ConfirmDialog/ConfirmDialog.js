import { createRef, useEffect, useState } from "react"
import { useTranslation } from "react-i18next";
import { Modal, View } from "react-native";
import className from "./className";
import { Text } from "react-native";
import { TouchableOpacity } from "react-native";

const confirmDialogRef = createRef();


const ConfirmDialog = () => {

    const { t } = useTranslation();

    const defaultConfig = {
        message: "Bạn nhận được 1 cái nịt",
        onConfirm: () => 1,
        onCancel: () => 1,
        buttonTextLeft: t("Cancel"),
        buttonTextRight: t("Agree")
    }

    const [show, setShow] = useState(false);
    const [config, setConfig] = useState({ ...defaultConfig });

    useEffect(() => {
        confirmDialogRef.current = {
            show: (config) => {
                setConfig({
                    ...defaultConfig,
                    ...config
                });
                setShow(true);
            },
            hide: () => {
                setShow(false);
                setConfig({ ...defaultConfig });
            }
        }
    });

    return (
        <Modal
            visible={show}
            transparent={true}>
            <View className={className.overlay}>
                <View className={className.content}>
                    <Text className={className.title}>{t("Notify")}</Text>
                    <Text
                        className={className.message}
                        style={{ textAlignVertical: "center" }}>
                        {config.message}
                    </Text>
                    <View className={className.group}>
                        <TouchableOpacity
                            className={className.buttonLeft}
                            activeOpacity={0.8}
                            onPress={() => {
                                config.onCancel();
                                setShow(false);
                            }}>
                            <Text className={className.buttonText}>{config.buttonTextLeft.toUpperCase()}</Text>
                        </TouchableOpacity>
                        <TouchableOpacity
                            className={className.buttonRight}
                            activeOpacity={0.8}
                            onPress={() => {
                                config.onConfirm();
                                setShow(false);
                            }}>
                            <Text className={className.buttonText}>{config.buttonTextRight.toUpperCase()}</Text>
                        </TouchableOpacity>
                    </View>
                </View>
            </View>
        </Modal>
    )
}

export default ConfirmDialog
export {
    confirmDialogRef
}