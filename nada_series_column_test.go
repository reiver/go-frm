package frm


import (
	"testing"
)


func TestNadaSeriesColumn0(t *testing.T) {

	tests := []struct{
		Slice []struct{}
	}{
		{
			Slice:    nil,
		},



		{
			Slice:    []struct{}{},
		},



		{
			Slice:    []struct{}{ struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},
	}


	for testNumber, test := range tests {

		s := NadaSeries( test.Slice )

		c := s.Column(0)

		if expected, actual := len(test.Slice), len(c.(NadaSeries)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Slice {
			actual := c.(NadaSeries)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%v\", but actually got \"%v\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}


func TestNadaSeriesColumnNot0(t *testing.T) {

	tests := []struct{
		Slice []struct{}
	}{
		{
			Slice:    nil,
		},



		{
			Slice:    []struct{}{},
		},



		{
			Slice:    []struct{}{ struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},
	}


	for testNumber, test := range tests {

		for columnNumber := -11; columnNumber < 12; columnNumber++ {

			if 0 == columnNumber {
				continue
			}

			s := NadaSeries( test.Slice )

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
