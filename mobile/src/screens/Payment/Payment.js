import { Text, View } from "react-native";
import Header from "../../components/Header";
import { useTranslation } from "react-i18next";

const Payment = () => {

    const { t } = useTranslation();


    return (
        <View style={{ flex: 1 }}>
            <Header
                title={t("Payment")} />
        </View>
    );
}

export default Payment