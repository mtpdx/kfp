// Code generated by go-swagger; DO NOT EDIT.

package job_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// APIResourceType api resource type
// swagger:model apiResourceType
type APIResourceType string

const (

	// APIResourceTypeUNKNOWNRESOURCETYPE captures enum value "UNKNOWN_RESOURCE_TYPE"
	APIResourceTypeUNKNOWNRESOURCETYPE APIResourceType = "UNKNOWN_RESOURCE_TYPE"

	// APIResourceTypeEXPERIMENT captures enum value "EXPERIMENT"
	APIResourceTypeEXPERIMENT APIResourceType = "EXPERIMENT"

	// APIResourceTypeJOB captures enum value "JOB"
	APIResourceTypeJOB APIResourceType = "JOB"

	// APIResourceTypePIPELINE captures enum value "PIPELINE"
	APIResourceTypePIPELINE APIResourceType = "PIPELINE"

	// APIResourceTypePIPELINEVERSION captures enum value "PIPELINE_VERSION"
	APIResourceTypePIPELINEVERSION APIResourceType = "PIPELINE_VERSION"

	// APIResourceTypeNAMESPACE captures enum value "NAMESPACE"
	APIResourceTypeNAMESPACE APIResourceType = "NAMESPACE"
)

// for schema
var apiResourceTypeEnum []interface{}

func init() {
	var res []APIResourceType
	if err := json.Unmarshal([]byte(`["UNKNOWN_RESOURCE_TYPE","EXPERIMENT","JOB","PIPELINE","PIPELINE_VERSION","NAMESPACE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiResourceTypeEnum = append(apiResourceTypeEnum, v)
	}
}

func (m APIResourceType) validateAPIResourceTypeEnum(path, location string, value APIResourceType) error {
	if err := validate.Enum(path, location, value, apiResourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this api resource type
func (m APIResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPIResourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
