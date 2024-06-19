// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TangConnectivityResponse tang connectivity response
//
// swagger:model tang_connectivity_response
type TangConnectivityResponse struct {

	// Tang check result.
	IsSuccess bool `json:"is_success,omitempty"`

	// tang server response
	TangServerResponse []*TangServerResponse `json:"tang_server_response"`
}

// Validate validates this tang connectivity response
func (m *TangConnectivityResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTangServerResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TangConnectivityResponse) validateTangServerResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.TangServerResponse) { // not required
		return nil
	}

	for i := 0; i < len(m.TangServerResponse); i++ {
		if swag.IsZero(m.TangServerResponse[i]) { // not required
			continue
		}

		if m.TangServerResponse[i] != nil {
			if err := m.TangServerResponse[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tang_server_response" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tang_server_response" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tang connectivity response based on the context it is used
func (m *TangConnectivityResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTangServerResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TangConnectivityResponse) contextValidateTangServerResponse(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TangServerResponse); i++ {

		if m.TangServerResponse[i] != nil {
			if err := m.TangServerResponse[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tang_server_response" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tang_server_response" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TangConnectivityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TangConnectivityResponse) UnmarshalBinary(b []byte) error {
	var res TangConnectivityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TangServerResponse tang server response
//
// swagger:model TangServerResponse
type TangServerResponse struct {

	// Tang response payload.
	Payload string `json:"payload,omitempty"`

	// signatures
	Signatures []*TangServerSignatures `json:"signatures"`

	// Tang URL.
	TangURL string `json:"tang_url,omitempty"`
}

// Validate validates this tang server response
func (m *TangServerResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TangServerResponse) validateSignatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Signatures) { // not required
		return nil
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tang server response based on the context it is used
func (m *TangServerResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TangServerResponse) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signatures); i++ {

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TangServerResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TangServerResponse) UnmarshalBinary(b []byte) error {
	var res TangServerResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TangServerSignatures tang server signatures
//
// swagger:model TangServerSignatures
type TangServerSignatures struct {

	// protected
	Protected string `json:"protected,omitempty"`

	// signature
	Signature string `json:"signature,omitempty"`
}

// Validate validates this tang server signatures
func (m *TangServerSignatures) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tang server signatures based on context it is used
func (m *TangServerSignatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TangServerSignatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TangServerSignatures) UnmarshalBinary(b []byte) error {
	var res TangServerSignatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
