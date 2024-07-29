package raw_jpeg

import (
	"allo/internal/file_info"
	"allo/internal/signatures"
)

type allocator struct {
}

func New() allocator {
	return allocator{}
}

func (a allocator) Allocate(fileInfos []file_info.FileInfo) (destinations []string, err error) {
	destinations = make([]string, len(fileInfos))
	for i, f := range fileInfos {
		switch {
		case signatures.IsJPEG(f.Signature):
			destinations[i] = "jpeg"
		case signatures.IsRaw(f.Signature):
			destinations[i] = "raw"
		default:
			destinations[i] = "other"
		}
	}
	return destinations, nil
}
