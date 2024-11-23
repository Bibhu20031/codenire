// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
package api

import (
	"encoding/json"
	"fmt"
)

// SubmissionRequest defines model for SubmissionRequest.
type SubmissionRequest struct {
	Args string `json:"args"`

	// Files Files
	Files      map[string]string `json:"files"`
	TemplateId string            `json:"templateId"`
}

// SubmissionResponse defines model for SubmissionResponse.
type SubmissionResponse struct {
	Errors *[]string                  `json:"errors,omitempty"`
	Events []SubmissionResponseEvents `json:"events"`
	Meta   *SubmissionResponse_Meta   `json:"meta,omitempty"`
	Time   *string                    `json:"time,omitempty"`
}

// SubmissionResponse_Meta defines model for SubmissionResponse.Meta.
type SubmissionResponse_Meta struct {
	Version              *string                `json:"version,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// SubmissionResponseEvents defines model for SubmissionResponseEvents.
type SubmissionResponseEvents struct {
	Kind    string `json:"kind"`
	Message string `json:"message"`
}

// RunSubmissionJSONRequestBody defines body for RunSubmission for application/json ContentType.
type RunSubmissionJSONRequestBody = SubmissionRequest

// Getter for additional properties for SubmissionResponse_Meta. Returns the specified
// element and whether it was found
func (a SubmissionResponse_Meta) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for SubmissionResponse_Meta
func (a *SubmissionResponse_Meta) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for SubmissionResponse_Meta to handle AdditionalProperties
func (a *SubmissionResponse_Meta) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["version"]; found {
		err = json.Unmarshal(raw, &a.Version)
		if err != nil {
			return fmt.Errorf("error reading 'version': %w", err)
		}
		delete(object, "version")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for SubmissionResponse_Meta to handle AdditionalProperties
func (a SubmissionResponse_Meta) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Version != nil {
		object["version"], err = json.Marshal(a.Version)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'version': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}
