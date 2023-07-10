package base

import (
	"errors"
)

var (
	ErrNoGameServerAssigned = errors.New("no gameserver assigned")
	ErrNoFsServerAssigned   = errors.New("no fightserver assigned")
	ErrServerNotConnected   = errors.New("no fightserver not connected")
)
