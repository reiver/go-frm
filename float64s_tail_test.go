package frm


import (
	"testing"
)


func TestFloat64sTail(t *testing.T) {

	tests := []struct{
		Slice    []float64
		Expected []float64
	}{
		{
			Slice:    nil,
			Expected: []float64{},
		},



		{
			Slice:    []float64{},
			Expected: []float64{},
		},



		{
			Slice:    []float64{1.1,},
			Expected: []float64{1.1,},
		},



		{
			Slice:    []float64{1.1, 2.2,},
			Expected: []float64{1.1, 2.2,},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9,},
			Expected: []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9,},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1},
			Expected: []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11},
			Expected: []float64{     2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11, 12.12},
			Expected: []float64{          3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11, 12.12},
		},
	}


	for testNumber, test := range tests {

		s := Float64s( test.Slice )

		tail := s.Tail()

		if expected, actual := len(test.Expected), len(tail.(Float64s)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Expected {
			actual := tail.(Float64s)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%f\", but actually got \"%f\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}
