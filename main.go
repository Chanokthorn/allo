package main

import (
	create_date "allo/internal/allocator/create-date"
	"allo/internal/allocator/mock"
	"allo/internal/processor"
)

func main() {
	processor := processor.Processor{}
	processor.AddAllocator(mock.DebugAllocator{})
	processor.AddAllocator(create_date.New(create_date.WithMode(create_date.ModeYear)))
	processor.AddAllocator(create_date.New(create_date.WithMode(create_date.ModeMonth)))
	processor.AddAllocator(create_date.New(create_date.WithMode(create_date.ModeDay)))
	processor.Run(".")
}
