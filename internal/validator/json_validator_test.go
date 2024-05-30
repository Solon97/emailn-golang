package validator

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type structSchema struct {
	Name     string   `json:"name"`
	Contacts []string `json:"contacts"`
}

var schema = `{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "type": "object",
    "properties": {
      "name": {
        "type": "string"
      },
      "contacts": {
        "type": "array",
        "items": {
          "type": "string",
          "format": "email"
        }
      }
    },
    "required": ["name", "contacts"]
  }`

func TestValidateJSON(t *testing.T) {
	assert := assert.New(t)

	t.Run("Valid JSON", func(t *testing.T) {
		requestBody := io.NopCloser(strings.NewReader(`{"name": "value", "contacts": ["value@test.com"]}`))
		structValue, validationMessage, err := ValidateJSON[structSchema](requestBody, schema)
		assert.Nil(err)
		assert.Empty(validationMessage)
		assert.Equal("value", structValue.Name)
		assert.Equal([]string{"value@test.com"}, structValue.Contacts)
	})

	t.Run("Empty request body", func(t *testing.T) {
		_, _, err := ValidateJSON[structSchema](nil, schema)
		assert.NotNil(err)
	})

	t.Run("Empty JSON schema", func(t *testing.T) {
		requestBody := io.NopCloser(strings.NewReader(`{"key": "value"}`))
		_, _, err := ValidateJSON[structSchema](requestBody, "")
		assert.NotNil(err)
	})

	t.Run("Invalid JSON in request body", func(t *testing.T) {
		requestBody := io.NopCloser(strings.NewReader(`{"key": "value"}`))
		_, validationMessage, err := ValidateJSON[structSchema](requestBody, schema)
		assert.Nil(err)
		assert.NotEmpty(validationMessage)
	})

	t.Run("JSON with invalid type", func(t *testing.T) {
		requestBody := io.NopCloser(strings.NewReader(`{"name": 1, "contacts": ["value@test.com"]}`))
		_, validationMessage, err := ValidateJSON[structSchema](requestBody, schema)
		assert.Nil(err)
		assert.NotEmpty(validationMessage)
		assert.Equal("name: Invalid type. Expected: string, given: integer", validationMessage)
	})
}
