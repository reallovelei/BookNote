package _interface

import (
	"testing"
)

type TestStruct struct {
}

func NilOrNot(v interface{}) bool {
	return v == nil
}

func TestNil(t *testing.T) {
	var s *TestStruct
	t.Logf("%+v", s)
	t.Logf("%+v", NilOrNot(s))
}
