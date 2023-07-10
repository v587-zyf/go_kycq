package builder

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/protobuf/pb"
)

func BuildMailNtfs(mails []*modelGame.Mail) []*pb.MailNtf {
	msgs := make([]*pb.MailNtf, len(mails))
	for i, v := range mails {
		msgs[i] = BuildMailNtf(v)
	}
	return msgs
}

func BuildMailNtf(m *modelGame.Mail) *pb.MailNtf {

	items := make([]*pb.ItemUnit, 0)
	if len(m.Items) > 0 {
		for _, v := range m.Items {
			pbItem := &pb.ItemUnit{ItemId: int32(v.ItemId), Count: int64(v.Count)}
			items = append(items, pbItem)
		}
	}
	var redeemedAt int32
	if !m.RedeemedAt.IsZero() {
		redeemedAt = int32(m.RedeemedAt.Unix())
	}
	return &pb.MailNtf{
		Id:         int32(m.Id),
		Sender:     m.Sender,
		Title:      m.Title,
		Content:    m.Content,
		Status:     int32(m.Status),
		ExpireAt:   int32(m.ExpireAt.Unix()),
		CreatedAt:  int32(m.CreatedAt.Unix()),
		RedeemedAt: redeemedAt,
		Args:       m.Args,
		Items:      items,
	}
}
