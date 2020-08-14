package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnATwoWordName(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		ReturnATwoWordName("Gabriel Toshio Omiya"),
		"Gabriel Toshio",
		`The result of the string "Gabriel Toshio Omiya" in the stripIfNecessary function should be "Gabriel Toshio"`,
	)

	assert.Equal(
		ReturnATwoWordName("Gabriel Toshio"),
		"Gabriel Toshio",
		`The result of the string "Gabriel Toshio" in the stripIfNecessary function should be "Gabriel Toshio"`,
	)

	assert.Equal(
		ReturnATwoWordName("Gabriel"),
		"Gabriel",
		`The result of the string "Gabriel" in the stripIfNecessary function should be "Gabriel"`,
	)
}
