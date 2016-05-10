package frm


import (
	"testing"
)


func TestNadaSeriesHead(t *testing.T) {

	tests := []struct{
		Slice    []struct{}
		Expected []struct{}
	}{
		{
			Slice:    nil,
			Expected: []struct{}{},
		},



		{
			Slice:    []struct{}{},
			Expected: []struct{}{},
		},



		{
			Slice:    []struct{}{ struct{}{}, },
			Expected: []struct{}{ struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, },
			Expected: []struct{}{ struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
			Expected: []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
			Expected: []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
			Expected: []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},



		{
			Slice:    []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
			Expected: []struct{}{ struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, struct{}{}, },
		},
	}


	for testNumber, test := range tests {

		s := NadaSeries( test.Slice )

		head := s.Head()

		if expected, actual := len(test.Expected), len(head.(NadaSeries)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Expected {
			actual := head.(NadaSeries)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%v\", but actually got \"%v\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}
