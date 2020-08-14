package wca

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpcomingCompetitionsReturn(t *testing.T) {
	assert := assert.New(t)

	result, _ := UpcomingCopetitions("New Jersey")

	assert.IsType(
		result,
		31415,
		"The return of UpcomingCopetitions should be an interger",
	)
}
