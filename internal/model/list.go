package model

type ListResult[T any] struct {
	Count   int
	Results []T
}
