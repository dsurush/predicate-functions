package main

import "testing"

func TestIndex(t *testing.T) {
	got := Index([]int{1, -2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	want := 1
	if got != want{
		t.Error("Index() got ", got, " want ", want)
	}

	got = Index([]int{1, 2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	want = -1
	if got != want{
		t.Error("Index() got ", got, " want ", want)
	}
}

func TestFind(t *testing.T) {

	want := -2
	got := Find([]int{1, -2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	if got != want{
		t.Error("Find() got ", got, " want ", want)
	}

	defer func() {
		err := recover()
		if err == nil {
			t.Error("need panic")
		}
	}()

	Find([]int{1, 2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
}

func TestAny(t *testing.T) {
	got := Any([]int{1, 2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	want := false
	if got != want{
		t.Error("Any() got", got, "want", want)
	}

	got = Any([]int{1, -2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	want = true
	if got != want{
		t.Error("Any() got", got, "want", want)
	}
}

func TestNone(t *testing.T) {
	got := None([]int{1, 2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	want := true
	if got != want{
		t.Error("None() got", got, "want", want)
	}

	got = None([]int{1, -2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	want = false
	if got != want{
		t.Error("None() got", got, "want", want)
	}
}

func TestAll(t *testing.T) {
	got := All([]int{1, 2, 3, 4, 10}, func(a int) bool {
		return a < 0
	})
	want := false
	if got != want{
		t.Error("All() got", got, "want", want)
	}

	got = All([]int{-1, -2, -3, -4, -10}, func(a int) bool {
		return a < 0
	})
	want = true
	if got != want{
		t.Error("All() got", got, "want", want)
	}
}