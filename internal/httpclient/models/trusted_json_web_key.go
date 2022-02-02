// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrustedJSONWebKey trusted Json web key
//
// swagger:model trustedJsonWebKey
type TrustedJSONWebKey struct {

	// The "key_id" is key unique identifier (same as kid header in jws/jwt).
	Kid string `json:"kid,omitempty"`

	// The "set" is basically a name for a group(set) of keys. Will be the same as "issuer" in grant.
	Set string `json:"set,omitempty"`
}

// Validate validates this trusted Json web key
func (m *TrustedJSONWebKey) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrustedJSONWebKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustedJSONWebKey) UnmarshalBinary(b []byte) error {
	var res TrustedJSONWebKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}