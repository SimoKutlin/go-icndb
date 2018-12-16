// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/baez90/go-icndb/models"
)

// GetRandomJokeOKCode is the HTTP code returned for type GetRandomJokeOK
const GetRandomJokeOKCode int = 200

/*GetRandomJokeOK OK

swagger:response getRandomJokeOK
*/
type GetRandomJokeOK struct {

	/*
	  In: Body
	*/
	Payload *models.FactResponse `json:"body,omitempty"`
}

// NewGetRandomJokeOK creates GetRandomJokeOK with default headers values
func NewGetRandomJokeOK() *GetRandomJokeOK {

	return &GetRandomJokeOK{}
}

// WithPayload adds the payload to the get random joke o k response
func (o *GetRandomJokeOK) WithPayload(payload *models.FactResponse) *GetRandomJokeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get random joke o k response
func (o *GetRandomJokeOK) SetPayload(payload *models.FactResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRandomJokeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
