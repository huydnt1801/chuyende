import { Text, View, Button, StyleSheet, TouchableOpacity, TextInput } from "react-native";
import MapView, { Marker } from "react-native-maps"
import className from "./className"
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faHeart, faL, faPersonFalling } from "@fortawesome/free-solid-svg-icons";
import MapViewDirections from "react-native-maps-directions";
import { useSelector } from "react-redux";
import { useState } from "react";
import { useRef } from "react";
import BottomSheet from "./components/BottomSheet";
import Api from "../../api";
import Utils from "../../share/Utils";
import ModalFinding from "./components/ModalFinding";
import { memo } from "react";

const MarkerType = {
    START: 1,
    DESTINATION: 2
}

const CustomMarker = ({ type = MarkerType.START, place }) => {
    console.log(place);
    return (
        <View className={className.wrapper}>
            <Text>{place}</Text>
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

const Direction = memo(({ source, destination }) => {

    console.log("re-render");

    return (
        <MapViewDirections
            origin={source}
            destination={destination}
            apikey={"AIzaSyBt5Cp2LUkwqb8wq-wgDDjIN1KZTeHebY4"}
            strokeWidth={4}
            strokeColor="#111111"
        />
    )
});


const TripDirection = () => {

    const { t } = useTranslation();
    const { source, destination } = useSelector(state => state.trip)

    const defaultPosition = {
        latitude: (source.latitude + destination.latitude) / 2,
        longitude: (source.longitude + destination.longitude) / 2
    };

    const [vehicle, setVehicle] = useState("motor");
    const [showModalFinding, setShowModalFinding] = useState(false);

    console.log(Utils.data["cookie"]);

    const handleClickOrder = async () => {
        Utils.showLoading();
        const result = await Api.trip.createTrip({
            startX: source.latitude,
            startY: source.latitude,
            startLocation: source.description,
            endX: destination.latitude,
            endY: destination.latitude,
            endLocation: destination.description,
            distance: 10,
            type: vehicle
        });
        await Utils.wait(500);
        Utils.hideLoading();
        console.log(result);
        await Utils.wait(1000);
        setShowModalFinding(true);
    }

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
                    la>
                    <CustomMarker
                        type={MarkerType.DESTINATION}
                        place={destination.place}
                    />
                </Marker>
                <MapViewDirections
                    origin={source}
                    destination={destination}
                    apikey={"AIzaSyBt5Cp2LUkwqb8wq-wgDDjIN1KZTeHebY4"}
                    strokeWidth={4}
                    strokeColor="#111111"
                />
            </MapView>
            <BottomSheet
                vehicle={vehicle}
                onChangeVehicle={vehicle => setVehicle(vehicle)} />
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
                onPressCancel={() => setShowModalFinding(false)} />
        </View>
    )
}

export default TripDirection