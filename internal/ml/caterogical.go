package ml

import (
	"runtime"
	"sync"
)

// CategoricalToNumericConverter converts passed struct
// to categorical []float64 format where
// each field is one element of the returned
// slice.
type CategoricalToNumericConverter[T any] struct {
	// dataSlice is a slice of other entities of the same type
	// used to convert categorical data to numeric one
	dataSlice []T

	// convertRoutinesCount represents how many routines will be working
	// in order to calculate given tree.
	convertRoutinesCount int

	mu sync.Mutex
}

// NewCategoricalConverter creates a new instance of the categorical converter
func NewCategoricalConverter[T any](dataSlice []T) *CategoricalToNumericConverter[T] {
	return &CategoricalToNumericConverter[T]{
		dataSlice:            dataSlice,
		convertRoutinesCount: runtime.NumCPU(), // basically there is no need have more workers than cores.
	}
}

// Convert converts given struct fields into []float64 and returns it.
// It converts all the fields into numeric format.
func (c *CategoricalToNumericConverter[T]) Convert(obj T) []float64 {
	c.mu.Lock()
	defer c.mu.Unlock()

	var wg sync.WaitGroup
	for i := range c.convertRoutinesCount {

	}

}
