// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DeployEnvResponse deploy env response
// swagger:model DeployEnvResponse
type DeployEnvResponse struct {

	// env
	Env string `json:"env,omitempty"`
}

// Validate validates this deploy env response
func (m *DeployEnvResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeployEnvResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeployEnvResponse) UnmarshalBinary(b []byte) error {
	var res DeployEnvResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}