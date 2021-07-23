package unicast

import (
	"google.golang.org/protobuf/proto"
)

func WrapPlainTxt(txt string) ([]byte, error) {
	chatMessage := &ChatMessage{
		Payload: &ChatMessage_PlainTxt{PlainTxt: txt},
	}
	rawData, err := proto.Marshal(chatMessage)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapImage(data []byte) ([]byte, error) {
	chatMessage := &ChatMessage{
		Payload: &ChatMessage_Image{Image: data},
	}
	rawData, err := proto.Marshal(chatMessage)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapLocation(l, a float32, name string) ([]byte, error) {
	chatMessage := &ChatMessage{
		Payload: &ChatMessage_Location{Location: &Location{
			Latitude:  a,
			Longitude: l,
			Name:      name,
		}},
	}

	rawData, err := proto.Marshal(chatMessage)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapVoice(p []byte, l int) ([]byte, error) {
	chatMessage := &ChatMessage{
		Payload: &ChatMessage_Voice{Voice: &Voice{
			Data:   p,
			Length: int32(l),
		}},
	}

	rawData, err := proto.Marshal(chatMessage)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func WrapSyncGroup(groupId string) ([]byte, error) {
	chatMessage := &ChatMessage{
		Payload: &ChatMessage_SyncGroupId{SyncGroupId: groupId},
	}

	rawData, err := proto.Marshal(chatMessage)
	if err != nil {
		return nil, err
	}

	return rawData, nil

}
