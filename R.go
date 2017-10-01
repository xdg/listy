// DO NOT EDIT!  This file was generated via `go generate`

// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy

// R wraps a slice of rune
type R []rune

// Contains checks if a value is in the list
func (xs R) Contains(v rune) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Filter returns a new list of elements matching a predicate
func (xs R) Filter(f func(rune) bool) R {
	ys := make(R, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}

// Head returns the first value in the list.  Panics if the list is empty.
func (xs R) Head() rune {
	return xs[0]
}

// Init returns a new list with all values except the last.  Panics if the
// list is empty.
func (xs R) Init() R {
	ys := make(R, len(xs)-1)
	copy(ys, xs[:len(xs)-1])
	return ys
}

// List returns the last value in the list.  Panics if the list is empty.
func (xs R) Last() rune {
	return xs[len(xs)-1]
}

// Len returns the length of the list.
func (xs R) Len() int {
	return len(xs)
}

// Less reports whether the element with index i should sort before the
// element with index j.
func (xs R) Less(i, j int) bool {
	return xs[i] < xs[j]
}

// Map returns a new list with every element transformed by a function
func (xs R) Map(f func(rune) rune) R {
	ys := make(R, 0, len(xs))
	for _, x := range xs {
		ys = append(ys, f(x))
	}
	return ys
}

// Swap swaps the elements with indexes i and j and returns the original list.
func (xs R) Swap(i, j int) R {
	xs[i], xs[j] = xs[j], xs[i]
	return xs
}

// Tail returns a new list with all values except the head.  Panics if the
// list is empty.
func (xs R) Tail() R {
	ys := make(R, len(xs)-1)
	copy(ys, xs[1:len(xs)])
	return ys
}

// Unbox returns the list as a native slice
func (xs R) Unbox() []rune {
	return []rune(xs)
}
