package frm


import (
	"fmt"
)


func mustSeriesLength(s Series) int {
	l, err := seriesLength(s)
	if nil != err {
		panic(err)
	}

	return l
}


func seriesLength(s Series) (int, error) {
	switch x := s.(type) {
	case Bools:
		return len(x), nil
	case Float64s:
		return len(x), nil
	default:
		return 0, fmt.Errorf("Could not figure out length of %T", s)
	}
}
