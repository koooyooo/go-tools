package testify

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestRequire(t *testing.T) {
	// assert returns bool value
	if ok := assert.Equal(t, 1, 1); !ok {
		return
	}
	// require doesn't return bool value
	// instead, it finishes the testcase immediately
	require.Equal(t, 1, 1)
}

func TestAssertMany(t *testing.T) {
	a := assert.New(t)
	// custom instance doesn't require first argument of *testing.T
	a.Equal(1, 1)
	a.Equal(2, 2)
	a.Equal(3, 3)
}

type Doer interface {
	DoSomething(string) (string, error)
}

type DoerImpl struct{}

func (i DoerImpl) DoSomething(s string) (string, error) {
	return s + s, nil
}

type MockObject struct {
	mock.Mock
}

// implements mock without any fixed value
func (m MockObject) DoSomething(s string) (string, error) {
	args := m.Called(s)
	return args.String(0), args.Error(1)
}

func TestMock(t *testing.T) {
	// creating testObj and told that
	testObj := new(MockObject)
	// when DoSomething called, assert arg is "hello" and return "test arg", nil
	testObj.On("DoSomething", "hello").Return("test arg", nil)

	// then, set mock to the interface
	var d Doer
	d = testObj
	s, err := d.DoSomething("hello")
	assert.NoError(t, err)
	assert.Equal(t, "test arg", s)
}
