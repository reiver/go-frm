package frm


import (
	"testing"
)


func TestFloat64sHead(t *testing.T) {

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
			Expected: []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11, 12.12},
			Expected: []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1},
		},
	}


	for testNumber, test := range tests {

		s := Float64s( test.Slice )

		head := s.Head()

		if expected, actual := len(test.Expected), len(head.(Float64s)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Expected {
			actual := head.(Float64s)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%f\", but actually got \"%f\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}
