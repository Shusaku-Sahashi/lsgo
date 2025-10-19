package main

func toPtr[T any](value T) *T {
	return &value
}
