package jdt

/*
JDateTime is an internal structure to aggregate JDate, MoladTime
*/
type JDateTime struct {
	D JDate
	T MoladTime
}

/*
NewJDateTime
E.g. NewJDateTime(NewJDate2(gDate), moladTime)
*/
func NewJDateTime(jDate JDate, moladTime MoladTime) JDateTime {
	return JDateTime{D: jDate, T: moladTime}
}

func (t JDateTime) Validate() {
	t.D.Validate()
	t.T.Validate()
}
