package is_test

import (
	"errors"
	"fmt"
	"io"
	"strings"
	"testing"

	"maragu.dev/is"
)

func TestNil(t *testing.T) {
	t.Run("succeeds when nil", func(t *testing.T) {
		mt := &mockT{}

		var n *int
		is.Nil(mt, n)

		is.True(t, mt.helperCalled)
		is.Equal(t, "", mt.message)
		is.True(t, !mt.failed)
	})

	t.Run("fails when not nil", func(t *testing.T) {
		mt := &mockT{}

		n := 1
		is.Nil(mt, &n)

		is.True(t, mt.helperCalled)
		is.Equal(t, `Expected nil, but got "1" (type *int)`, strings.TrimSpace(mt.message))
		is.True(t, mt.failed)
	})

	t.Run("fails with custom message", func(t *testing.T) {
		mt := &mockT{}

		n := 1
		is.Nil(mt, &n, "custom message")

		is.True(t, mt.helperCalled)
		is.True(t, strings.Contains(mt.message, `Expected nil, but got "1" (type *int)`))
		is.True(t, strings.Contains(mt.message, "custom message"))
		is.True(t, mt.failed)
	})
}

func TestNotNil(t *testing.T) {
	t.Run("succeeds when not nil", func(t *testing.T) {
		mt := &mockT{}

		n := 1
		is.NotNil(mt, &n)

		is.True(t, mt.helperCalled)
		is.Equal(t, "", mt.message)
		is.True(t, !mt.failed)
	})

	t.Run("fails when nil", func(t *testing.T) {
		mt := &mockT{}

		var n *int
		is.NotNil(mt, n)

		is.True(t, mt.helperCalled)
		is.Equal(t, `Expected not nil, but got nil (type *int)`, strings.TrimSpace(mt.message))
		is.True(t, mt.failed)
	})

	t.Run("fails with custom message", func(t *testing.T) {
		mt := &mockT{}

		var n *int
		is.NotNil(mt, n, "custom message")

		is.True(t, mt.helperCalled)
		is.True(t, strings.Contains(mt.message, `Expected not nil, but got nil (type *int)`))
		is.True(t, strings.Contains(mt.message, "custom message"))
		is.True(t, mt.failed)
	})
}

func TestError(t *testing.T) {
	t.Run("succeeds when error", func(t *testing.T) {
		mt := &mockT{}

		err := io.EOF
		is.Error(mt, io.EOF, err)

		is.True(t, mt.helperCalled)
		is.Equal(t, "", mt.message)
		is.True(t, !mt.failed)
	})

	t.Run("fails when nil error", func(t *testing.T) {
		mt := &mockT{}

		var err error
		is.Error(mt, io.EOF, err)

		is.True(t, mt.helperCalled)
		is.Equal(t, `Expected "EOF" (type *errors.errorString), but got "<nil>" (type <nil>)`, strings.TrimSpace(mt.message))
		is.True(t, mt.failed)
	})

	t.Run("fails when different error", func(t *testing.T) {
		mt := &mockT{}

		err := io.ErrNoProgress
		is.Error(mt, io.EOF, err)

		is.True(t, mt.helperCalled)
		is.Equal(t, `Expected "EOF" (type *errors.errorString), but got `+
			`"multiple Read calls return no data or error" (type *errors.errorString)`, strings.TrimSpace(mt.message))
		is.True(t, mt.failed)
	})

	t.Run("fails with custom message", func(t *testing.T) {
		mt := &mockT{}

		err := io.ErrNoProgress
		is.Error(mt, io.EOF, err, "custom message")

		is.True(t, mt.helperCalled)
		is.True(t, strings.Contains(mt.message, `Expected "EOF" (type *errors.errorString), but got `+
			`"multiple Read calls return no data or error" (type *errors.errorString)`))
		is.True(t, strings.Contains(mt.message, "custom message"))
		is.True(t, mt.failed)
	})
}

func TestNotError(t *testing.T) {
	t.Run("succeeds when not error", func(t *testing.T) {
		mt := &mockT{}

		var err error
		is.NotError(mt, err)

		is.True(t, mt.helperCalled)
		is.Equal(t, "", mt.message)
		is.True(t, !mt.failed)
	})

	t.Run("fails when error", func(t *testing.T) {
		mt := &mockT{}

		err := errors.New("bleh")
		is.NotError(mt, err)

		is.True(t, mt.helperCalled)
		is.Equal(t, `Expected nil error, but got "bleh" (type *errors.errorString)`, strings.TrimSpace(mt.message))
		is.True(t, mt.failed)
	})

	t.Run("fails with custom message", func(t *testing.T) {
		mt := &mockT{}

		err := errors.New("bleh")
		is.NotError(mt, err, "custom message")

		is.True(t, mt.helperCalled)
		is.True(t, strings.Contains(mt.message, `Expected nil error, but got "bleh" (type *errors.errorString)`))
		is.True(t, strings.Contains(mt.message, "custom message"))
		is.True(t, mt.failed)
	})
}

func TestEqual(t *testing.T) {
	t.Run("succeeds when equal", func(t *testing.T) {
		mt := &mockT{}

		is.Equal(mt, "123", "123")

		is.True(t, mt.helperCalled)
		is.Equal(t, "", mt.message)
		is.True(t, !mt.failed)
	})

	t.Run("fails when not equal", func(t *testing.T) {
		mt := &mockT{}

		is.Equal(mt, "123", "234")

		is.True(t, mt.helperCalled)
		is.Equal(t, `Expected "123", but got "234" (type string)`, strings.TrimSpace(mt.message))
		is.True(t, mt.failed)
	})

	t.Run("fails with custom message", func(t *testing.T) {
		mt := &mockT{}

		is.Equal(mt, "123", "234", "custom message")

		is.True(t, mt.helperCalled)
		is.True(t, strings.Contains(mt.message, `Expected "123", but got "234" (type string)`))
		is.True(t, strings.Contains(mt.message, "custom message"))
		is.True(t, mt.failed)
	})
}

func TestTrue(t *testing.T) {
	t.Run("succeeds when true", func(t *testing.T) {
		mt := &mockT{}

		is.True(mt, true)

		is.True(t, mt.helperCalled)
		is.Equal(t, "", mt.message)
		is.True(t, !mt.failed)
	})

	t.Run("fails when false", func(t *testing.T) {
		mt := &mockT{}

		is.True(mt, false)

		is.True(t, mt.helperCalled)
		is.Equal(t, `Not true`, strings.TrimSpace(mt.message))
		is.True(t, mt.failed)
	})

	t.Run("fails with custom message", func(t *testing.T) {
		mt := &mockT{}

		is.True(mt, false, "custom message")

		is.True(t, mt.helperCalled)
		is.True(t, strings.Contains(mt.message, `Not true`))
		is.True(t, strings.Contains(mt.message, "custom message"))
		is.True(t, mt.failed)
	})

	t.Run("fails with multiple arguments in message", func(t *testing.T) {
		mt := &mockT{}

		is.True(mt, false, "expected ", 123, " to be valid")

		is.True(t, mt.helperCalled)
		is.True(t, strings.Contains(mt.message, `Not true`))
		is.True(t, strings.Contains(mt.message, "expected 123 to be valid"))
		is.Equal(t, true, mt.failed)
	})
}

type mockT struct {
	helperCalled bool
	message      string
	failed       bool
}

func (m *mockT) Helper() {
	m.helperCalled = true
}

func (m *mockT) Log(args ...any) {
	m.message += fmt.Sprint(args...) + "\n"
}

func (m *mockT) Logf(format string, args ...any) {
	m.message += fmt.Sprintf(format, args...) + "\n"
}

func (m *mockT) FailNow() {
	m.failed = true
}
