package cron

import (
	"time"
)

type ParseOption int

const (
	Second ParseOption = 1 << iota
	SecondOptional
	Minute
	Hour
	Dom
	Month
	Dow
	DowOptional
	Descriptor
)

var places = []ParseOption{
	Second,
	Minute,
	Hour,
	Dom,
	Month,
	Dow,
}

var defaults = []string{
	"0",
	"0",
	"0",
	"*",
	"*",
	"*",
}

type Parser struct {
	options ParseOption
}

func NewParser(options ParseOption) Parser { _ = "STUB: not implemented"; return *new(Parser) }

func (p Parser) Parse(spec string) (Schedule, error) {
	_ = "STUB: not implemented"
	return *new(Schedule), nil
}

func normalizeFields(fields []string, options ParseOption) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

var standardParser = NewParser(
	Minute | Hour | Dom | Month | Dow | Descriptor,
)

func ParseStandard(standardSpec string) (Schedule, error) {
	_ = "STUB: not implemented"
	return *new(Schedule), nil
}

func getField(field string, r bounds) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func getRange(expr string, r bounds) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

func parseIntOrName(expr string, names map[string]uint) (uint, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func mustParseInt(expr string) (uint, error) { _ = "STUB: not implemented"; return 0, nil }

func getBits(min, max, step uint) uint64 { _ = "STUB: not implemented"; return 0 }

func all(r bounds) uint64 { _ = "STUB: not implemented"; return 0 }

func parseDescriptor(descriptor string, loc *time.Location) (Schedule, error) {
	_ = "STUB: not implemented"
	return *new(Schedule), nil
}
