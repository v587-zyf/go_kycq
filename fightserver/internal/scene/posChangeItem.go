package scene

const (
	PosChangeTypeAdd    = 1
	PosChangeTypeRemove = 2
	PosChangeTypeMove   = 3
	PosChangeTypeUpdate = 4
	PosChangeTypeRelive = 5
	PosChangeTypeAdds   = 6
)

type posChangeItem struct {
	typ   int
	obj   ISceneObj
	point *Point
	arg   []interface{}
}

func newPosChangeItem(typ int, obj ISceneObj, point *Point, arg ...interface{}) *posChangeItem {
	return &posChangeItem{
		typ:   typ,
		obj:   obj,
		point: point,
		arg:   arg,
	}
}
