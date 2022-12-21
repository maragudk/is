package is

import (
	"errors"
)

type t interface {
	Helper()
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

func Nil[T any](t t, v *T) {
	t.Helper()

	if v != nil {
		t.Fatalf(`Expected nil, but got "%v" (type %T)`, *v, v)
	}
}

func NotNil[T any](t t, v *T) {
	t.Helper()

	if v == nil {
		t.Fatalf(`Expected not nil, but got nil (type %T)`, v)
	}
}

func Error(t t, expected, actual error) {
	t.Helper()

	if !errors.Is(actual, expected) {
		t.Fatalf(`Expected "%v" (type %T), but got "%v" (type %T)`, expected, expected, actual, actual)
	}
}

func NotError(t t, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf(`Expected nil error, but got "%v" (type %T)`, err, err)
	}
}

func Equal[T comparable](t t, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Fatalf(`Expected "%v", but got "%v" (type %T)`, expected, actual, actual)
	}
}

func True(t t, expression bool) {
	t.Helper()

	if !expression {
		t.Fatal("Not true")
	}
}
