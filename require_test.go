package require_field

import (
	"strings"
	"testing"
)

func TestRequire(t *testing.T) {
	type T struct {
		Home     string `required:"t"`
		Optional string
	}

	var i int
	if err := Check(i); err != nil {
		t.Errorf("err must be nil. i is %T", i)
	}

	var st = &T{"", "one"}
	err := Check(st)

	if err == nil {
		t.Errorf("err cannot be nil. Struct T have required field")
	}

	if !strings.Contains(err.Error(), "Home") {
		t.Errorf("err must contain Home; have: %s", err.Error())
	}

	type T2 struct {
		Other string
		T     T `required:"t"`
	}

	var st2 = T2{Other: "", T: T{
		Optional: "one",
	}}
	err = Check(st2)

	if err == nil {
		t.Errorf("err cannot be nil. Struct T2 have required field T")
	}

	if !strings.Contains(err.Error(), "Home") {
		t.Errorf("err must contain Home; have: %s", err.Error())
	}
}
