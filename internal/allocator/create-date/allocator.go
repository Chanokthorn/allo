package create_date

import (
	"log"
	"time"

	"allo/internal/file_info"
)

func getDate(mode Mode, t time.Time) string {
	switch mode {
	case ModeYear:
		return t.Format("2006")
	case ModeMonth:
		return t.Format("2006-01")
	case ModeDay:
		return t.Format("2006-01-02")
	default:
		log.Panicf("invalid create date allocator mode: %v", mode)
		return ""
	}
}

type Mode int

const (
	ModeYear Mode = iota
	ModeMonth
	ModeDay
)

type opts struct {
	mode Mode
}

type Allocator struct {
	opts
}

type Option func(*opts)

func WithMode(mode Mode) Option {
	return func(o *opts) {
		o.mode = mode
	}
}

func New(options ...Option) Allocator {
	opts := opts{
		mode: ModeDay,
	}
	for _, o := range options {
		o(&opts)
	}
	return Allocator{opts}
}

func (a Allocator) Allocate(fileInfos []file_info.FileInfo) (destinations []string, err error) {
	destinations = make([]string, len(fileInfos))
	for i, f := range fileInfos {
		destinations[i] = getDate(a.opts.mode, f.CreateDate)
	}
	return destinations, nil
}
