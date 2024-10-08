// Code generated by github.com/stokelp/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"

type StringAdditionalProperties struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`

	AdditionalProperties map[string]string
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *StringAdditionalProperties) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain StringAdditionalProperties
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw[""]; !ok || v == nil {
		plain.AdditionalProperties = map[string]string{}
	}
	*j = StringAdditionalProperties(plain)
	return nil
}
