package frm


func (slice Float64s) Head() Series {
	if len(slice) <= defaultHeadLength {
		return slice
	}

	slice2 := slice[:defaultHeadLength]

	return slice2
}
