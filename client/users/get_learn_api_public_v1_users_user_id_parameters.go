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

// NewGetLearnAPIPublicV1UsersUserIDParams creates a new GetLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1UsersUserIDParams() *GetLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1UsersUserIDParamsWithTimeout creates a new GetLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1UsersUserIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1UsersUserIDParamsWithContext creates a new GetLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1UsersUserIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1UsersUserIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1UsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1UsersUserIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1UsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV1UsersUserIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1UsersUserIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 users user ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1UsersUserIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
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

// WithTimeout adds the timeout to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1UsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1UsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1UsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) WithFields(fields *string) *GetLearnAPIPublicV1UsersUserIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) WithUserID(userID string) *GetLearnAPIPublicV1UsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get learn API public v1 users user ID params
func (o *GetLearnAPIPublicV1UsersUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1UsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
