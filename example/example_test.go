package example

import (
	"testing"

	"github.com/gustavosbarreto/structsnapshot"
	"github.com/stretchr/testify/assert"
)

func TestUserStruct(t *testing.T) {
	// Take snapshot of current `User` struct
	current, err := structsnapshot.TakeSnapshot(User{})
	assert.NoError(t, err)

	// Load snapshot for `User` struct from filesystem
	expected, err := structsnapshot.LoadSnapshot(User{})
	assert.NoError(t, err)

	assert.Equal(t, expected, current, "User struct has changed! Did you forget to update it? To update run `go generate`")
}
