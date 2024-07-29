package processor

import (
	"errors"
	"io"
	"log"
	"os"
	"path"
	"time"

	"allo/internal/allocator"
	"allo/internal/file_info"
	"allo/internal/signatures"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

type Processor struct {
	allocators []allocator.Allocator
}

func New() Processor {
	return Processor{}
}

func (p Processor) getFileInfo(directory string, fileName string) (file_info.FileInfo, error) {
	log.Print("getting file info: ", path.Join(directory, fileName))
	file, err := os.Open(path.Join(directory, fileName))
	if err != nil {
		log.Panicf("failed to open file: %v", err)
	}
	defer file.Close()

	// Read the first 8 bytes (enough for most file signatures)
	sig := signatures.Signature(make([]byte, 8))
	_, err = file.Read(sig)
	if err != nil {
		log.Printf("Failed to read file: %v\n", err)
		return file_info.FileInfo{}, err
	}

	// only process accepted signatures
	if !signatures.IsAcceptedSignature(sig) {
		log.Printf("File signature not accepted: %v\n", sig)
		return file_info.FileInfo{}, errors.New("file signature not accepted")
	}

	// Reset the read pointer to the start of the file before decoding EXIF
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Printf("failed to seek file: %v", err)
		return file_info.FileInfo{}, err
	}

	var createDate time.Time

	if x, err := exif.Decode(file); err == nil {
		// get created date of file
		createDate, err = x.DateTime()
		if err != nil {
			log.Printf("failed to get created date: %v", err)
			return file_info.FileInfo{}, err
		}
	} else {
		log.Printf("failed to decode exif: %s", err)
		// get created date of file
		fileInfo, err := file.Stat()
		if err != nil {
			log.Printf("failed to get file info: %v", err)
			return file_info.FileInfo{}, err
		}
		createDate = fileInfo.ModTime()
	}

	return file_info.FileInfo{
		Signature:  sig,
		Name:       fileName,
		CreateDate: createDate,
	}, nil
}

func (p Processor) moveFiles(directory string, fileInfos []file_info.FileInfo, destinations []string) {
	if len(fileInfos) != len(destinations) {
		log.Panicf("fileInfos and destinations length mismatch")
	}

	destinationMap := make(map[string]bool)
	for i := range fileInfos {
		if _, ok := destinationMap[destinations[i]]; !ok {
			err := os.MkdirAll(path.Join(directory, destinations[i]), 0755)
			if err != nil {
				log.Panicf("failed to create directory: %v", err)
			}
			destinationMap[destinations[i]] = true
		}
		log.Print("moving file: ", path.Join(directory, fileInfos[i].Name), " to: ", path.Join(directory, destinations[i], fileInfos[i].Name))
		os.Rename(path.Join(directory, fileInfos[i].Name), path.Join(directory, destinations[i], fileInfos[i].Name))
	}
}

func (p *Processor) AddAllocator(allocator allocator.Allocator) {
	p.allocators = append(p.allocators, allocator)
}

func (p Processor) Run(directory string) {
	exif.RegisterParsers(mknote.All...)
	log.Printf("processing directory: %s", directory)
	// list file infos in the current directory
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Panicf("failed to list files: %v", err)
	}

	fileInfos := []file_info.FileInfo{}
	for _, f := range files {
		if f.Type().IsDir() {
			continue
		}
		info, err := p.getFileInfo(directory, f.Name())
		if err != nil {
			log.Printf("failed to get file info: %v", err)
			continue
		}
		fileInfos = append(fileInfos, info)
	}

	// run the first allocator in list
	destinations, err := p.allocators[0].Allocate(fileInfos)
	if err != nil {
		log.Panicf("failed to allocate: %v", err)
	}

	// move files to the allocated directories
	p.moveFiles(directory, fileInfos, destinations)

	// new processors for each allocated directory, remove the latest allocator
	if len(p.allocators) == 1 {
		return
	}
	newProcessor := Processor{
		allocators: p.allocators[1:len(p.allocators)],
	}

	// run the new processors
	for _, d := range destinations {
		newProcessor.Run(path.Join(directory, d))
	}
}
