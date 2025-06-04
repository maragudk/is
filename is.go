package is

import (
	"errors"
)

type t interface {
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	FailNow()
}

func Nil[T any](t t, v *T, messages ...any) {
	t.Helper()

	if v != nil {
		t.Logf(`Expected nil, but got "%v" (type %T)`, *v, v)
		if len(messages) > 0 {
			t.Log(messages...)
		}
		t.FailNow()
	}
}

func NotNil[T any](t t, v *T, messages ...any) {
	t.Helper()

	if v == nil {
		t.Logf(`Expected not nil, but got nil (type %T)`, v)
		if len(messages) > 0 {
			t.Log(messages...)
		}
		t.FailNow()
	}
}

func Error(t t, expected, actual error, messages ...any) {
	t.Helper()

	if !errors.Is(actual, expected) {
		t.Logf(`Expected "%v" (type %T), but got "%v" (type %T)`, expected, expected, actual, actual)
		if len(messages) > 0 {
			t.Log(messages...)
		}
		t.FailNow()
	}
}

func NotError(t t, err error, messages ...any) {
	t.Helper()

	if err != nil {
		t.Logf(`Expected nil error, but got "%v" (type %T)`, err, err)
		if len(messages) > 0 {
			t.Log(messages...)
		}
		t.FailNow()
	}
}

func Equal[T comparable](t t, expected, actual T, messages ...any) {
	t.Helper()

	if expected != actual {
		t.Logf(`Expected "%v", but got "%v" (type %T)`, expected, actual, actual)
		if len(messages) > 0 {
			t.Log(messages...)
		}
		t.FailNow()
	}
}

func EqualSlice[S ~[]E, E comparable](t t, expected, actual S, messages ...any) {
	t.Helper()

	if len(expected) != len(actual) {
		t.Logf(`Expected slice of length %v, but got %v`, len(expected), len(actual))
		if len(messages) > 0 {
			t.Log(messages...)
		}
		t.FailNow()
	}
	for i := range expected {
		if expected[i] != actual[i] {
			t.Logf(`Expected "%v", but got "%v" (type %T) at index %v`, expected[i], actual[i], actual[i], i)
			if len(messages) > 0 {
				t.Log(messages...)
			}
			t.FailNow()
		}
	}
}

func True(t t, expression bool, messages ...any) {
	t.Helper()

	if !expression {
		t.Log("Not true")
		if len(messages) > 0 {
			t.Log(messages...)
		}
		t.FailNow()
	}
}
