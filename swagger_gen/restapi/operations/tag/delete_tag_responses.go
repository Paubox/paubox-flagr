// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/paubox/paubox-flagr/swagger_gen/models"
)

// DeleteTagOKCode is the HTTP code returned for type DeleteTagOK
const DeleteTagOKCode int = 200

/*
DeleteTagOK deleted

swagger:response deleteTagOK
*/
type DeleteTagOK struct {
}

// NewDeleteTagOK creates DeleteTagOK with default headers values
func NewDeleteTagOK() *DeleteTagOK {

	return &DeleteTagOK{}
}

// WriteResponse to the client
func (o *DeleteTagOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*
DeleteTagDefault generic error response

swagger:response deleteTagDefault
*/
type DeleteTagDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteTagDefault creates DeleteTagDefault with default headers values
func NewDeleteTagDefault(code int) *DeleteTagDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteTagDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete tag default response
func (o *DeleteTagDefault) WithStatusCode(code int) *DeleteTagDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete tag default response
func (o *DeleteTagDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete tag default response
func (o *DeleteTagDefault) WithPayload(payload *models.Error) *DeleteTagDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete tag default response
func (o *DeleteTagDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteTagDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
