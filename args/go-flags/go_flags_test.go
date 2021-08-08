package go_flags

import (
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
)

type opts struct {
	Verbose []bool       `short:"v" long:"verbose" description:"Show verbose debug information"`
	Offset  uint         `long:"offset" description:"Offset"`
	Call    func(string) `short:"c" description:"Call phone number"`
	Name    string       `short:"n" long:"name" description:"A name" required:"true"`
}

func TestDo(t *testing.T) {
	args := []string{
		"-vv",
		"--offset=5",
		"-n", "Me",
		"-c", "000-0000-0000",
		"arg1",
		"arg2",
		"arg3",
	}

	var o opts
	var callLog string
	o.Call = func(s string) {
		callLog = "call..." + s
	}

	remainingArgs, err := flags.ParseArgs(&o, args)
	assert.NoError(t, err)

	assert.Equal(t, []string{"arg1", "arg2", "arg3"}, remainingArgs)

	assert.Equal(t, []bool{true, true}, o.Verbose)
	assert.Equal(t, uint(5), o.Offset)
	assert.Equal(t, "call...000-0000-0000", callLog)
	assert.Equal(t, "Me", o.Name)
}
