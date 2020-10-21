// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPatchLearnAPIPublicV1UsersUserIDParams creates a new PatchLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized.
func NewPatchLearnAPIPublicV1UsersUserIDParams() *PatchLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1UsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearnAPIPublicV1UsersUserIDParamsWithTimeout creates a new PatchLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLearnAPIPublicV1UsersUserIDParamsWithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1UsersUserIDParams{

		timeout: timeout,
	}
}

// NewPatchLearnAPIPublicV1UsersUserIDParamsWithContext creates a new PatchLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLearnAPIPublicV1UsersUserIDParamsWithContext(ctx context.Context) *PatchLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1UsersUserIDParams{

		Context: ctx,
	}
}

// NewPatchLearnAPIPublicV1UsersUserIDParamsWithHTTPClient creates a new PatchLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLearnAPIPublicV1UsersUserIDParamsWithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &PatchLearnAPIPublicV1UsersUserIDParams{
		HTTPClient: client,
	}
}

/*PatchLearnAPIPublicV1UsersUserIDParams contains all the parameters to send to the API endpoint
for the patch learn API public v1 users user ID operation typically these are written to a http.Request
*/
type PatchLearnAPIPublicV1UsersUserIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input
	  JSON input object.

	*/
	Input PatchLearnAPIPublicV1UsersUserIDBody
	/*UserID
	 The user ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:jsmith                     |
	| userName   | userName:jsmith                       |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) WithTimeout(timeout time.Duration) *PatchLearnAPIPublicV1UsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) WithContext(ctx context.Context) *PatchLearnAPIPublicV1UsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) WithHTTPClient(client *http.Client) *PatchLearnAPIPublicV1UsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) WithFields(fields *string) *PatchLearnAPIPublicV1UsersUserIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) WithInput(input PatchLearnAPIPublicV1UsersUserIDBody) *PatchLearnAPIPublicV1UsersUserIDParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) SetInput(input PatchLearnAPIPublicV1UsersUserIDBody) {
	o.Input = input
}

// WithUserID adds the userID to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) WithUserID(userID string) *PatchLearnAPIPublicV1UsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch learn API public v1 users user ID params
func (o *PatchLearnAPIPublicV1UsersUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearnAPIPublicV1UsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if err := r.SetBodyParam(o.Input); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}