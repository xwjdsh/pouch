// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SystemInfo system info
// swagger:model SystemInfo

type SystemInfo struct {

	// architecture
	Architecture string `json:"Architecture,omitempty"`

	// containers
	Containers int64 `json:"Containers,omitempty"`

	// containers paused
	ContainersPaused int64 `json:"ContainersPaused,omitempty"`

	// containers running
	ContainersRunning int64 `json:"ContainersRunning,omitempty"`

	// containers stopped
	ContainersStopped int64 `json:"ContainersStopped,omitempty"`
}

/* polymorph SystemInfo Architecture false */

/* polymorph SystemInfo Containers false */

/* polymorph SystemInfo ContainersPaused false */

/* polymorph SystemInfo ContainersRunning false */

/* polymorph SystemInfo ContainersStopped false */

// Validate validates this system info
func (m *SystemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SystemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemInfo) UnmarshalBinary(b []byte) error {
	var res SystemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
