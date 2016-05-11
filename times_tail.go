package frm


func (slice Times) Tail() Series {
	if len(slice) <= defaultTailLength {
		return slice
	}

	slice2 := slice[len(slice)-defaultTailLength:]

	return slice2
}
