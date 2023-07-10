package ai

import "cqserver/fightserver/internal/base"

type IFightStageInterface interface {
	getFightStageTarget() base.Actor
}
