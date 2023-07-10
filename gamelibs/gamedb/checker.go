package gamedb

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Checkable interface {
	check() string
}

type tagChecker struct {
	fn   func(a ...interface{}) bool
	desc string
}

var tagCheckerMap = make(map[string]tagChecker)

var itemTypeChecker = make(map[int]func(a ...interface{}) bool)

func init() {
	tagCheckerMap["item"] = newTagChecker(itemExist, "item.xlsx 中没有物品 或者数量不对: %d")
	tagCheckerMap["prop"] = newTagChecker(propExist, "property.xlsx 中没有属性,或者数值不对 %d")
	tagCheckerMap["task"] = newTagChecker(taskExist, "任务中没有 %d")
	tagCheckerMap["condition"] = newTagChecker(conditionExist, "条件中没有 %d (如果是32,可能是任务id不存在)")

}

func newTagChecker(fn func(a ...interface{}) bool, desc string) tagChecker {
	return tagChecker{fn: fn, desc: desc}
}

func (this *GameDb) Check() error {

	checkers := []func(*GameDb) error{
		//(*GameDb).checkConf,
	}
	for _, checker := range checkers {
		err := checker(this)
		if err != nil {
			return err
		}
	}
	err := this.checkAll()
	if err != nil {
		return err
	}

	return nil
}

func (this *GameDb) checkAll() error {
	err := checkGameDbBase()
	if err != nil {
		return err
	}
	err = checkAllViaObj(GetConf())
	if err != nil {
		return err
	}
	err = this.checkPlot()
	if err != nil {
		return err
	}
	return nil

}

func checkFunc(value interface{}) (bool, []string) {
	allErrorMsg := make([]string, 0)
	hasCheck := false
	//如果结构体实现了Checkable
	if c, ok := value.(Checkable); ok {
		errMsg := c.check()
		if len(errMsg) > 0 {
			allErrorMsg = append(allErrorMsg, errMsg)
		}
		hasCheck = true
	}
	//如果，去检查Tag 中配置的检查
	typeValue := reflect.ValueOf(value)
	if gameDb.hasTagCheck(typeValue) {
		err := gameDb.checkOneDataMember(typeValue)
		if err != nil {
			allErrorMsg = append(allErrorMsg, "Struct:"+reflect.TypeOf(value).String()+" "+err.Error())
		}
		hasCheck = true
	}
	return hasCheck, allErrorMsg
}

func checkGameDbBase() error {

	allErrorMsg := make([]string, 0)
	gameDbBaseValueOf := reflect.ValueOf(reflect.ValueOf(gameDb).Elem().FieldByName("GameDbBase").Interface()).Elem()

	for j := 0; j < gameDbBaseValueOf.NumField(); j++ {
		f := gameDbBaseValueOf.Field(j)
		fKind := f.Kind()
		if fKind == reflect.Map {
			for _, v := range f.MapKeys() {
				value := f.MapIndex(v).Interface()
				hasCheck, err := checkFunc(value)
				if !hasCheck {
					break
				}
				if len(err) > 0 {
					allErrorMsg = append(allErrorMsg, err...)
				}
			}
		}
	}

	if len(allErrorMsg) < 1 {
		return nil
	}

	return errors.New(strings.Join(allErrorMsg, "\n"))
}

func checkAllViaObj(objBeChecker interface{}) error {

	values := reflect.ValueOf(objBeChecker).Elem()

	allErrorMsg := make([]string, 0)

	for i := 0; i < values.NumField(); i++ {
		f := values.Field(i)
		fKind := f.Kind()
		switch fKind {
		case reflect.Slice:
			for i := 0; i < f.Len(); i++ {
				value := f.Index(i).Interface()
				hasCheck, err := checkFunc(value)
				if !hasCheck {
					break
				}
				if len(err) > 0 {
					allErrorMsg = append(allErrorMsg, err...)
				}
			}
		case reflect.Map:
			for _, v := range f.MapKeys() {
				value := f.MapIndex(v).Interface()
				hasCheck, err := checkFunc(value)
				if !hasCheck {
					break
				}
				if len(err) > 0 {
					allErrorMsg = append(allErrorMsg, err...)
				}
			}
		case reflect.Ptr:
			hasCheck, err := checkFunc(f.Interface())
			if !hasCheck {
				break
			}
			if len(err) > 0 {
				allErrorMsg = append(allErrorMsg, err...)
			}

		default:
			//fmt.Println("unhandle type in gamedb:", fKind)
		}
	}
	if len(allErrorMsg) < 1 {
		return nil
	}

	return errors.New(strings.Join(allErrorMsg, "\n"))
}

//公用的检查方法
func propExist(args ...interface{}) bool {
	//k := args[0].(int)
	//if gameDb.GetProperty(k) == nil {
	//	return false
	//}
	//if len(args) > 1 {
	//	count := args[1].(int)
	//	if count < 1 {
	//
	//	}
	//}
	return true
}

func itemExist(args ...interface{}) bool {
	return itemExistWithFlag(false, args...)
}

func itemExistWithFlag(allowZero bool, args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}
	itemId := args[0].(int)
	if itemId == 0 {
		return true
	}
	itemT := GetItemBaseCfg(itemId)
	if itemT == nil {
		//logger.Error("道具不存在：%v", itemId)
		return false
	}
	if len(args) > 1 {
		count := args[1].(int)
		if !allowZero && count < 1 {
			//logger.Error("道具数量不能为0，道具Id:%v", itemId)
			return false
		}
	}
	fn := itemTypeChecker[itemT.Type]
	if fn != nil {
		flag := fn(itemId)
		if !flag {
			fmt.Printf("aaa3 itemType=%v %v", itemT.Type, args)
		}
		return flag
	}
	return true
}

func taskExist(args ...interface{}) bool {
	//id := args[0].(int)
	//if id == 0 {
	//	return true
	//}
	//return gameDb.GetTask(id) != nil
	return true
}

func conditionExist(args ...interface{}) bool {
	//id := args[0].(int)
	//if gameDb.GetCondition(id) == nil {
	//	return false
	//}
	//if len(args) < 2 {
	//	fmt.Printf("conditionExist:why args is %v", args)
	//	return false
	//}
	//value := args[1].(int)
	//
	//if len(args) > 2 {
	//	subs := args[2].(map[int]int)
	//	for subType, v := range subs {
	//		if gameDb.ConditionSubTypes[id*100+subType] == nil || v < 1 {
	//			return false
	//		}
	//	}
	//}
	//
	//switch id {
	//case pb.CONDITIONTYPE_FINISH_MAIN_LINE_TASK:
	//	return taskExist(value)
	//default:
	//	return true
	//}
	return true
}

func itemsExist(propInfos PropInfos) bool {
	if len(propInfos) < 1 {
		return false
	}
	for _, v := range propInfos {
		if !itemExist(v.K) {
			return false
		}
	}
	return true
}

func checkWalkable(sceneT *SceneConf, points PropInfos, copyId int, msg string) string {
	if len(points) < 1 {
		return ""
	}
	for _, p := range points {
		if !sceneT.Walkable(int32(p.K), int32(p.V)) {
			return fmt.Sprintf(msg, copyId, sceneT.Id, p.K, p.V)
		}
	}
	return ""
}

func (this *GameDb) hasTagCheck(v reflect.Value) bool {
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.Elem().Kind() == reflect.Struct {
		for j := 0; j < v.Elem().NumField(); j++ {
			if len(v.Elem().Type().Field(j).Tag.Get("checker")) > 0 {
				return true
			}
		}
	}
	return false
}

func (this *GameDb) checkOneFiledTest() {
	//for _, v := range this.Tasks {
	//	err := this.checkOneDataMember(reflect.ValueOf(v))
	//	if err != nil {
	//		fmt.Printf("err = %+v\n", err)
	//	}
	//	break
	//}
}

func checkOneId(errMsgs *[]string, tc tagChecker, fieldName string, args ...interface{}) {
	if !tc.fn(args...) {
		*errMsgs = append(*errMsgs, fmt.Sprintf("字段 %s 在%v", fieldName, fmt.Sprintf(tc.desc, args)))
	}
}

func (this *GameDb) checkOneDataMember(v reflect.Value) error {

	errMsgsSlice := make([]string, 0)
	errMsgs := &errMsgsSlice

	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.Elem().Kind() == reflect.Struct {
		for j := 0; j < v.Elem().NumField(); j++ {
			checkerName := v.Elem().Type().Field(j).Tag.Get("checker")
			fieldType := v.Elem().Type().Field(j).Type.Name()
			fieldName := v.Elem().Type().Field(j).Name
			checkerName = strings.TrimSpace(checkerName)
			if len(checkerName) < 1 {
				continue
			}
			checkerNameArr := strings.Split(checkerName, "_")
			checkerFnName := checkerNameArr[0]
			checkerNameExtra := checkerNameArr[1:]
			checkerFnName = strings.ToLower(checkerFnName)
			var tagc tagChecker
			if _, ok := tagCheckerMap[checkerFnName]; ok {
				tagc = tagCheckerMap[checkerFnName]
			} else {
				fmt.Printf("No checker for %s\n", checkerFnName)
				continue
			}

			switch fieldType {
			case "PropInfo":
				data := v.Elem().Field(j).Interface().(PropInfo)
				checkOneId(errMsgs, tagc, fieldName, data.K, data.V)
			case "PropInfos":
				data := v.Elem().Field(j).Interface().(PropInfos)
				for m := 0; m < len(data); m++ {
					checkOneId(errMsgs, tagc, fieldName, (data)[m].K, (data)[m].V)
				}
			case "int":
				id := int(v.Elem().Field(j).Int())
				if id == 0 {
					continue
				}
				checkOneId(errMsgs, tagc, fieldName, id)
			case "ProbItems":
				data := v.Elem().Field(j).Interface().(ProbItems)
				for m := 0; m < len(data); m++ {
					checkOneId(errMsgs, tagc, fieldName, (data)[m].Id, (data)[m].Count)
				}
			case "ProbValues":
				data := v.Elem().Field(j).Interface().(ProbValues)
				for m := 0; m < len(data); m++ {
					checkOneId(errMsgs, tagc, fieldName, (data)[m].Value, (data)[m].Prob)
				}
			case "IntSlice":
				data := v.Elem().Field(j).Interface().(IntSlice)
				for m := 0; m < len(data); m++ {
					checkOneId(errMsgs, tagc, fieldName, data[m])
				}
			case "IntMap":
				data := v.Elem().Field(j).Interface().(IntMap)
				for k, v := range data {
					checkOneId(errMsgs, tagc, fieldName, k, v)
				}
			case "ItemInfos":
				data := v.Elem().Field(j).Interface().(ItemInfos)
				for m := 0; m < len(data); m++ {
					checkOneId(errMsgs, tagc, fieldName, (data)[m].ItemId, (data)[m].Count)
				}
			case "ItemInfo":
				data := v.Elem().Field(j).Interface().(ItemInfo)
				checkOneId(errMsgs, tagc, fieldName, (data).ItemId, (data).Count)

			case "string":
				str := v.Elem().Field(j).String()
				checkOneId(errMsgs, tagc, fieldName, str, checkerNameExtra)
			case "StringSlice":
				data := v.Elem().Field(j).Interface().(StringSlice)
				for m := 0; m < len(data); m++ {
					checkOneId(errMsgs, tagc, fieldName, data[m], checkerNameExtra)
				}

			case "Condition":
				data := v.Elem().Field(j).Interface().(Condition)
				checkOneId(errMsgs, tagc, fieldName, data.K, data.V, data.Subs)
			case "Conditions":
				data := v.Elem().Field(j).Interface().(Conditions)
				for m := 0; m < len(data); m++ {
					checkOneId(errMsgs, tagc, fieldName, (data)[m].K, (data)[m].V, (data)[m].Subs)
				}
			case "MedalCheckId":
				id := int(v.Elem().Field(j).Int())
				checkOneId(errMsgs, tagc, fieldName, id)
			default:
				//fmt.Printf("fieldType = %+v no handler ,for fieldName %s \n", fieldType, fieldName)
			}
		}
	}
	if len(*errMsgs) < 1 {
		return nil
	}
	return errors.New(strings.Join(*errMsgs, "\n"))
}
