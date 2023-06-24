import { ScrollView, StatusBar, Text, View } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faAngleRight, faCar, faEnvelope, faGear, faGem, faHeadphones, faLink, faMoneyBill, faShareNodes, faShieldAlt } from "@fortawesome/free-solid-svg-icons";

import Avatar from "./components/Avatar";
import ButtonRow from "./components/ButtonRow";
import className from "./className";

const Account = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();

    return (
        <View className={className.container}>
            <ScrollView>
                <Avatar
                    name={"Truong Quang Phu"}
                    phone={"+84123132123"}
                    rate={4.8}
                    onPress={() => 1} />
                <ButtonRow
                    className={``}
                    title={t("Liên Kết Tài Khoản")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faLink}
                            size={22}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                // iconRight={
                //     <FontAwesomeIcon
                //         icon={faAngleRight}
                //         size={16}
                //         style={{ color: "rgb(107 114 128)" }} />
                // }
                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("Cài đặt chuyến đi")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faCar}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                    iconRight={
                        <FontAwesomeIcon
                            icon={faAngleRight}
                            size={14}
                            style={{ color: "rgb(107 114 128)" }} />
                    }
                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("Bảo hiểm chuyến đi")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faShieldAlt}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("Khuyến mãi")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faShieldAlt}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-[1px]`}
                    title={t("Gói tiết kiệm")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faGem}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-[1px]`}
                    title={t("Giới thiệu & Nhận ưu đãi")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faShareNodes}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-[1px]`}
                    title={t("Thanh toán")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faMoneyBill}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("Hộp thư")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faEnvelope}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-[1px]`}
                    title={t("Hỗ trợ")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faHeadphones}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                    iconRight={
                        <FontAwesomeIcon
                            icon={faAngleRight}
                            size={16}
                            style={{ color: "rgb(107 114 128)" }} />
                    }
                />
                <ButtonRow
                    classNames={`mt-[1px]`}
                    title={t("Cài đặt")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faGear}
                            size={18}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                    iconRight={
                        <FontAwesomeIcon
                            icon={faAngleRight}
                            size={16}
                            style={{ color: "rgb(107 114 128)" }} />
                    }
                    onPress={() => navigation.navigate("Setting")}
                />
                <Text className={className.version}>{"Phiên bản: 1.1.0"}</Text>
                <View className={className.bottom}></View>
            </ScrollView>
        </View>
    );
}

export default Account