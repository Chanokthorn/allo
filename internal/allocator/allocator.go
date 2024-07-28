package allocator

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

type Allocator interface {
	Allocate(fileInfos []FileInfo) (destinations []string, err error)
}

type FileInfo struct {
	Name       string
	CreateDate time.Time
}

type Processor struct {
	allocators []Allocator
}

func (p Processor) getFileInfo(directory string, fileName string) (FileInfo, error) {
	file, err := os.Open(path.Join(directory, fileName))
	if err != nil {
		log.Panicf("failed to open file: %v", err)
	}
	defer file.Close()

	x, err := exif.Decode(file)
	if err != nil {
		log.Printf("failed to decode exif: %v", err)
		return FileInfo{}, err
	}

	// get created date of file
	createDate, err := x.DateTime()
	if err != nil {
		log.Printf("failed to get created date: %v", err)
		return FileInfo{}, err
	}

	return FileInfo{
		Name:       fileName,
		CreateDate: createDate,
	}, nil
}

func (p Processor) moveFiles(directory string, fileInfos []FileInfo, destinations []string) {
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

func (p *Processor) AddAllocator(allocator Allocator) {
	p.allocators = append(p.allocators, allocator)
}

func (p Processor) Run(directory string) {
	log.Printf("processing directory: %s", directory)
	// list file infos in the current directory
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Panicf("failed to list files: %v", err)
	}

	fileInfos := []FileInfo{}
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
