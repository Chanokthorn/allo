package file_info

import (
	"allo/internal/signatures"
	"time"
)

type FileInfo struct {
	Signature  signatures.Signature
	Name       string
	CreateDate time.Time
}
