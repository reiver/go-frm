package frm


import (
	"testing"
)


func TestFloat64sColumn0(t *testing.T) {

	tests := []struct{
		Slice []float64
	}{
		{
			Slice:    nil,
		},



		{
			Slice:    []float64{},
		},



		{
			Slice:    []float64{1.1,},
		},



		{
			Slice:    []float64{1.1, 2.2,},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9,},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11, 12.12},
		},
	}


	for testNumber, test := range tests {

		s := Float64s( test.Slice )

		c := s.Column(0)

		if expected, actual := len(test.Slice), len(c.(Float64s)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Slice {
			actual := c.(Float64s)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%f\", but actually got \"%f\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}


func TestFloat64sColumnNot0(t *testing.T) {

	tests := []struct{
		Slice []float64
	}{
		{
			Slice:    nil,
		},



		{
			Slice:    []float64{},
		},



		{
			Slice:    []float64{1.1,},
		},



		{
			Slice:    []float64{1.1, 2.2,},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9,},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11},
		},



		{
			Slice:    []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1, 11.11, 12.12},
		},
	}


	for testNumber, test := range tests {

		for columnNumber := -11; columnNumber < 12; columnNumber++ {

			if 0 == columnNumber {
				continue
			}

			s := Float64s( test.Slice )

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
