package email

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnATwoWordName(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(
		ReturnATwoWordName("Gabriel Toshio Omiya"),
		"Gabriel Toshio",
		`The result of the string "Gabriel Toshio Omiya" in the ReturnATwoWordName function should be "Gabriel Toshio"`,
	)

	assert.Equal(
		ReturnATwoWordName("Gabriel Toshio"),
		"Gabriel Toshio",
		`The result of the string "Gabriel Toshio" in the ReturnATwoWordName function should be "Gabriel Toshio"`,
	)

	assert.Equal(
		ReturnATwoWordName("Gabriel"),
		"Gabriel",
		`The result of the string "Gabriel" in the ReturnATwoWordName function should be "Gabriel"`,
	)
}

func ExampleReturnATwoWordName_withAThreeWordName() {
	n := ReturnATwoWordName("Gabriel Toshio Omiya")
	fmt.Println(n)
	// Output: "Gabriel Toshio"
}

func ExampleReturnATwoWordName_withATwoWordName() {
	n := ReturnATwoWordName("Gabriel Toshio")
	fmt.Println(n)
	// Output: "Gabriel Toshio"
}

func ExampleReturnATwoWordName_withAOneWordName() {
	n := ReturnATwoWordName("Gabriel")
	fmt.Println(n)
	// Output: "Gabriel"
}
