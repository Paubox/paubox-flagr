// Code generated by go-swagger; DO NOT EDIT.

package variant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/paubox/paubox-flagr/swagger_gen/models"
)

// NewPutVariantParams creates a new PutVariantParams object
//
// There are no default values defined in the spec.
func NewPutVariantParams() PutVariantParams {

	return PutVariantParams{}
}

// PutVariantParams contains all the bound params for the put variant operation
// typically these are obtained from a http.Request
//
// swagger:parameters putVariant
type PutVariantParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*update a variant
	  Required: true
	  In: body
	*/
	Body *models.PutVariantRequest
	/*numeric ID of the flag
	  Required: true
	  Minimum: 1
	  In: path
	*/
	FlagID int64
	/*numeric ID of the variant
	  Required: true
	  Minimum: 1
	  In: path
	*/
	VariantID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutVariantParams() beforehand.
func (o *PutVariantParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.PutVariantRequest
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body", ""))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body", ""))
	}

	rFlagID, rhkFlagID, _ := route.Params.GetOK("flagID")
	if err := o.bindFlagID(rFlagID, rhkFlagID, route.Formats); err != nil {
		res = append(res, err)
	}

	rVariantID, rhkVariantID, _ := route.Params.GetOK("variantID")
	if err := o.bindVariantID(rVariantID, rhkVariantID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFlagID binds and validates parameter FlagID from path.
func (o *PutVariantParams) bindFlagID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("flagID", "path", "int64", raw)
	}
	o.FlagID = value

	if err := o.validateFlagID(formats); err != nil {
		return err
	}

	return nil
}

// validateFlagID carries on validations for parameter FlagID
func (o *PutVariantParams) validateFlagID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("flagID", "path", o.FlagID, 1, false); err != nil {
		return err
	}

	return nil
}

// bindVariantID binds and validates parameter VariantID from path.
func (o *PutVariantParams) bindVariantID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("variantID", "path", "int64", raw)
	}
	o.VariantID = value

	if err := o.validateVariantID(formats); err != nil {
		return err
	}

	return nil
}

// validateVariantID carries on validations for parameter VariantID
func (o *PutVariantParams) validateVariantID(formats strfmt.Registry) error {

	if err := validate.MinimumInt("variantID", "path", o.VariantID, 1, false); err != nil {
		return err
	}

	return nil
}
