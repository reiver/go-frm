package frm


import (
	"testing"
)


func TestBoolsTail(t *testing.T) {

	tests := []struct{
		Slice    []bool
		Expected []bool
	}{
		{
			Slice:    nil,
			Expected: []bool{},
		},



		{
			Slice:    []bool{},
			Expected: []bool{},
		},



		{
			Slice:    []bool{false,},
			Expected: []bool{false,},
		},
		{
			Slice:    []bool{true,},
			Expected: []bool{true,},
		},



		{
			Slice:    []bool{false, false,},
			Expected: []bool{false, false,},
		},
		{
			Slice:    []bool{false, true,},
			Expected: []bool{false, true,},
		},
		{
			Slice:    []bool{true, false,},
			Expected: []bool{true, false,},
		},
		{
			Slice:    []bool{true, true,},
			Expected: []bool{true, true,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false,},
			Expected: []bool{false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,},
			Expected: []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,},
		},

		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false,},
			Expected: []bool{false, true,  false, false, true,  true,  false, false, false,},
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,},
			Expected: []bool{true,  false, true,  true,  false, false, true,  true,  true,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false,},
			Expected: []bool{false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,},
			Expected: []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,},
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true, },
			Expected: []bool{false, true,  false, false, true,  true,  false, false, false, true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false,},
			Expected: []bool{true,  false, true,  true,  false, false, true,  true,  true,  false,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false, false,},
			Expected: []bool{       false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true,},
			Expected: []bool{       true,  true,  true,  true,  true,  true,  true,  true,  true,  true,},
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true,  true, },
			Expected: []bool{       true,  false, false, true,  true,  false, false, false, true, true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false, false,},
			Expected: []bool{       false, true,  true,  false, false, true,  true,  true,  false, false,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false, false, false,},
			Expected: []bool{              false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true,},
			Expected: []bool{              true,  true,  true,  true,  true,  true,  true,  true,  true,  true,},
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true,  true,  true, },
			Expected: []bool{              false, false, true,  true,  false, false, false, true,  true,  true, },
		},
		{
			Slice:    []bool{true, false,  true,  true,  false, false, true,  true,  true,  false, false, false,},
			Expected: []bool{              true,  true,  false, false, true,  true,  true,  false, false, false,},
		},
	}


	for testNumber, test := range tests {

		s := Bools(test.Slice)

		head := s.Tail()

		if expected, actual := len(test.Expected), len(head.(Bools)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Expected {
			actual := head.(Bools)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%t\", but actually got \"%t\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}
