package allocator

import (
	"math/rand"
	"strconv"
)

type MockAllocator struct {
}

func (a MockAllocator) Allocate(fileInfos []FileInfo) (destinations []string, err error) {
	// return random running dates as destinatino according to fileInfos length
	destinations = make([]string, len(fileInfos))
	for i := range fileInfos {
		// append random between 2020-01-01 and 2020-01-05
		destinations[i] = "2020-01-0" + strconv.Itoa(rand.Intn(4)+1)
	}
	return destinations, nil
}

type AnotherMockAllocator struct{}

func (a AnotherMockAllocator) Allocate(fileInfos []FileInfo) (destinations []string, err error) {
	// return random alphabets between a and e as destination according to fileInfos length
	destinations = make([]string, len(fileInfos))
	for i := range fileInfos {
		// append random between 2020-01-01 and 2020-01-05
		destinations[i] = string(rune(rand.Intn(5) + 97))
	}
	return destinations, nil
}

type DebugMockAllocator struct{}

func (a DebugMockAllocator) Allocate(fileInfos []FileInfo) (destinations []string, err error) {
	destinations = make([]string, len(fileInfos))
	for i := range fileInfos {
		destinations[i] = "debug"
	}
	return destinations, nil
}
