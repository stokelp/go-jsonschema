// Code generated by github.com/stokelp/go-jsonschema, DO NOT EDIT.

package test

type ExtRef struct {
	// MyThing corresponds to the JSON schema field "myThing".
	MyThing *Thing `json:"myThing,omitempty" yaml:"myThing,omitempty" mapstructure:"myThing,omitempty"`

	// MyThing2 corresponds to the JSON schema field "myThing2".
	MyThing2 *Thing `json:"myThing2,omitempty" yaml:"myThing2,omitempty" mapstructure:"myThing2,omitempty"`
}

type Thing struct {
	// Name corresponds to the JSON schema field "name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
}
