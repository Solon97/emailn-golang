package jsonvalidator

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/xeipuuv/gojsonschema"
)

func ValidateJSON[T any](requestBody io.ReadCloser, jsonSchema string) (structBody *T, validationMessage string, err error) {
	if requestBody == nil {
		return nil, "", fmt.Errorf("request body is nil")
	}
	if jsonSchema == "" {
		return nil, "", fmt.Errorf("json schema is empty")
	}

	content, err := io.ReadAll(requestBody)
	if err != nil {
		return
	}

	loader := gojsonschema.NewStringLoader(string(content))
	schemaLoader := gojsonschema.NewStringLoader(jsonSchema)

	result, err := gojsonschema.Validate(schemaLoader, loader)
	if err != nil {
		return nil, "", err
	}
	if !result.Valid() {
		for _, desc := range result.Errors() {
			return nil, desc.String(), nil
		}
	}

	structBody = new(T)
	err = json.Unmarshal([]byte(content), structBody)
	if err != nil {
		return nil, "", err
	}
	return structBody, "", nil
}
