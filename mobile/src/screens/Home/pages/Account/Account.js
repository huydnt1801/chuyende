import { ScrollView, Text, View } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { useSelector } from "react-redux";
import { faAngleRight, faCar, faEnvelope, faGear, faGem, faHeadphones, faLink, faMoneyBill, faShareNodes, faShieldAlt } from "@fortawesome/free-solid-svg-icons";

import Avatar from "./components/Avatar";
import ButtonRow from "./components/ButtonRow";
import className from "./className";

const Account = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();
    const { account } = useSelector(state => state.account)

    return (
        <View className={className.container}>
            <ScrollView>
                <Avatar
                    name={account?.fullName}
                    phone={`+84 ${account?.phoneNumber?.slice(1)}`}
                    rate={4.8}
                    onPress={() => 1} />
                <ButtonRow
                    className={``}
                    title={t("LinkAccount")}
                    iconLeft={
                        <FontAwesomeIcon
                            icon={faLink}
                            size={22}
                            style={{
                                color: "rgb(107 114 128)",
                                marginRight: 8
                            }} />}
                />
                <ButtonRow
                    classNames={`mt-2`}
                    title={t("TripSetting")}
                    onPress={() => navigation.navigate("TripSetting")}
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
                    title={t("TripInsurance")}
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
                    title={t("Promotion")}
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
                    title={t("SavingPackage")}
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
                    title={t("ReferAndGetDeals")}
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
                    title={t("Payment")}
                    onPress={() => navigation.navigate("Payment")}
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
                    title={t("Mail")}
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
                    title={t("Support")}
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
                    title={t("Setting")}
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