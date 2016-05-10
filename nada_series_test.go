package frm


import (
	"testing"
)


func TestNadaSeriesSeries(t *testing.T) {

	var series Series

	series = NadaSeries{}

	if nil == series {
		t.Errorf("This should never happen.")
		return
	}
}
