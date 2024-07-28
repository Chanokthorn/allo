package allocator

import (
	"allo/internal/file_info"
)

type Allocator interface {
	Allocate(fileInfos []file_info.FileInfo) (destinations []string, err error)
}
