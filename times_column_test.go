package frm


import (
	"time"

	"testing"
)


func TestTimesColumn0(t *testing.T) {

	now := time.Now()


	tests := []struct{
		Slice []time.Time
	}{
		{
			Slice: nil,
		},



		{
			Slice: []time.Time{},
		},



		{
			Slice: []time.Time{ now.Add(1 * time.Hour), },
		},



		{
			Slice: []time.Time{ now.Add(1 * time.Hour), now.Add(2 * time.Hour), },
		},



		{
			Slice: []time.Time{
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
			Slice: []time.Time{
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
			Slice: []time.Time{
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
		},



		{
			Slice: []time.Time{
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
		},
	}


	for testNumber, test := range tests {

		s := Times( test.Slice )

		c := s.Column(0)

		if expected, actual := len(test.Slice), len(c.(Times)); expected != actual {
			t.Errorf("For test #%d, expected %d, but actually got %d.", testNumber, expected, actual)
			continue
		}

		for seriesIndex, expected := range test.Slice {
			actual := c.(Times)[seriesIndex]

			if expected != actual {
				t.Errorf("For test #%d and series item #%d, expected \"%f\", but actually got \"%f\".", testNumber, seriesIndex, expected, actual)
				continue
			}
		}
	}
}


func TestTimesColumnNot0(t *testing.T) {

	now := time.Now()


	tests := []struct{
		Slice []time.Time
	}{
		{
			Slice: nil,
		},



		{
			Slice: []time.Time{},
		},



		{
			Slice: []time.Time{ now.Add(1 * time.Hour), },
		},



		{
			Slice: []time.Time{ now.Add(1 * time.Hour), now.Add(2 * time.Hour), },
		},



		{
			Slice: []time.Time{
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
			Slice: []time.Time{
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
			Slice: []time.Time{
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
		},



		{
			Slice: []time.Time{
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
		},
	}


	for testNumber, test := range tests {

		for columnNumber := -11; columnNumber < 12; columnNumber++ {

			if 0 == columnNumber {
				continue
			}

			s := Times( test.Slice )

			c := s.Column(columnNumber)

			if _, ok := c.(NadaSeries); !ok {
				t.Errorf("For test #%d and column #%d, expected series to be NadaSeries, but actually was %T.", testNumber, columnNumber, c)
				continue
			}

			if expected, actual := len(test.Slice), len(c.(NadaSeries)); expected != actual {
				t.Errorf("For test #%d and column #%d, expected \"%v\", but actually got \"%v\".", testNumber, columnNumber, expected, actual)
				continue
			}
		}
	}
}
