package frm


func (slice Times) Column(n int) Series {
	if 0 != n {
		return NadaSeries( make([]struct{}, len(slice)) )
	}

	return slice
}
