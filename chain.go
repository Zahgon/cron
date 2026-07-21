package cron

type JobWrapper func(Job) Job

type Chain struct {
	wrappers []JobWrapper
}

func NewChain(c ...JobWrapper) Chain { _ = "STUB: not implemented"; return *new(Chain) }

func (c Chain) Then(j Job) Job { _ = "STUB: not implemented"; return *new(Job) }

func Recover(logger Logger) JobWrapper { _ = "STUB: not implemented"; return *new(JobWrapper) }

func DelayIfStillRunning(logger Logger) JobWrapper {
	_ = "STUB: not implemented"
	return *new(JobWrapper)
}

func SkipIfStillRunning(logger Logger) JobWrapper {
	_ = "STUB: not implemented"
	return *new(JobWrapper)
}
