// Code generated by go-swagger; DO NOT EDIT.

package pipeline_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1beta1ListPipelineVersionsResponse v1beta1 list pipeline versions response
// swagger:model v1beta1ListPipelineVersionsResponse
type V1beta1ListPipelineVersionsResponse struct {

	// The token to list the next page of pipeline versions.
	NextPageToken string `json:"next_page_token,omitempty"`

	// The total number of pipeline versions for the given query.
	TotalSize int32 `json:"total_size,omitempty"`

	// versions
	Versions []*V1beta1PipelineVersion `json:"versions"`
}

// Validate validates this v1beta1 list pipeline versions response
func (m *V1beta1ListPipelineVersionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1beta1ListPipelineVersionsResponse) validateVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1beta1ListPipelineVersionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1beta1ListPipelineVersionsResponse) UnmarshalBinary(b []byte) error {
	var res V1beta1ListPipelineVersionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
