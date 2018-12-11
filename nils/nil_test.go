package nils

import "testing"

func TestC_GetAName(t *testing.T) {
	var c Third

	if c.GetFirstName() != "none" {
		t.Error("can't get none string")
	}
}
