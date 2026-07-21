package cron

import (
	"context"
	"sync"
	"time"
)

type Cron struct {
	entries   []*Entry
	chain     Chain
	stop      chan struct{}
	add       chan *Entry
	remove    chan EntryID
	snapshot  chan chan []Entry
	running   bool
	logger    Logger
	runningMu sync.Mutex
	location  *time.Location
	parser    ScheduleParser
	nextID    EntryID
	jobWaiter sync.WaitGroup
}

type ScheduleParser interface {
	Parse(spec string) (Schedule, error)
}

type Job interface {
	Run()
}

type Schedule interface {
	Next(time.Time) time.Time
}

type EntryID int

type Entry struct {
	ID EntryID

	Schedule Schedule

	Next time.Time

	Prev time.Time

	WrappedJob Job

	Job Job
}

func (e Entry) Valid() bool { _ = "STUB: not implemented"; return false }

type byTime []*Entry

func (s byTime) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s byTime) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (s byTime) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func New(opts ...Option) *Cron { _ = "STUB: not implemented"; return nil }

type FuncJob func()

func (f FuncJob) Run() { _ = "STUB: not implemented"; return }

func (c *Cron) AddFunc(spec string, cmd func()) (EntryID, error) {
	_ = "STUB: not implemented"
	return *new(EntryID), nil
}

func (c *Cron) AddJob(spec string, cmd Job) (EntryID, error) {
	_ = "STUB: not implemented"
	return *new(EntryID), nil
}

func (c *Cron) Schedule(schedule Schedule, cmd Job) EntryID {
	_ = "STUB: not implemented"
	return *new(EntryID)
}

func (c *Cron) Entries() []Entry { _ = "STUB: not implemented"; return nil }

func (c *Cron) Location() *time.Location { _ = "STUB: not implemented"; return nil }

func (c *Cron) Entry(id EntryID) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (c *Cron) Remove(id EntryID) { _ = "STUB: not implemented"; return }

func (c *Cron) Start() { _ = "STUB: not implemented"; return }

func (c *Cron) Run() { _ = "STUB: not implemented"; return }

func (c *Cron) run() { _ = "STUB: not implemented"; return }

func (c *Cron) startJob(j Job) { _ = "STUB: not implemented"; return }

func (c *Cron) now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *Cron) Stop() context.Context { _ = "STUB: not implemented"; return *new(context.Context) }

func (c *Cron) entrySnapshot() []Entry { _ = "STUB: not implemented"; return nil }

func (c *Cron) removeEntry(id EntryID) { _ = "STUB: not implemented"; return }
