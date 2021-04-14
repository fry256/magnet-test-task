// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteUserParams creates a new DeleteUserParams object
//
// There are no default values defined in the spec.
func NewDeleteUserParams() DeleteUserParams {

	return DeleteUserParams{}
}

// DeleteUserParams contains all the bound params for the delete user operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteUser
type DeleteUserParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The userID that needs to be deleted
	  Required: true
	  In: path
	*/
	UID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteUserParams() beforehand.
func (o *DeleteUserParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rUID, rhkUID, _ := route.Params.GetOK("uid")
	if err := o.bindUID(rUID, rhkUID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindUID binds and validates parameter UID from path.
func (o *DeleteUserParams) bindUID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("uid", "path", "int64", raw)
	}
	o.UID = value

	return nil
}
