package system

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func (this *SystemManager) PreferenceSet(user *objs.User, preferences []*pb.Preference, ack *pb.PreferenceSetAck) error {
	userPreference := user.Preference
	for _, preference := range preferences {
		userPreference[int(preference.Key)] = preference.Value
	}
	user.Dirty = true
	ack.Preference = preferences
	return nil
}

func (this *SystemManager) PreferenceLoad(user *objs.User, ack *pb.PreferenceLoadAck) error {
	for key, value := range user.Preference {
		ack.Preference = append(ack.Preference, &pb.Preference{
			Key:   int32(key),
			Value: value,
		})
	}
	return nil
}
