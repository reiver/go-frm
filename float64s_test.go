package frm


import (
	"testing"
)


func TestFloat64sSeries(t *testing.T) {

	var series Series

	series = Float64s{}

	if nil == series {
		t.Errorf("This should never happen.")
		return
	}
}
