// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// MeArticlesDraftsCreate me articles drafts create
// swagger:model MeArticlesDraftsCreate
type MeArticlesDraftsCreate struct {

	// body
	Body string `json:"body,omitempty"`

	// eye catch url
	EyeCatchURL string `json:"eye_catch_url,omitempty"`

	// overview
	Overview string `json:"overview,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this me articles drafts create
func (m *MeArticlesDraftsCreate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MeArticlesDraftsCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeArticlesDraftsCreate) UnmarshalBinary(b []byte) error {
	var res MeArticlesDraftsCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}