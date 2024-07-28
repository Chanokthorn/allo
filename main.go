package main

import (
	"allo/internal/allocator"
)

func main() {
	processor := allocator.Processor{}
	processor.AddAllocator(allocator.DebugMockAllocator{})
	processor.AddAllocator(allocator.NewCreateDateAllocator(allocator.WithMode(allocator.ModeYear)))
	processor.AddAllocator(allocator.NewCreateDateAllocator(allocator.WithMode(allocator.ModeMonth)))
	processor.AddAllocator(allocator.NewCreateDateAllocator(allocator.WithMode(allocator.ModeDay)))
	processor.Run(".")
}
