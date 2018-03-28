package disjointSetTree

import (
	"testing"
	"reflect"
	"fmt"
)

func TestOfflineMinimum(t *testing.T) {
	seq := []int{4, 8,
	OFFLINEMINIMUM_EXTRACT,
	3,
	OFFLINEMINIMUM_EXTRACT,
	9,2,6,
	OFFLINEMINIMUM_EXTRACT,
	OFFLINEMINIMUM_EXTRACT,
	OFFLINEMINIMUM_EXTRACT,
	1,7,
	OFFLINEMINIMUM_EXTRACT,
	5}
	exp := []int{4,3,2,6,8,1}
	extractSeq := offLineMinimum(seq)
	if !reflect.DeepEqual(extractSeq, exp) {
		t.Log(fmt.Sprintf("expct :%+v, actual:%+v", exp, extractSeq))
		t.Fail()
	}
}