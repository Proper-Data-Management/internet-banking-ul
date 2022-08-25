package tools

import (
	"strconv"
	"strings"
)

type Number interface {
	int | int64 | float64
}

func ToStr[T Number](a T, precs ...int) (r string) {
	switch p := any(a).(type) {
	case int:
		r = strconv.Itoa(p)
	case int64:
		r = strconv.FormatInt(p, 10)
	case float64:
		var prec int
		if len(precs) == 0 {
			prec = 13
		} else {
			prec = precs[0]
		}
		r = strconv.FormatFloat(p, 'f', prec, 64)
	}

	return
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return ToInt(s) == 1
	}
	return b
}

func RestorePhoneNumber(phone string) (num int) {
	var err error
	phone = strings.TrimSpace(phone)
	if len(phone) < 7 {
		return
	}

	var masked strings.Builder

	switch {
	case strings.HasPrefix(phone, "8"): // 87071234456
		masked.WriteString(phone[1:4])
	case strings.HasPrefix(phone, "7") && len(phone) == 10: // 7071234456
		masked.WriteString(phone[:3])
	case strings.HasPrefix(phone, "7") && len(phone) == 11: // 77071234456
		masked.WriteString(phone[1:4])
	case strings.HasPrefix(phone, "+7"): // +77071234456
		masked.WriteString(phone[2:5])
	}
	masked.WriteString(phone[len(phone)-7:])

	num, err = strconv.Atoi(masked.String())
	if err != nil {
		return
	}

	return num
}

// ToPtr returns a pointer copy of value.
func ToPtr[T any](x T) *T {
	return &x
}

// FromPtr returns the pointer value or empty.
func FromPtr[T any](x *T) T {
	if x == nil {
		return Empty[T]()
	}

	return *x
}

// ToSlicePtr returns a slice of pointer copy of value.
func ToSlicePtr[T any](collection []T) []*T {
	return Map(collection, func(x T, _ int) *T {
		return &x
	})
}

// ToAnySlice returns a slice with all elements mapped to `any` type
func ToAnySlice[T any](collection []T) []any {
	result := make([]any, len(collection))
	for i, item := range collection {
		result[i] = item
	}
	return result
}

// FromAnySlice returns an `any` slice with all elements mapped to a type.
// Returns false in case of type conversion failure.
func FromAnySlice[T any](in []any) (out []T, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			out = []T{}
			ok = false
		}
	}()

	result := make([]T, len(in))
	for i, item := range in {
		result[i] = item.(T)
	}
	return result, true
}

// Empty returns an empty value.
func Empty[T any]() T {
	var zero T
	return zero
}

// IsEmpty returns true if argument is a zero value.
func IsEmpty[T comparable](v T) bool {
	var zero T
	return zero == v
}

// IsNotEmpty returns true if argument is not a zero value.
func IsNotEmpty[T comparable](v T) bool {
	var zero T
	return zero != v
}

// Coalesce returns the first non-empty arguments. Arguments must be comparable.
func Coalesce[T comparable](v ...T) (result T, ok bool) {
	for _, e := range v {
		if e != result {
			result = e
			ok = true
			return
		}
	}

	return
}
