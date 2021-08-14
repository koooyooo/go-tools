package testify

import (
	"testing"

	"github.com/stretchr/testify/suite"

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
	Do()
	DoSomething(string) (string, error)
}

type MockObject struct {
	mock.Mock
}

func (m *MockObject) Do() {
	// record calls
	m.Called()
}

// be aware! the receiver should be pointer, unless method call won't be counted
func (m *MockObject) DoSomething(s string) (string, error) {
	// record calls
	args := m.Called(s)
	// setting return values
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

	// assert everything specified with On and Return was called as expected.
	testObj.AssertExpectations(t)
}

func TestCallsPlane(t *testing.T) {
	// creating testObj and told that
	testObj := new(MockObject)
	testObj.On("Do").Return()

	// then, set mock to the interface
	var d Doer
	d = testObj
	d.Do()
	d.Do()
	d.Do()

	// on AssertCalled, argument is required
	testObj.AssertCalled(t, "Do")
	testObj.AssertNumberOfCalls(t, "Do", 3)
}

func TestCallsWithArgsAndReturns(t *testing.T) {
	// creating testObj and told that
	testObj := new(MockObject)
	testObj.On("DoSomething", "hello").Return("test arg", nil)

	// then, set mock to the interface
	var d Doer
	d = testObj
	var err error
	_, err = d.DoSomething("hello")
	_, err = d.DoSomething("hello")
	assert.NoError(t, err)

	// on AssertCalled, argument is required
	testObj.AssertCalled(t, "DoSomething", "hello")
	testObj.AssertNumberOfCalls(t, "DoSomething", 2)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(MySuite))
}

type MySuite struct {
	suite.Suite
	initialState int
}

func (s *MySuite) SetupTest() {
	s.initialState = 10
}

func (s *MySuite) TestInitialized() {
	s.Equal(10, s.initialState)
}
