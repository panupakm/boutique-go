package service

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTemplateFromFile(t *testing.T) {
	tmpls, err := createMapOfTemplates("../../template")
	assert.NoError(t, err)
	assert.IsType(t, &template.Template{}, tmpls["confirmation.html"])
}
