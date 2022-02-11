// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PreviousConsentSession PreviousConsentSession PreviousConsentSession PreviousConsentSession PreviousConsentSession PreviousConsentSession The response used to return used consent requests
// same as HandledLoginRequest, just with consent_request exposed as json
//
// swagger:model PreviousConsentSession
type PreviousConsentSession struct {

	// consent request
	ConsentRequest *ConsentRequest `json:"consent_request,omitempty"`

	// grant access token audience
	GrantAccessTokenAudience StringSlicePipeDelimiter `json:"grant_access_token_audience,omitempty"`

	// grant scope
	GrantScope StringSlicePipeDelimiter `json:"grant_scope,omitempty"`

	// handled at
	// Format: date-time
	HandledAt NullTime `json:"handled_at,omitempty"`

	// Remember, if set to true, tells ORY Hydra to remember this consent authorization and reuse it if the same
	// client asks the same user for the same, or a subset of, scope.
	Remember bool `json:"remember,omitempty"`

	// RememberFor sets how long the consent authorization should be remembered for in seconds. If set to `0`, the
	// authorization will be remembered indefinitely.
	RememberFor int64 `json:"remember_for,omitempty"`

	// session
	Session *ConsentRequestSession `json:"session,omitempty"`
}

// Validate validates this previous consent session
func (m *PreviousConsentSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConsentRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantAccessTokenAudience(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGrantScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHandledAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreviousConsentSession) validateConsentRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsentRequest) { // not required
		return nil
	}

	if m.ConsentRequest != nil {
		if err := m.ConsentRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consent_request")
			}
			return err
		}
	}

	return nil
}

func (m *PreviousConsentSession) validateGrantAccessTokenAudience(formats strfmt.Registry) error {
	if swag.IsZero(m.GrantAccessTokenAudience) { // not required
		return nil
	}

	if err := m.GrantAccessTokenAudience.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grant_access_token_audience")
		}
		return err
	}

	return nil
}

func (m *PreviousConsentSession) validateGrantScope(formats strfmt.Registry) error {
	if swag.IsZero(m.GrantScope) { // not required
		return nil
	}

	if err := m.GrantScope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grant_scope")
		}
		return err
	}

	return nil
}

func (m *PreviousConsentSession) validateHandledAt(formats strfmt.Registry) error {
	if swag.IsZero(m.HandledAt) { // not required
		return nil
	}

	if err := m.HandledAt.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("handled_at")
		}
		return err
	}

	return nil
}

func (m *PreviousConsentSession) validateSession(formats strfmt.Registry) error {
	if swag.IsZero(m.Session) { // not required
		return nil
	}

	if m.Session != nil {
		if err := m.Session.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this previous consent session based on the context it is used
func (m *PreviousConsentSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConsentRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrantAccessTokenAudience(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGrantScope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHandledAt(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSession(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreviousConsentSession) contextValidateConsentRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsentRequest != nil {
		if err := m.ConsentRequest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consent_request")
			}
			return err
		}
	}

	return nil
}

func (m *PreviousConsentSession) contextValidateGrantAccessTokenAudience(ctx context.Context, formats strfmt.Registry) error {

	if err := m.GrantAccessTokenAudience.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grant_access_token_audience")
		}
		return err
	}

	return nil
}

func (m *PreviousConsentSession) contextValidateGrantScope(ctx context.Context, formats strfmt.Registry) error {

	if err := m.GrantScope.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("grant_scope")
		}
		return err
	}

	return nil
}

func (m *PreviousConsentSession) contextValidateHandledAt(ctx context.Context, formats strfmt.Registry) error {

	if err := m.HandledAt.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("handled_at")
		}
		return err
	}

	return nil
}

func (m *PreviousConsentSession) contextValidateSession(ctx context.Context, formats strfmt.Registry) error {

	if m.Session != nil {
		if err := m.Session.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreviousConsentSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreviousConsentSession) UnmarshalBinary(b []byte) error {
	var res PreviousConsentSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
