// Copyright 2017 by David A. Golden. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package listy_test

import (
	"encoding/json"
	"testing"

	"github.com/xdg/listy"
	"github.com/xdg/testy"
)

func getTestData(is *testy.T, v interface{}) {
	err := json.Unmarshal(testdata["ints"], v)
	if err != nil {
		is.Fatalf("Error unmarshaling ints.json: %s", err)
	}
}

func TestListI8Box(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Box struct {
			Input []int8
			Unbox []int8
		}
	}
	getTestData(is, &data)

	is.Equal(listy.I8(data.Box.Input).Unbox(), data.Box.Unbox)
}

func TestListI8Contains(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Contains struct {
			Input         []int8
			ContainsTrue  []int8
			ContainsFalse []int8
		}
	}
	getTestData(is, &data)

	xs := listy.I8(data.Contains.Input)

	for _, x := range data.Contains.ContainsTrue {
		is.True(xs.Contains(x))
	}
	for _, x := range data.Contains.ContainsFalse {
		is.False(xs.Contains(x))
	}
}

func TestListI8Filter(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Filter struct {
			Input    []int8
			Lessthan int8
			Filtered []int8
		}
	}
	getTestData(is, &data)

	xs := listy.I8(data.Filter.Input)

	ys := xs.Filter(func(v int8) bool { return v < data.Filter.Lessthan })

	is.Equal(ys.Unbox(), data.Filter.Filtered)
}

func TestListI8Map(t *testing.T) {
	is := testy.New(t)
	defer func() { t.Logf(is.Done()) }()

	var data struct {
		Map struct {
			Input  []int8
			Add    int8
			Mapped []int8
		}
	}
	getTestData(is, &data)

	xs := listy.I8(data.Map.Input)
	ys := xs.Map(func(v int8) int8 { return v + data.Map.Add })

	is.Equal(ys.Unbox(), data.Map.Mapped)
}
