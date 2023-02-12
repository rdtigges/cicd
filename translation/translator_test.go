package translation_test

import (
	"github.com/rdtigges/cicd/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	// Arrange
	word := "hello"
	language := "english"

	// Act
	res := translation.Translate(word, language)

	// Assert
	if res != "hello" {
		t.Errorf(`expected "hello" but received "%s"`, res)
	}
}
