package handler

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managers"
)

var (
	m = managers.Get()
)

func gameDb() *gamedb.GameDb {
	return gamedb.GetDb()
}
