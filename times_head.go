package frm


func (slice Times) Head() Series {
	if len(slice) <= defaultHeadLength {
		return slice
	}

	slice2 := slice[:defaultHeadLength]

	return slice2
}
