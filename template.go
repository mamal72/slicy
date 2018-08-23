package slicy

//go:generate genny -in=$GOFILE -out=generated.go gen "Something=byte,int,int8,int16,int32,int64,float32,float64,uint,uint8,uint16,uint32,uint64,string,interface{}"

import (
	"github.com/cheekybits/genny/generic"
)

// Something is obviously a Something
type Something generic.Type

// SliceOfSomething is a struct containing needed data for storing slice items data
type SliceOfSomething struct {
	items []Something
}

// NewFromSomethingSlice creates an instance of SliceOfSomething and returns a reference to it
func NewFromSomethingSlice(items []Something) *SliceOfSomething {
	slicy := &SliceOfSomething{items}
	return slicy
}

// Filter gets a filter function which gets a single Something and only preserve items with true return value from that function
func (s *SliceOfSomething) Filter(filterFunc func(Something) bool) *SliceOfSomething {
	var newItems []Something
	for _, value := range s.items {
		if filterFunc(value) {
			newItems = append(newItems, value)
		}
	}
	return s
}

// Map gets a mapper function and runs that on every single Something items and sets that item with returned value from function
func (s *SliceOfSomething) Map(mapFunc func(Something) Something) *SliceOfSomething {
	for index, value := range s.items {
		s.items[index] = mapFunc(value)
	}
	return s
}

// Shift removes first Something item from the slice
func (s *SliceOfSomething) Shift() *SliceOfSomething {
	s.items = s.items[1:]
	return s
}

// Unshift adds the given Something to the start of the slice
func (s *SliceOfSomething) Unshift(item Something) *SliceOfSomething {
	s.items = append([]Something{item}, s.items...)
	return s
}

// Append adds the given item to the end of the slice
func (s *SliceOfSomething) Append(item Something) *SliceOfSomething {
	s.items = append(s.items, item)
	return s
}

// Concat concats another slice and adds the items in that to the end of current slice
func (s *SliceOfSomething) Concat(items []Something) *SliceOfSomething {
	s.items = append(s.items, items...)
	return s
}

// Push pushes the Something item at the end of the slice
func (s *SliceOfSomething) Push(item Something) *SliceOfSomething {
	s.items = append(s.items, item)
	return s
}

// Pop deletes the Something item from the end of the slice
func (s *SliceOfSomething) Pop() *SliceOfSomething {
	s.items = s.items[:len(s.items)-1]
	return s
}

// Every gets a checker function and runs that on every single Something items and returns true if the function returne true for all of the them
func (s *SliceOfSomething) Every(checkerFunc func(Something) bool) bool {
	for _, value := range s.items {
		if checkerFunc(value) == false {
			return false
		}
	}
	return true
}

// Some gets a checker function and runs that on every single Something items and returns true if the function returne true for one or more of the them
func (s *SliceOfSomething) Some(checkerFunc func(Something) bool) bool {
	for _, value := range s.items {
		if checkerFunc(value) {
			return true
		}
	}
	return false
}

// Includes returns true if the slice contains given item
func (s *SliceOfSomething) Includes(item Something) bool {
	for _, value := range s.items {
		if value == item {
			return true
		}
	}
	return false
}

// Len returns the length of items in current slice
func (s *SliceOfSomething) Len() int {
	return len(s.items)
}

// GetSlice returns a pointer to the final slice of Something
func (s *SliceOfSomething) GetSlice() *[]Something {
	return &s.items
}
