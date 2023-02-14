package jdt

/*
MoladTime is an internal structure to track the molad time
*/
type MoladTime struct {
	// the internal count of <em>molad</em> hours
	Hours MoladHours
	// the internal count of <em>molad</em> minutes
	Minutes MoladMinutes
	// the internal count of <em>molad</em> <em>chalakim</em>
	Chalakim MoladChalakim
}

func NewMoladTime(hours MoladHours, minutes MoladMinutes, chalakim MoladChalakim) MoladTime {
	return MoladTime{Hours: hours, Minutes: minutes, Chalakim: chalakim}
}

func NewMoladTime0() MoladTime {
	return NewMoladTime(0, 0, 0)
}

func (t MoladTime) Validate() {
	t.Hours.Validate()
	t.Minutes.Validate()
	t.Chalakim.Validate()
}
