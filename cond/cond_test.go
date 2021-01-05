package cond_test

import (
	"github.com/rafaelbreno/go-cond/cond"
	"reflect"
	"testing"
)

func testNewCondition(t *testing.T) {
	t.Helper()

	want := cond.Condition{}

	if got := cond.NewCondition(); !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v\nGot %v", want, got)
	}

}
