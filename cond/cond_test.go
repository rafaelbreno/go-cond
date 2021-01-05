package cond_test

import (
	"github.com/rafaelbreno/go-cond/cond"
	"reflect"
	"testing"
)

func TestNewCondition(t *testing.T) {
	t.Helper()

	want := cond.Condition{}

	if got := cond.NewCondition(); !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v\nGot %v", want, got)
	}

}

func TestAnd(t *testing.T) {
	t.Helper()

	want := false

	c := cond.NewCondition()

	got := c.And(true, false, false, false, true).CurrentVal

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want %v\nGot %v", want, got)
	}
}
