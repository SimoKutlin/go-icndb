package models

// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FactResponse fact response
// swagger:model FactResponse
type FactResponse struct {

	// id
	ID int64 `json:"id,omitempty"`

	// joke
	// Required: true
	Joke       string   `json:"joke"`

	// categories
	// Required: true
	Categories []string `json:"categories"`
}

// Validate validates this fact response
func (m *FactResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoke(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FactResponse) validateCategories(formats strfmt.Registry) error {

	if err := validate.Required("categories", "body", m.Categories); err != nil {
		return err
	}

	return nil
}

func (m *FactResponse) validateJoke(formats strfmt.Registry) error {

	if err := validate.Required("joke", "body", m.Joke); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FactResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FactResponse) UnmarshalBinary(b []byte) error {
	var res FactResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
