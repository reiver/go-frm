package frm


func (slice NadaSeries) Tail() Series {
	if len(slice) <= defaultTailLength {
		return slice
	}

	slice2 := slice[len(slice)-defaultTailLength:]

	return slice2
}
