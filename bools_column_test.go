package frm


import (
	"testing"
)


func TestBoolsColumn0(t *testing.T) {

	tests := []struct{
		Slice []bool
	}{
		{
			Slice:    nil,
		},



		{
			Slice:    []bool{},
		},



		{
			Slice:    []bool{false,},
		},
		{
			Slice:    []bool{true, },
		},



		{
			Slice:    []bool{false, false,},
		},
		{
			Slice:    []bool{false, true, },
		},
		{
			Slice:    []bool{true,  false,},
		},
		{
			Slice:    []bool{true,  true, },
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false,},
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true,  true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false, false,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true,  true,  true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false, false, false, },
		},
	}


	for testNumber, test := range tests {

		s := Bools( test.Slice )

		c := s.Column(0)

		if expected, actual := len(test.Slice), len(c.(Bools)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Slice {
			actual := c.(Bools)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%f\", but actually got \"%f\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}


func TestBoolsColumnNot0(t *testing.T) {

	tests := []struct{
		Slice []bool
	}{
		{
			Slice:    nil,
		},



		{
			Slice:    []bool{},
		},



		{
			Slice:    []bool{false,},
		},
		{
			Slice:    []bool{true, },
		},



		{
			Slice:    []bool{false, false,},
		},
		{
			Slice:    []bool{false, true, },
		},
		{
			Slice:    []bool{true,  false,},
		},
		{
			Slice:    []bool{true,  true, },
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false,},
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true,  true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false, false,},
		},



		{
			Slice:    []bool{false, false, false, false, false, false, false, false, false, false, false, false,},
		},
		{
			Slice:    []bool{true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true,  true, },
		},
		{
			Slice:    []bool{false, true,  false, false, true,  true,  false, false, false, true,  true,  true, },
		},
		{
			Slice:    []bool{true,  false, true,  true,  false, false, true,  true,  true,  false, false, false, },
		},
	}


	for testNumber, test := range tests {

		for columnNumber := -11; columnNumber < 12; columnNumber++ {

			if 0 == columnNumber {
				continue
			}

			s := Bools( test.Slice )

			c := s.Column(columnNumber)

			if _, ok := c.(NadaSeries); !ok {
				t.Errorf("For test #%d and column #%d, expected series to be NadaSeries, but actually was %T.", testNumber, columnNumber, c)
				continue
			}

			if expected, actual := len(test.Slice), len(c.(NadaSeries)); expected != actual {
				t.Errorf("For test #%d and column #%d, expected %d, but actually got %d.", testNumber, columnNumber, expected, actual)
				continue
			}
		}
	}
}
