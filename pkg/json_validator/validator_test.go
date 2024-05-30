package jsonvalidator

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		validationMessage, err := ValidateJSON(requestBody, schema)
		assert.Nil(err)
		assert.Empty(validationMessage)
	})

	t.Run("Empty request body", func(t *testing.T) {
		_, err := ValidateJSON(nil, schema)
		assert.NotNil(err)
	})

	t.Run("Empty JSON schema", func(t *testing.T) {
		requestBody := io.NopCloser(strings.NewReader(`{"key": "value"}`))
		_, err := ValidateJSON(requestBody, "")
		assert.NotNil(err)
	})

	t.Run("Invalid JSON in request body", func(t *testing.T) {
		requestBody := io.NopCloser(strings.NewReader(`{"key": "value"}`))
		validationMessage, err := ValidateJSON(requestBody, schema)
		assert.Nil(err)
		assert.NotEmpty(validationMessage)
	})

	t.Run("JSON with invalid type", func(t *testing.T) {
		requestBody := io.NopCloser(strings.NewReader(`{"name": 1, "contacts": ["value@test.com"]}`))
		validationMessage, err := ValidateJSON(requestBody, schema)
		assert.Nil(err)
		assert.NotEmpty(validationMessage)
		assert.Equal("name: Invalid type. Expected: string, given: integer", validationMessage)
	})
}
