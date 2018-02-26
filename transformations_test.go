package main

import (
	"testing"
)

func TestTrimRight(t *testing.T) {
	beforestring := "6887Hi:89888"
	afterstring := trimMrnFromColon(beforestring)
	if afterstring != "89888" {
		t.Errorf("trim mrn from colon failed")
	}

}
