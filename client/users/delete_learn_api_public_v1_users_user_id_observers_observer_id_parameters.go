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

// NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams creates a new DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams() *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	var ()
	return &DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParamsWithTimeout creates a new DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	var ()
	return &DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParamsWithContext creates a new DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	var ()
	return &DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	var ()
	return &DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v1 users user ID observers observer ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*ObserverID
	  The user ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	 | ID type    | Example                               |
	 |------------|---------------------------------------|
	 | primary    | _123_1                                |
	 | externalId | externalId:jsmith                     |
	 | userName   | userName:jsmith                       |
	 | uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	**Since**: 3500.5.0

	*/
	ObserverID string
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

// WithTimeout adds the timeout to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) WithFields(fields *string) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithObserverID adds the observerID to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) WithObserverID(observerID string) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	o.SetObserverID(observerID)
	return o
}

// SetObserverID adds the observerId to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) SetObserverID(observerID string) {
	o.ObserverID = observerID
}

// WithUserID adds the userID to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) WithUserID(userID string) *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete learn API public v1 users user ID observers observer ID params
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV1UsersUserIDObserversObserverIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param observerId
	if err := r.SetPathParam("observerId", o.ObserverID); err != nil {
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
