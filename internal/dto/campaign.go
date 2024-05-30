package dto

import (
	jsonvalidator "emailn/pkg/json_validator"
	"io"
)

var jsonSchema = `{
    "$schema": "http://json-schema.org/draft-07/schema#",
    "type": "object",
    "properties": {
      "name": {
        "type": "string"
      },
      "content": {
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
    "required": ["name", "content", "contacts"]
  }`

type NewCampaign struct {
	Name     string   `json:"name"`
	Content  string   `json:"content"`
	Contacts []string `json:"contacts"`
}

func NewCampaignDto(requestBody io.ReadCloser) (*NewCampaign, string, error) {
	newCampaign, validation, err := jsonvalidator.ValidateJSON[NewCampaign](requestBody, jsonSchema)
	if err != nil {
		return nil, "", err
	}
	if validation != "" {
		return nil, validation, nil
	}
	return newCampaign, "", nil
}
