package objs

import (
	"cqserver/gamelibs/modelGame"
)

type Mining struct {
	*modelGame.MiningDb

	Combat int //机器人战力
}

func NewMining(mining *modelGame.MiningDb) *Mining {
	m := &Mining{
		MiningDb: mining,
	}
	return m
}
