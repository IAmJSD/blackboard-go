// Code generated by go-swagger; DO NOT EDIT.

package lti

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

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams creates a new DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams() *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	var ()
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParamsWithTimeout creates a new DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	var ()
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParamsWithContext creates a new DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	var ()
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV1LtiPlacementsPlacementIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	var ()
	return &DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v1 lti placements placement ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*PlacementID*/
	PlacementID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) WithFields(fields *string) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithPlacementID adds the placementID to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) WithPlacementID(placementID string) *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams {
	o.SetPlacementID(placementID)
	return o
}

// SetPlacementID adds the placementId to the delete learn API public v1 lti placements placement ID params
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) SetPlacementID(placementID string) {
	o.PlacementID = placementID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV1LtiPlacementsPlacementIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param placementId
	if err := r.SetPathParam("placementId", o.PlacementID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
