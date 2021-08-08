package go_flags

import (
	"os"
	"testing"

	"github.com/jessevdk/go-flags"
	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	var opts struct {
		Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	}
	remainingArgs, err := flags.ParseArgs(&opts, []string{
		"-vv",
		"arg1",
		"arg2",
		"arg3",
	})
	assert.NoError(t, err)
	assert.Equal(t, []string{"arg1", "arg2", "arg3"}, remainingArgs)
	assert.Equal(t, []bool{true, true}, opts.Verbose)
}

func TestDefault(t *testing.T) {
	var opts struct {
		Default string `short:"d" long:"default" description:"Show default" default:"default-value"`
	}
	_, err := flags.ParseArgs(&opts, []string{})
	assert.NoError(t, err)
	assert.Equal(t, "default-value", opts.Default)
}

func TestEnv(t *testing.T) {
	var opts struct {
		Env string `short:"e" long:"env" description:"possible to override by env" default:"dev" env:"ENV"`
	}
	var err error
	_, err = flags.ParseArgs(&opts, []string{"-e", "stg"})
	assert.NoError(t, err)
	assert.Equal(t, "stg", opts.Env)

	// use default value
	_, err = flags.ParseArgs(&opts, []string{})
	assert.NoError(t, err)
	assert.Equal(t, "dev", opts.Env)

	// override default value
	os.Setenv("ENV", "prod")
	_, err = flags.ParseArgs(&opts, []string{})
	assert.NoError(t, err)
	assert.Equal(t, "prod", opts.Env)
}

func TestAutomaticMarshalling(t *testing.T) {
	var opts struct {
		Offset uint `long:"offset" description:"Offset"`
	}

	_, err := flags.ParseArgs(&opts, []string{
		"--offset=5",
	})
	assert.NoError(t, err)
	assert.Equal(t, uint(5), opts.Offset)
}

func TestFunction(t *testing.T) {
	var opts struct {
		Call func(string) `short:"c" description:"Call phone number"`
	}

	var callLog string
	opts.Call = func(s string) {
		callLog = "call..." + s
	}

	_, err := flags.ParseArgs(&opts, []string{
		"-c", "000-0000-0000",
	})
	assert.NoError(t, err)

	assert.Equal(t, "call...000-0000-0000", callLog)
}

func TestRequired(t *testing.T) {
	var opts struct {
		Name string `short:"n" long:"name" description:"A name" required:"true"`
	}

	_, err := flags.ParseArgs(&opts, []string{
		"-n", "Me",
	})
	assert.NoError(t, err)
	assert.Equal(t, "Me", opts.Name)

	// if required=true argument not found, an error occurs
	_, err = flags.ParseArgs(&opts, []string{})
	assert.Error(t, err, "the required flag `-n, --name' was not specified")
}

func TestChoice(t *testing.T) {
	var opts struct {
		Animal string `long:"animal" choice:"cat" choice:"dog"`
	}

	_, err := flags.ParseArgs(&opts, []string{
		"--animal", "cat",
	})
	assert.NoError(t, err)
	assert.Equal(t, "cat", opts.Animal)

	_, err = flags.ParseArgs(&opts, []string{
		"--animal", "pig",
	})
	assert.Error(t, err, "Invalid value `pig' for option `--animal'. Allowed values are: cat or dog")
}

func TestValueName(t *testing.T) {
	var opts struct {
		File string `short:"f" long:"file" description:"A file" value-name:"FILE"`
	}
	_, err := flags.ParseArgs(&opts, []string{
		"-f", "/tmp/file",
	})
	assert.NoError(t, err)
	assert.Equal(t, "/tmp/file", opts.File)

	// TODO: describe how to use value-name
}

func TestPointer(t *testing.T) {
	var opts struct {
		Ptr *int `short:"p" description:"A pointer to an integer"`
	}
	_, err := flags.ParseArgs(&opts, []string{
		"-p", "3",
	})
	assert.NoError(t, err)
	assert.Equal(t, 3, *opts.Ptr)
}

func TestPointerSlice(t *testing.T) {
	var opts struct {
		StringSlice []*string `long:"ptrslice" description:"A slice of pointers to string"`
	}
	_, err := flags.ParseArgs(&opts, []string{
		"--ptrslice", "hello",
		"--ptrslice", "world",
	})
	assert.NoError(t, err)
	assert.Equal(t, "hello", *opts.StringSlice[0])
	assert.Equal(t, "world", *opts.StringSlice[1])
}

func TestStrIntMap(t *testing.T) {
	var opts struct {
		StrIntMap map[string]int `long:"strintmap" description:"A map from string to int"`
	}
	_, err := flags.ParseArgs(&opts, []string{
		"--strintmap", "a:1",
		"--strintmap", "b:2",
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, opts.StrIntMap["a"])
	assert.Equal(t, 2, opts.StrIntMap["b"])
}
