// Code generated by go-swagger; DO NOT EDIT.

package experiment_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ExperimentStorageState experiment storage state
// swagger:model ExperimentStorageState
type ExperimentStorageState string

const (

	// ExperimentStorageStateSTORAGESTATEUNSPECIFIED captures enum value "STORAGESTATE_UNSPECIFIED"
	ExperimentStorageStateSTORAGESTATEUNSPECIFIED ExperimentStorageState = "STORAGESTATE_UNSPECIFIED"

	// ExperimentStorageStateSTORAGESTATEAVAILABLE captures enum value "STORAGESTATE_AVAILABLE"
	ExperimentStorageStateSTORAGESTATEAVAILABLE ExperimentStorageState = "STORAGESTATE_AVAILABLE"

	// ExperimentStorageStateSTORAGESTATEARCHIVED captures enum value "STORAGESTATE_ARCHIVED"
	ExperimentStorageStateSTORAGESTATEARCHIVED ExperimentStorageState = "STORAGESTATE_ARCHIVED"
)

// for schema
var experimentStorageStateEnum []interface{}

func init() {
	var res []ExperimentStorageState
	if err := json.Unmarshal([]byte(`["STORAGESTATE_UNSPECIFIED","STORAGESTATE_AVAILABLE","STORAGESTATE_ARCHIVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		experimentStorageStateEnum = append(experimentStorageStateEnum, v)
	}
}

func (m ExperimentStorageState) validateExperimentStorageStateEnum(path, location string, value ExperimentStorageState) error {
	if err := validate.Enum(path, location, value, experimentStorageStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this experiment storage state
func (m ExperimentStorageState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateExperimentStorageStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}