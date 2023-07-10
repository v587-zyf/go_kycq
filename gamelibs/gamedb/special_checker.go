package gamedb

import (
	"errors"
)

func (this *GameDb) checkPlot() error {

	errs := ""

	if len(errs) != 0 {
		return errors.New(errs)
	}
	return nil
}
