// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/baez90/go-icndb/models"
)

// GetJokeByIDOKCode is the HTTP code returned for type GetJokeByIDOK
const GetJokeByIDOKCode int = 200

/*GetJokeByIDOK Ok

swagger:response getJokeByIdOK
*/
type GetJokeByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.FactResponse `json:"body,omitempty"`
}

// NewGetJokeByIDOK creates GetJokeByIDOK with default headers values
func NewGetJokeByIDOK() *GetJokeByIDOK {

	return &GetJokeByIDOK{}
}

// WithPayload adds the payload to the get joke by Id o k response
func (o *GetJokeByIDOK) WithPayload(payload *models.FactResponse) *GetJokeByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get joke by Id o k response
func (o *GetJokeByIDOK) SetPayload(payload *models.FactResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetJokeByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetJokeByIDNotFoundCode is the HTTP code returned for type GetJokeByIDNotFound
const GetJokeByIDNotFoundCode int = 404

/*GetJokeByIDNotFound Not found

swagger:response getJokeByIdNotFound
*/
type GetJokeByIDNotFound struct {
}

// NewGetJokeByIDNotFound creates GetJokeByIDNotFound with default headers values
func NewGetJokeByIDNotFound() *GetJokeByIDNotFound {

	return &GetJokeByIDNotFound{}
}

// WriteResponse to the client
func (o *GetJokeByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
