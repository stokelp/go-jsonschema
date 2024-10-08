// Code generated by github.com/stokelp/go-jsonschema, DO NOT EDIT.

package test

type OnlyModels struct {
	// MyBoolean corresponds to the JSON schema field "myBoolean".
	MyBoolean *bool `json:"myBoolean,omitempty" yaml:"myBoolean,omitempty" mapstructure:"myBoolean,omitempty"`

	// MyEnum corresponds to the JSON schema field "myEnum".
	MyEnum *OnlyModelsMyEnum `json:"myEnum,omitempty" yaml:"myEnum,omitempty" mapstructure:"myEnum,omitempty"`

	// MyInteger corresponds to the JSON schema field "myInteger".
	MyInteger *int `json:"myInteger,omitempty" yaml:"myInteger,omitempty" mapstructure:"myInteger,omitempty"`

	// MyNull corresponds to the JSON schema field "myNull".
	MyNull interface{} `json:"myNull,omitempty" yaml:"myNull,omitempty" mapstructure:"myNull,omitempty"`

	// MyNumber corresponds to the JSON schema field "myNumber".
	MyNumber *float64 `json:"myNumber,omitempty" yaml:"myNumber,omitempty" mapstructure:"myNumber,omitempty"`

	// MyString corresponds to the JSON schema field "myString".
	MyString *string `json:"myString,omitempty" yaml:"myString,omitempty" mapstructure:"myString,omitempty"`
}

type OnlyModelsMyEnum string

const OnlyModelsMyEnumX OnlyModelsMyEnum = "x"
const OnlyModelsMyEnumY OnlyModelsMyEnum = "y"
