package fightModule

type MainCity struct {
	*DefaultFight

	birthPointX, birthPointY int
}

func NewMainCity(stageId int) (*MainCity, error) {

	maincity := &MainCity{}
	var err error
	maincity.DefaultFight, err = NewDefaultFight(stageId, maincity)
	if err != nil {
		return nil, err
	}
	maincity.SetLifeTime(-1)
	maincity.Start()
	return maincity, nil
}
