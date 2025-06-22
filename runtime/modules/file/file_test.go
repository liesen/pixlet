package file_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"tidbyt.dev/pixlet/runtime"
)

var fileSource = `
load("file.star", "file")

def assert(success, message=None):
    if not success:
        fail(message or "assertion failed")

buf = file.read("testfile.txt")

assert(buf == "test", buf)

def main():
	return []
`

func TestFile(t *testing.T) {
	app, err := runtime.NewApplet("file_app.star", []byte(fileSource))
	assert.NoError(t, err)

	screens, err := app.Run(context.Background())
	assert.NoError(t, err)
	assert.NotNil(t, screens)
}
