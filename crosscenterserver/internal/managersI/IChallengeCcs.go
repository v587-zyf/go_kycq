package managersI

type IChallengeCcs interface {
	BeginGroup(roundIndex int)

	InSeasonCompetition(roundIndex int)

	ToGsUp(crossFsId int)

	BeginInsertRobotUser()
}
