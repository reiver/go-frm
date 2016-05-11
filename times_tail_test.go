package frm


import (
	"time"

	"testing"
)


func TestTimessTail(t *testing.T) {

	now := time.Now()


	tests := []struct{
		Slice    []time.Time
		Expected []time.Time
	}{
		{
			Slice:    nil,
			Expected: []time.Time{},
		},



		{
			Slice:    []time.Time{},
			Expected: []time.Time{},
		},



		{
			Slice:    []time.Time{ now.Add(1 * time.Hour), },
			Expected: []time.Time{ now.Add(1 * time.Hour), },
		},



		{
			Slice:    []time.Time{ now.Add(1 * time.Hour), now.Add(2 * time.Hour), },
			Expected: []time.Time{ now.Add(1 * time.Hour), now.Add(2 * time.Hour), },
		},



		{
			Slice:    []time.Time{
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
				now.Add(3 * time.Hour),
				now.Add(4 * time.Hour),
				now.Add(5 * time.Hour),
				now.Add(6 * time.Hour),
				now.Add(7 * time.Hour),
				now.Add(8 * time.Hour),
				now.Add(9 * time.Hour),
			},
			Expected: []time.Time{
				now.Add(1 * time.Hour),
				now.Add(2 * time.Hour),
				now.Add(3 * time.Hour),
				now.Add(4 * time.Hour),
				now.Add(5 * time.Hour),
				now.Add(6 * time.Hour),
				now.Add(7 * time.Hour),
				now.Add(8 * time.Hour),
				now.Add(9 * time.Hour),
			},
		},



		{
			Slice:    []time.Time{
				now.Add( 1 * time.Hour),
				now.Add( 2 * time.Hour),
				now.Add( 3 * time.Hour),
				now.Add( 4 * time.Hour),
				now.Add( 5 * time.Hour),
				now.Add( 6 * time.Hour),
				now.Add( 7 * time.Hour),
				now.Add( 8 * time.Hour),
				now.Add( 9 * time.Hour),
				now.Add(10 * time.Hour),
			},
			Expected: []time.Time{
				now.Add( 1 * time.Hour),
				now.Add( 2 * time.Hour),
				now.Add( 3 * time.Hour),
				now.Add( 4 * time.Hour),
				now.Add( 5 * time.Hour),
				now.Add( 6 * time.Hour),
				now.Add( 7 * time.Hour),
				now.Add( 8 * time.Hour),
				now.Add( 9 * time.Hour),
				now.Add(10 * time.Hour),
			},
		},



		{
			Slice:    []time.Time{
				now.Add( 1 * time.Hour),
				now.Add( 2 * time.Hour),
				now.Add( 3 * time.Hour),
				now.Add( 4 * time.Hour),
				now.Add( 5 * time.Hour),
				now.Add( 6 * time.Hour),
				now.Add( 7 * time.Hour),
				now.Add( 8 * time.Hour),
				now.Add( 9 * time.Hour),
				now.Add(10 * time.Hour),
				now.Add(11 * time.Hour),
			},
			Expected: []time.Time{
				now.Add( 2 * time.Hour),
				now.Add( 3 * time.Hour),
				now.Add( 4 * time.Hour),
				now.Add( 5 * time.Hour),
				now.Add( 6 * time.Hour),
				now.Add( 7 * time.Hour),
				now.Add( 8 * time.Hour),
				now.Add( 9 * time.Hour),
				now.Add(10 * time.Hour),
				now.Add(11 * time.Hour),
			},
		},



		{
			Slice:    []time.Time{
				now.Add( 1 * time.Hour),
				now.Add( 2 * time.Hour),
				now.Add( 3 * time.Hour),
				now.Add( 4 * time.Hour),
				now.Add( 5 * time.Hour),
				now.Add( 6 * time.Hour),
				now.Add( 7 * time.Hour),
				now.Add( 8 * time.Hour),
				now.Add( 9 * time.Hour),
				now.Add(10 * time.Hour),
				now.Add(11 * time.Hour),
				now.Add(12 * time.Hour),
			},
			Expected: []time.Time{
				now.Add( 3 * time.Hour),
				now.Add( 4 * time.Hour),
				now.Add( 5 * time.Hour),
				now.Add( 6 * time.Hour),
				now.Add( 7 * time.Hour),
				now.Add( 8 * time.Hour),
				now.Add( 9 * time.Hour),
				now.Add(10 * time.Hour),
				now.Add(11 * time.Hour),
				now.Add(12 * time.Hour),
			},
		},
	}


	for testNumber, test := range tests {

		s := Times( test.Slice )

		tail := s.Tail()

		if expected, actual := len(test.Expected), len(tail.(Times)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Expected {
			actual := tail.(Times)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%v\", but actually got \"%v\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}
