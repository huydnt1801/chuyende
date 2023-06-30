import { Text, View, Button, StyleSheet, TouchableOpacity, TextInput } from "react-native";
import MapView, { Marker } from "react-native-maps"
import className from "./className"
import { useTranslation } from "react-i18next";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import { faHeart, faL, faPersonFalling } from "@fortawesome/free-solid-svg-icons";
import MapViewDirections from "react-native-maps-directions";
import { useSelector } from "react-redux";
import { useState } from "react";

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


const TripDirection = () => {

    const { source, destination } = useSelector(state => state.trip)

    const defaultPosition = {
        latitude: (source.latitude + destination.latitude) / 2,
        longitude: (source.longitude + destination.longitude) / 2
    };

    console.log(defaultPosition);
    // const defaultPosition = {
    //     latitude: 21.0285,
    //     longitude: 105.8542
    // };

    const start = {
        latitude: 21.03429931508327,
        longitude: 105.8507011685404
    }
    const end = {
        latitude: 21.01826747575612,
        longitude: 105.85388186669118
    }

    const x = [1, 2, 3]
    console.log(x);
    console.log(x.slice(2, 5));

    const { t } = useTranslation();

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
                    coordinate={destination}>
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
        </View>
    )
}

export default TripDirection