package frm


import (
	"testing"
)


func TestBoolsSeries(t *testing.T) {

	var series Series

	series = Bools{}

	if nil == series {
		t.Errorf("This should never happen.")
		return
	}
}
