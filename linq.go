package linq

import "fmt"

type List[T any] struct {
	slice []T
}

// From is constructor of List.
func From[T any](s []T) *List[T] {
	return &List[T]{
		slice: s,
	}
}

// First gets first element of List.
// If element is not found, then it returns error.
func (l *List[T]) First(filter ...func(value T, index int) bool) (T, error) {
	if len(l.slice) == 0 {
		return *new(T), fmt.Errorf("length is 0")
	}

	if len(filter) > 0 {
		for i, t := range l.slice {
			if filter[0](t, i) {
				return t, nil
			}
		}

		return *new(T), fmt.Errorf("not found")
	}

	return l.slice[0], nil
}

// MustFirst gets first element of List.
// If element is not found, then it raises panic.
func (l *List[T]) MustFirst(filter ...func(value T, index int) bool) T {
	first, err := l.First(filter...)
	if err != nil {
		panic(err)
	}

	return first
}

// FirstOrDefault gets first element of List.
// If element is not found, then it returns default value.
func (l *List[T]) FirstOrDefault(filter ...func(value T, index int) bool) T {
	first, err := l.First(filter...)
	if err != nil {
		return *new(T)
	}

	return first
}

// Last gets last element of List.
// If element is not found, then it returns error.
func (l *List[T]) Last(filter ...func(value T, index int) bool) (T, error) {
	if len(l.slice) == 0 {
		return *new(T), fmt.Errorf("length is 0")
	}

	if len(filter) > 0 {
		for i := len(l.slice) - 1; i >= 0; i-- {
			if filter[0](l.slice[i], i) {
				return l.slice[i], nil
			}
		}

		return *new(T), fmt.Errorf("not found")
	}

	return l.slice[len(l.slice)-1], nil
}

// MustLast gets last element of List.
// If element is not found, then it raises panic.
func (l *List[T]) MustLast(filter ...func(value T, index int) bool) T {
	last, err := l.Last(filter...)
	if err != nil {
		panic(err)
	}

	return last
}

// LastOrDefault gets last element of List.
// If element is not found, then it returns default value.
func (l *List[T]) LastOrDefault(filter ...func(value T, index int) bool) T {
	last, err := l.Last(filter...)
	if err != nil {
		return *new(T)
	}

	return last
}

// At returns specific element by index.
// If element is not found, then it returns error.
func (l *List[T]) At(index int) (T, error) {
	if index < 0 || len(l.slice) <= index {
		return *new(T), fmt.Errorf("out of index: %v", index)
	}
	return l.slice[index], nil
}

// MustAt returns specific element by index.
// If element is not found, then it raises panic.
func (l *List[T]) MustAt(index int) T {
	at, err := l.At(index)
	if err != nil {
		panic(err)
	}

	return at
}

// AtOrDefault returns specific element by index.
// If element is not found, then it returns default value.
func (l *List[T]) AtOrDefault(index int) T {
	at, err := l.At(index)
	if err != nil {
		return *new(T)
	}

	return at
}
