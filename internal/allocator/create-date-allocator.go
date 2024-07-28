package allocator

import (
	"log"
	"time"
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

type CreateDateAllocator struct {
	opts
}

type CreateDateAllocatorOption func(*opts)

func WithMode(mode Mode) CreateDateAllocatorOption {
	return func(o *opts) {
		o.mode = mode
	}
}

func NewCreateDateAllocator(options ...CreateDateAllocatorOption) CreateDateAllocator {
	opts := opts{
		mode: ModeDay,
	}
	for _, o := range options {
		o(&opts)
	}
	return CreateDateAllocator{opts}
}

func (a CreateDateAllocator) Allocate(fileInfos []FileInfo) (destinations []string, err error) {
	destinations = make([]string, len(fileInfos))
	for i, f := range fileInfos {
		destinations[i] = getDate(a.opts.mode, f.CreateDate)
	}
	return destinations, nil
}
