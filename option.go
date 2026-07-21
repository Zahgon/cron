package cron

import (
	"time"
)

type Option func(*Cron)

func WithLocation(loc *time.Location) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSeconds() Option { _ = "STUB: not implemented"; return *new(Option) }

func WithParser(p ScheduleParser) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithChain(wrappers ...JobWrapper) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogger(logger Logger) Option { _ = "STUB: not implemented"; return *new(Option) }
