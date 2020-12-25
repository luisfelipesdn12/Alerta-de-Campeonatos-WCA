package gspread

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/Iwark/spreadsheet.v2"
)

var spreadData, err = GetSpreadData()

func TestGetSpreadDataReturn(t *testing.T) {
	assert := assert.New(t)

	assert.IsType(
		spreadData,
		spreadsheet.Spreadsheet{},
		"GetSpreadData return should be of type spreadsheet.Spreadsheet",
	)

	if err != nil {
		assert.Empty(
			spreadData,
			"When an error occurs, the spreadsheet.Spreadsheet should be empty",
		)
	} else {
		assert.NotEmpty(
			spreadData,
			"When errors not occurs, the spreadsheet.Spreadsheet should not be empty",
		)
	}
}

func TestGetRecipientsDataReturn(t *testing.T) {
	assert := assert.New(t)

	recipientsData, _ := GetRecipientsData(spreadData)

	assert.IsType(
		recipientsData,
		[]RecipientStruct{},
		"GetRecipientsData return should be of type []RecipientStruct",
	)
}

func TestStripIfNecessary(t *testing.T) {
	assert := assert.New(t)

	fooString := "Foo"
	StripIfNecessary(&fooString)

	assert.Equal(
		fooString,
		"Foo",
		`The result of the string "Foo" in the stripIfNecessary function should be "Foo"`,
	)

	fooString = "  Foo  "
	StripIfNecessary(&fooString)

	assert.Equal(
		fooString,
		"Foo",
		`The result of the string "  Foo  " in the stripIfNecessary function should be "Foo"`,
	)
}

func ExampleStripIfNecessary() {
	fooString := "Foo"
	StripIfNecessary(&fooString)
	fmt.Println(fooString)

	fooString = "  Foo  "
	StripIfNecessary(&fooString)
	fmt.Println(fooString)

	// Output:
	// Foo
	// Foo
}

func TestGetCredentialsDataReturn(t *testing.T) {
	assert := assert.New(t)

	credentialsData, err := GetCredentialsData(spreadData)

	assert.IsType(
		credentialsData,
		CredentialStruct{},
		"GetCredentialsData return should be of type CredentialStruct",
	)

	if err != nil {
		assert.Empty(
			credentialsData,
			"When an error occurs, the gspread.CredentialStruct should be empty",
		)
	} else {
		assert.NotEmpty(
			credentialsData,
			"When errors not occurs, the gspread.CredentialStruct should not be empty",
		)
	}
}
