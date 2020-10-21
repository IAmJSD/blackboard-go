// Code generated by go-swagger; DO NOT EDIT.

package proctoring

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

// NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams creates a new GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams() *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	var ()
	return &GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParamsWithTimeout creates a new GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	var ()
	return &GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParamsWithContext creates a new GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	var ()
	return &GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParamsWithHTTPClient creates a new GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	var ()
	return &GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams contains all the parameters to send to the API endpoint
for the get learn API public v1 proctoring services proctoring service ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*ProctoringServiceID*/
	ProctoringServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) WithFields(fields *string) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProctoringServiceID adds the proctoringServiceID to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) WithProctoringServiceID(proctoringServiceID string) *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams {
	o.SetProctoringServiceID(proctoringServiceID)
	return o
}

// SetProctoringServiceID adds the proctoringServiceId to the get learn API public v1 proctoring services proctoring service ID params
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) SetProctoringServiceID(proctoringServiceID string) {
	o.ProctoringServiceID = proctoringServiceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1ProctoringServicesProctoringServiceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param proctoringServiceId
	if err := r.SetPathParam("proctoringServiceId", o.ProctoringServiceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
