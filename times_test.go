package frm


import (
	"testing"
)


func TestTimesSeries(t *testing.T) {

	var series Series

	series = Times{}

	if nil == series {
		t.Errorf("This should never happen.")
		return
	}
}
