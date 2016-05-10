package frm


type Series interface {
	Column(int) Series
	Head() Series
	Len() int
	Tail() Series
}
