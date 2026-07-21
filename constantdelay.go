package cron

import "time"

type ConstantDelaySchedule struct {
	Delay time.Duration
}

func Every(duration time.Duration) ConstantDelaySchedule {
	_ = "STUB: not implemented"
	return *new(ConstantDelaySchedule)
}

func (schedule ConstantDelaySchedule) Next(t time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
