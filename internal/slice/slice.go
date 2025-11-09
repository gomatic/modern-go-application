// Package slice provides functions for converting slices.
package slice

func StringSliceConverter[F string, T ~string](from []F) []T {
	return Convert(from, StringConverter[F, T])
}

func IntSliceConverter[F int, T ~int](from []F) []T {
	return Convert(from, IntConverter[F, T])
}

func Int64SliceConverter[F int64, T ~int64](from []F) []T {
	return Convert(from, Int64Converter[F, T])
}

// ConvertFunc is a function that converts from type F to type T.
type ConvertFunc[F, T any] func(F) T

func StringConverter[F string, T ~string](from F) T { return T(from) }
func IntConverter[F int, T ~int](from F) T          { return T(from) }
func Int64Converter[F int64, T ~int64](from F) T    { return T(from) }

// Convert transforms a slice of type []F to []T using the provided conversion function.
func Convert[F, T any](from []F, convert ConvertFunc[F, T]) []T {
	if from == nil {
		return nil
	}
	out := make([]T, len(from))
	for i, v := range from {
		out[i] = convert(v)
	}
	return out
}
