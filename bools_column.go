package frm


func (slice Bools) Column(n int) Series {
	if 0 != n {
		return NadaSeries( make([]struct{}, len(slice)) )
	}

	return slice
}
