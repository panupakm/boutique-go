package biz

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTemplateFromFile(t *testing.T) {
	tmpls, err := createMapOfTemplates("../../templates")
	assert.NoError(t, err)
	assert.IsType(t, &template.Template{}, tmpls["confirmation.html"])
}
