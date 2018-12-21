// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MeWalletBalance me wallet balance
// swagger:model MeWalletBalance
type MeWalletBalance struct {

	// private eth address
	PrivateEthAddress string `json:"private_eth_address,omitempty"`
}

// Validate validates this me wallet balance
func (m *MeWalletBalance) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MeWalletBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeWalletBalance) UnmarshalBinary(b []byte) error {
	var res MeWalletBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}