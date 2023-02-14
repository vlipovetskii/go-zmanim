package jdt

// A JWeekday specifies a day of the week (Sunday = 1, ...).
type JWeekday int

const (
	Sunday JWeekday = 1 + iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
