package common

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestSlice(t *testing.T) {
	src := []int{1, 2, 3, 4, 4, 5, 3, 10}

	ret := ConvertIntSlice2Int32Slice(src)

	for k, v := range ret {
		if src[k] != int(v) {
			t.Fatal("convert doesn't work")
		}
	}

	uniquned := SliceInt32Unique(ret)

	rightResult := []int32{1, 2, 3, 4, 5, 10}
	if !reflect.DeepEqual(uniquned, rightResult) {
		t.Fatal("unique doesn't work")
	}

}

func AddSliceElm(s []int, a int, t *testing.T) {
	addr := unsafe.Pointer(&s)
	fmt.Println(addr)
	s = append(s, a)
	fmt.Println(cap(s))
}

func TestSliceAppend(t *testing.T) {
	var s []int = make([]int, 3)
	addr := unsafe.Pointer(&s)

	fmt.Println(addr)
	t.Log(addr)
	s = append(s, 1)
	fmt.Println(cap(s))
	AddSliceElm(s, 2, t)
	fmt.Println(cap(s))
}

func TestWeekDay(t *testing.T) {

	now := time.Now().Weekday()
	fmt.Println(now)

	//context.WithTimeout()

}

func TestBreak(t *testing.T) {

	var v = []int{1, 2, 3, 4, 5, 6, 7}
	var idx int
FOR0:
	for i1 := 0; i1 < 10; i1++ {
		for i, itr := range v {
			idx = i
			if itr == 3 {
				break FOR0
			}
		}
	}

	fmt.Println(idx)

	//context.WithTimeout()

}

func TestSector(t *testing.T) {

	fmt.Println(math.Abs(1))

}

func TestRemove(t *testing.T) {
	a := []int{2, 3, 4, 1}
	for idx, v := range a {
		if v == 1 {
			a = append(a[0:idx], a[idx+1:]...)
		}
	}
	a = append(a, 2)
	for idx, v := range a {
		if v == 2 {
			a = append(a[0:idx], a[idx+1:]...)
		}
	}

}

func TestRemove2(t *testing.T) {

	a := []int{1, 2, 3, 4, 2, 5, 2, 1, 1, 2, 3, 4, 2, 5, 2, 2, 2}
	b := a

	for idx := len(b) - 1; idx >= 0; idx-- {
		v := a[idx]
		fmt.Println(idx, v, len(a))
		if v == 2 {
			a = append(a[:idx], a[idx+1:]...)
		}
	}
	fmt.Println(a)

}

type TestValueType struct {
	k int
	v int
}

func TestOrder(t *testing.T) {
	v := []*TestValueType{
		&TestValueType{1, 1},
		&TestValueType{2, 2},
		&TestValueType{6, 6},
		&TestValueType{3, 3},
		&TestValueType{4, 4},
		&TestValueType{5, 5},
	}
	var sv []*TestValueType
	var idx2 int
	var itr2 *TestValueType
	var bFind bool
	for _, itr := range v {
		idx2 = 0
		bFind = false
		for idx2, itr2 = range sv {
			if itr.k <= itr2.k && itr.v <= itr2.v {
				bFind = true
				break
			}
		}
		if bFind {
			sv = append(sv[:idx2], itr)
			sv = append(sv, sv[idx2:]...)
		} else {
			sv = append(sv, itr)
		}
	}

	for _, itr := range sv {
		fmt.Println(itr)
	}

}
