package main

import "iter"

type Iterator[T any] interface {
	Next() (T, bool)
}

func Iter[T any](it Iterator[T]) func(func(T) bool) {
	return func(yield func(T) bool) {
		for {
			v, ok := it.Next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}

type StringSliceIterator struct {
	slice []string
	i     int
}

func (it *StringSliceIterator) Next() (string, bool) {
	if it.i >= len(it.slice) {
		return "", false
	}
	v := it.slice[it.i]
	it.i++
	return v, true
}

func TestIter(arr []string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, v := range arr {
			if !yield(v) {
				return
			}
		}
	}
}

func main() {
	arr := []string{"a", "b", "c"}

	for x := range TestIter(arr) {
		println(x)
	}

	arrIter := StringSliceIterator{slice: arr}
	for x := range Iter(&arrIter) {
		println(x)
	}
}
