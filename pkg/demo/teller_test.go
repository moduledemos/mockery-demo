package demo

import (
	"testing"

	demoMock "github.com/moduledemos/mockery-demo/mocks/pkg/demo"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTeller_Tell(t *testing.T) {
	teller := Teller{}
	story := new(demoMock.Story)
	noChapters := 20
	story.On("ReadChapter", mock.Anything).Return(
		[]byte(":)"), nil).Times(noChapters)
	for i := 0; i < noChapters; i++ {
		story.On("HasChapter", i).Return(true)
	}
	story.On("HasChapter", noChapters).Return(false)
	err := teller.Tell(story)
	assert.NoError(t, err)

	story.AssertExpectations(t)
}
