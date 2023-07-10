package fightModule

type DarkPalace struct {
	*DefaultFight
}

func NewDarkPalace(stageId int) (*DarkPalace, error) {

	darkPalace := &DarkPalace{}
	var err error
	darkPalace.DefaultFight, err = NewDefaultFight(stageId, darkPalace)
	if err != nil {
		return nil, err
	}
	darkPalace.SetLifeTime(-1)
	darkPalace.InitCollection()
	darkPalace.Start()
	return darkPalace, nil
}

func (this *DarkPalace) OnCollection(colllection map[int]int) {

}
