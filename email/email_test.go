package email

import (
	"fmt"
	"testing"

	"github.com/luisfelipesdn12/Alerta-de-Campeonatos-WCA/gspread"
	"github.com/stretchr/testify/assert"
	"gopkg.in/Iwark/spreadsheet.v2"
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

func ExampleReturnATwoWordName() {
	n := ReturnATwoWordName("Gabriel Toshio Omiya")
	fmt.Println("With three words:", n)

	n = ReturnATwoWordName("Gabriel Toshio")
	fmt.Println("With two words:", n)

	n = ReturnATwoWordName("Gabriel")
	fmt.Println("With one word:", n)

	// Output:
	// With three words: Gabriel Toshio
	// With two words: Gabriel Toshio
	// With one word: Gabriel
}

func TestSendEmail(t *testing.T) {
	mockRecipient := gspread.RecipientStruct{
		Name:     spreadsheet.Cell{Value: "Lorem Ipsum"},
		Email:    spreadsheet.Cell{Value: "example@mail.com"},
		City:     spreadsheet.Cell{Value: "Washington DC"},
		Language: spreadsheet.Cell{Value: "English"},
	}

	mockCredentials := gspread.CredentialStruct{
		Email:    "mock",
		Password: "mock",
	}

	_ = SendEmail(mockRecipient, mockCredentials)
}
