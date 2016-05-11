package frm


func (slice Bools) Head() Series {
	if len(slice) <= defaultHeadLength {
		return slice
	}

	slice2 := slice[:defaultHeadLength]

	return slice2
}
