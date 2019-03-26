// Code generated by go-swagger; DO NOT EDIT.

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceACLOKCode is the HTTP code returned for type ReplaceACLOK
const ReplaceACLOKCode int = 200

/*ReplaceACLOK ACL line replaced

swagger:response replaceAclOK
*/
type ReplaceACLOK struct {

	/*
	  In: Body
	*/
	Payload *models.ACL `json:"body,omitempty"`
}

// NewReplaceACLOK creates ReplaceACLOK with default headers values
func NewReplaceACLOK() *ReplaceACLOK {

	return &ReplaceACLOK{}
}

// WithPayload adds the payload to the replace Acl o k response
func (o *ReplaceACLOK) WithPayload(payload *models.ACL) *ReplaceACLOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Acl o k response
func (o *ReplaceACLOK) SetPayload(payload *models.ACL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceACLOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceACLAcceptedCode is the HTTP code returned for type ReplaceACLAccepted
const ReplaceACLAcceptedCode int = 202

/*ReplaceACLAccepted Configuration change accepted and reload requested

swagger:response replaceAclAccepted
*/
type ReplaceACLAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.ACL `json:"body,omitempty"`
}

// NewReplaceACLAccepted creates ReplaceACLAccepted with default headers values
func NewReplaceACLAccepted() *ReplaceACLAccepted {

	return &ReplaceACLAccepted{}
}

// WithReloadID adds the reloadId to the replace Acl accepted response
func (o *ReplaceACLAccepted) WithReloadID(reloadID string) *ReplaceACLAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace Acl accepted response
func (o *ReplaceACLAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace Acl accepted response
func (o *ReplaceACLAccepted) WithPayload(payload *models.ACL) *ReplaceACLAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Acl accepted response
func (o *ReplaceACLAccepted) SetPayload(payload *models.ACL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceACLAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceACLBadRequestCode is the HTTP code returned for type ReplaceACLBadRequest
const ReplaceACLBadRequestCode int = 400

/*ReplaceACLBadRequest Bad request

swagger:response replaceAclBadRequest
*/
type ReplaceACLBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceACLBadRequest creates ReplaceACLBadRequest with default headers values
func NewReplaceACLBadRequest() *ReplaceACLBadRequest {

	return &ReplaceACLBadRequest{}
}

// WithPayload adds the payload to the replace Acl bad request response
func (o *ReplaceACLBadRequest) WithPayload(payload *models.Error) *ReplaceACLBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Acl bad request response
func (o *ReplaceACLBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceACLBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceACLNotFoundCode is the HTTP code returned for type ReplaceACLNotFound
const ReplaceACLNotFoundCode int = 404

/*ReplaceACLNotFound The specified resource was not found

swagger:response replaceAclNotFound
*/
type ReplaceACLNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceACLNotFound creates ReplaceACLNotFound with default headers values
func NewReplaceACLNotFound() *ReplaceACLNotFound {

	return &ReplaceACLNotFound{}
}

// WithPayload adds the payload to the replace Acl not found response
func (o *ReplaceACLNotFound) WithPayload(payload *models.Error) *ReplaceACLNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Acl not found response
func (o *ReplaceACLNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceACLNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceACLDefault General Error

swagger:response replaceAclDefault
*/
type ReplaceACLDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceACLDefault creates ReplaceACLDefault with default headers values
func NewReplaceACLDefault(code int) *ReplaceACLDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceACLDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace Acl default response
func (o *ReplaceACLDefault) WithStatusCode(code int) *ReplaceACLDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace Acl default response
func (o *ReplaceACLDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace Acl default response
func (o *ReplaceACLDefault) WithPayload(payload *models.Error) *ReplaceACLDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Acl default response
func (o *ReplaceACLDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceACLDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
