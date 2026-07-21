package cron

import "time"

type SpecSchedule struct {
	Second, Minute, Hour, Dom, Month, Dow uint64

	Location *time.Location
}

type bounds struct {
	min, max uint
	names    map[string]uint
}

var (
	seconds = bounds{0, 59, nil}
	minutes = bounds{0, 59, nil}
	hours   = bounds{0, 23, nil}
	dom     = bounds{1, 31, nil}
	months  = bounds{1, 12, map[string]uint{
		"jan": 1,
		"feb": 2,
		"mar": 3,
		"apr": 4,
		"may": 5,
		"jun": 6,
		"jul": 7,
		"aug": 8,
		"sep": 9,
		"oct": 10,
		"nov": 11,
		"dec": 12,
	}}
	dow = bounds{0, 6, map[string]uint{
		"sun": 0,
		"mon": 1,
		"tue": 2,
		"wed": 3,
		"thu": 4,
		"fri": 5,
		"sat": 6,
	}}
)

const (
	starBit = 1 << 63
)

func (s *SpecSchedule) Next(t time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func dayMatches(s *SpecSchedule, t time.Time) bool { _ = "STUB: not implemented"; return false }
