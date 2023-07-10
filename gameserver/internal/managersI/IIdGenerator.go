package managersI

type IIdGeneratorManager interface {

	InitWorldNowId(nowId, guildNowId int)

	GetNextWorldId() int

	GetNextGuildId() int

}
