package conf

import "time"

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second * 100
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 4096
)

const (
	STATUS_NONE int32 = iota
	STATUS_CONNECTED
	STATUS_ENTERGAME_ING
	STATUS_ENTERGAME_DONE
	STATUS_CREATE_ROLE_ING
	STATUS_CREATE_ROLE_DONE
	STATUS_IN_GAME
	STATUS_OFF
	STATUS_CONNECTED2 int32 = 50
)