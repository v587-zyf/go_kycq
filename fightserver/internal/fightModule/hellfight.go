package fightModule

type HellFight struct {
	*DefaultFight
}

func NewHellFight(stageId int) (*HellFight, error) {
	fight := &HellFight{}
	var err error
	fight.DefaultFight, err = NewDefaultFight(stageId, fight)
	if err != nil {
		return nil, err
	}
	fight.SetLifeTime(-1)
	fight.InitCollection()
	fight.Start()
	return fight, nil
}
