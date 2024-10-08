// Code generated by github.com/stokelp/go-jsonschema, DO NOT EDIT.

package test

import "encoding/json"

type TypedDefaultEmpty struct {
	// TopLevelDomains corresponds to the JSON schema field "topLevelDomains".
	TopLevelDomains []string `json:"topLevelDomains,omitempty" yaml:"topLevelDomains,omitempty" mapstructure:"topLevelDomains,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TypedDefaultEmpty) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain TypedDefaultEmpty
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["topLevelDomains"]; !ok || v == nil {
		plain.TopLevelDomains = []string{}
	}
	*j = TypedDefaultEmpty(plain)
	return nil
}
