import { Text, View, StyleSheet, TouchableOpacity } from "react-native";
import MapView, { Marker } from "react-native-maps"
import className from "./className"
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faHeart, faPersonFalling } from "@fortawesome/free-solid-svg-icons";
import MapViewDirections from "react-native-maps-directions";
import { useSelector } from "react-redux";
import { useState } from "react";
import BottomSheet from "./components/BottomSheet";
import Api from "../../api";
import Utils from "../../share/Utils";
import ModalFinding from "./components/ModalFinding";
import { useEffect } from "react";
import { useNavigation } from "@react-navigation/native";
import { useRef } from "react";

const MarkerType = {
    START: 1,
    DESTINATION: 2
}

const CustomMarker = ({ type = MarkerType.START, place }) => {
    return (
        <View className={className.wrapper}>
            <View
                className={className.top}>
                {type == MarkerType.START && (
                    <View className={className.circleStart}>
                        <FontAwesomeIcon
                            icon={faPersonFalling}
                            transform={{ rotate: 45 }}
                            size={16}
                            style={{ color: "white" }} />
                    </View>
                )}
                {type == MarkerType.DESTINATION && (
                    <View className={className.circleEnd}>
                        <FontAwesomeIcon
                            icon={faHeart}
                            size={12}
                            style={{ color: "white" }} />
                    </View>
                )}
                <View className={className.line}></View>
            </View>
            <View className={className.shadow}></View>
        </View>
    );
}

const TripDirection = () => {

    const { t } = useTranslation();
    const navigation = useNavigation();
    const { source, destination } = useSelector(state => state.trip)

    const defaultPosition = {
        latitude: (source.latitude + destination.latitude) / 2,
        longitude: (source.longitude + destination.longitude) / 2
    };

    const [vehicle, setVehicle] = useState("motor");
    const [showModalFinding, setShowModalFinding] = useState(false);
    const [price, setPrice] = useState({
        motor: 50,
        car: 50
    });

    const [distance, setDistance] = useState(0);
    const [tripId, setTripId] = useState(0);
    const ref = useRef(true);

    const [methodPayment, setMethodPayment] = useState(true);

    const handleClickOrder = async () => {
        Utils.showLoading();
        const result = await Api.trip.createTrip({
            startX: source.latitude,
            startY: source.latitude,
            startLocation: source.description,
            endX: destination.latitude,
            endY: destination.latitude,
            endLocation: destination.description,
            distance: distance,
            type: vehicle,
            price: price[vehicle]
        });
        await Utils.wait(1000);
        Utils.hideLoading();
        // await Utils.wait(1000);
        setShowModalFinding(true);
        setTripId(result.data.data.id)
        let status = await Api.trip.getTripInformation(result.data.data.id);
        while (status.data.data[0].status == "waiting") {
            await Utils.wait(3000);
            console.log("ref.current", ref.current);
            if (!ref.current) {
                break;
            }
            status = await Api.trip.getTripInformation(result.data.data.id);
        }
        setShowModalFinding(false);
        if (status.data.data[0].status == "accept") {
            navigation.navigate("InTrip", { tripId: result.data.data.id })
        }
    }

    useEffect(() => {
        const getPrice = async () => {

            const distance = await Api.google.checkDistance(source.description, destination.description);
            setDistance(distance);
            const result = await Api.trip.checkPrice(distance);
            if (result.result == Api.ResultCode.SUCCESS) {
                setPrice(result.data.data)
            }
        }
        getPrice();
    }, [])

    return (
        <View className={className.container}>
            <MapView
                style={StyleSheet.absoluteFillObject}
                showsPointsOfInterest={false}
                showsUserLocation={false}
                initialRegion={{
                    latitude: defaultPosition.latitude,
                    longitude: defaultPosition.longitude,
                    latitudeDelta: Math.abs(source.latitude - destination.latitude) + 0.005,
                    longitudeDelta: Math.abs(source.longitude - destination.longitude) + 0.005,
                }}>
                <Marker
                    coordinate={source}>
                    <CustomMarker
                        type={MarkerType.START}
                        place={source.place}
                    />
                </Marker>
                <Marker
                    coordinate={destination}
                >
                    <CustomMarker
                        type={MarkerType.DESTINATION}
                        place={destination.place}
                    />
                </Marker>
                <MapViewDirections
                    origin={source}
                    destination={destination}
                    // apikey={"AIzaSyBt5Cp2LUkwqb8wq-wgDDjIN1KZTeHebY4"}
                    strokeWidth={4}
                    strokeColor="#111111"
                />
            </MapView>
            <BottomSheet
                vehicle={vehicle}
                onChangeVehicle={vehicle => setVehicle(vehicle)}
                methodPayment={methodPayment}
                source={source}
                destination={destination}
                distance={distance}
                onChangePayment={e => setMethodPayment(e)}
                price={price} />
            <View className={className.bottom}>
                <TouchableOpacity
                    activeOpacity={0.7}
                    className={className.buttonOrder}
                    onPress={handleClickOrder}>
                    <Text className={className.buttonText}>{"Tìm tài xế"}</Text>
                </TouchableOpacity>
            </View>
            <ModalFinding
                show={showModalFinding}
                onPressCancel={() => {
                    Utils.showConfirmDialog({
                        message: t("AreYouSureToCancelThisTrip"),
                        onConfirm: async () => {
                            setShowModalFinding(false);
                            ref.current = true;
                            await Api.trip.updateStatus(tripId, "cancel");
                        },
                        onCancel: () => {
                            setShowModalFinding(true);
                        }
                    });
                }} />
        </View>
    )
}

export default TripDirection