package linq

import (
	"fmt"
)

type List[T comparable] struct {
	slice []T
}

// From is constructor of List.
func From[T comparable](s []T) *List[T] {
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

// Skip returns elements after the specified index.
func (l *List[T]) Skip(index int) *List[T] {
	if index < 0 {
		index = 0
	}
	if index >= len(l.slice) {
		index = len(l.slice)
	}
	return From(l.slice[index:])
}

// SkipWhile returns elements after the specified condition.
func (l *List[T]) SkipWhile(f func(value T, index int) bool) *List[T] {
	for i, t := range l.slice {
		if f(t, i) {
			return From(l.slice[i:])
		}
	}
	return From(l.slice[len(l.slice):])
}

// Take returns elements up to the specified index.
func (l *List[T]) Take(count int) *List[T] {
	if count < 0 {
		count = 0
	}
	if count >= len(l.slice) {
		count = len(l.slice)
	}
	return From(l.slice[:count])
}

// TakeWhile returns elements up to the specified condition.
func (l *List[T]) TakeWhile(f func(value T, index int) bool) *List[T] {
	for i, t := range l.slice {
		if !f(t, i) {
			return From(l.slice[:i])
		}
	}
	return l
}

// DefaultIfEmpty returns default value if list is empty.
func (l *List[T]) DefaultIfEmpty(defaultT ...T) *List[T] {
	if len(l.slice) > 0 {
		return l
	}

	if len(defaultT) > 0 {
		return From([]T{defaultT[0]})
	}

	return From([]T{*new(T)})
}

// Where returns condition matched elements
func (l *List[T]) Where(f func(value T, index int) bool) *List[T] {
	s := make([]T, 0, len(l.slice))
	for i, t := range l.slice {
		if f(t, i) {
			s = append(s, t)
		}
	}

	return From(s)
}

// All returns true if all elements are matched
func (l *List[T]) All(f func(value T, index int) bool) bool {
	for i, t := range l.slice {
		if !f(t, i) {
			return false
		}
	}

	return true
}

// Any returns true if there is matched element
func (l *List[T]) Any(f ...func(value T, index int) bool) bool {
	if len(f) == 0 {
		return len(l.slice) > 0
	}

	for i, t := range l.slice {
		if f[0](t, i) {
			return true
		}
	}

	return false
}

// Contains returns true if there is matched element
func (l *List[T]) Contains(value T) bool {
	for _, t := range l.slice {
		if t == value {
			return true
		}
	}

	return false
}

// SequenceEqual return true if all element of two list are the same values
func (l *List[T]) SequenceEqual(other *List[T]) bool {
	if len(l.slice) != len(other.slice) {
		return false
	}
	for i, t := range l.slice {
		if other.slice[i] != t {
			return false
		}
	}

	return true
}

// Count returns number of element
func (l *List[T]) Count(f ...func(value T, index int) bool) int {
	if len(f) == 0 {
		return len(l.slice)
	}

	return len(l.Where(f[0]).slice)
}

// Max returns maximum element of list
func (l *List[T]) Max(f func(value T, index int) float64) T {
	if len(l.slice) == 0 {
		return *new(T)
	}

	maxV := f(l.slice[0], 0)
	max := l.slice[0]
	for i, t := range l.slice[1:] {
		v := f(t, i)
		if maxV < v {
			maxV = v
			max = t
		}
	}

	return max
}

// Min returns minimum element of list
func (l *List[T]) Min(f func(value T, index int) float64) T {
	if len(l.slice) == 0 {
		return *new(T)
	}

	minV := f(l.slice[0], 0)
	min := l.slice[0]
	for i, t := range l.slice[1:] {
		v := f(t, i)
		if minV > v {
			minV = v
			min = t
		}
	}

	return min
}

// Average returns average of list
func (l *List[T]) Average(f func(value T, index int) float64) float64 {
	if len(l.slice) == 0 {
		return 0
	}

	sum := 0.0
	for i, t := range l.slice {
		sum += f(t, i)
	}

	return sum / float64(len(l.slice))
}

// Sum returns sum of elements
func (l *List[T]) Sum(f func(value T, index int) float64) float64 {
	sum := 0.0
	for i, t := range l.slice {
		sum += f(t, i)
	}

	return sum
}

// ToSlice returns slice of elements
func (l *List[T]) ToSlice() []T {
	return l.slice
}

// Reverse returns reversed list
func (l *List[T]) Reverse() *List[T] {
	s := make([]T, len(l.slice))
	for i := 0; i < len(l.slice); i++ {
		s[i] = l.slice[len(l.slice)-i-1]
	}
	return From(s)
}

// Distinct returns list excluding duplicate elements
func (l *List[T]) Distinct() *List[T] {
	m := make(map[T]bool)
	s := make([]T, 0, len(l.slice))
	for _, t := range l.slice {
		if _, ok := m[t]; !ok {
			m[t] = true
			s = append(s, t)
		}
	}
	return From(s)
}
