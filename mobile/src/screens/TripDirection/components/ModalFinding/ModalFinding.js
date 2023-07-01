import { ActivityIndicator, Modal, Text, TouchableOpacity, View } from "react-native"


const ModalFinding = ({ show, onPressCancel }) => {

    return (
        <Modal
            visible={show}
            transparent={true}
            statusBarTranslucent={true}>
            <View className="flex-1 bg-black/30 items-center justify-center px-6">
                <View className='bg-white p-1 w-[240px] rounded-sm'>
                    <View className={'scale-[3] mb-8 mt-8'}>
                        <ActivityIndicator color={"rgb(234,179,8)"} />
                    </View>
                    <Text className={"text-center font-semibold mb-3"}>{"Đang tìm tài xế"}</Text>
                    <TouchableOpacity
                        className="bg-yellow-400 rounded px-2 py-2 items-center justify-center"
                        activeOpacity={0.7}
                        onPress={onPressCancel}>
                        <Text className="font-semibold text-white">
                            Hủy
                        </Text>
                    </TouchableOpacity>
                </View>
            </View>
        </Modal>
    )
}

export default ModalFinding