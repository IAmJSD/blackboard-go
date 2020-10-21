// Code generated by go-swagger; DO NOT EDIT.

package data_sources

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

// NewPostLearnAPIPublicV1DataSourcesParams creates a new PostLearnAPIPublicV1DataSourcesParams object
// with the default values initialized.
func NewPostLearnAPIPublicV1DataSourcesParams() *PostLearnAPIPublicV1DataSourcesParams {
	var ()
	return &PostLearnAPIPublicV1DataSourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV1DataSourcesParamsWithTimeout creates a new PostLearnAPIPublicV1DataSourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV1DataSourcesParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV1DataSourcesParams {
	var ()
	return &PostLearnAPIPublicV1DataSourcesParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV1DataSourcesParamsWithContext creates a new PostLearnAPIPublicV1DataSourcesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV1DataSourcesParamsWithContext(ctx context.Context) *PostLearnAPIPublicV1DataSourcesParams {
	var ()
	return &PostLearnAPIPublicV1DataSourcesParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV1DataSourcesParamsWithHTTPClient creates a new PostLearnAPIPublicV1DataSourcesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV1DataSourcesParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV1DataSourcesParams {
	var ()
	return &PostLearnAPIPublicV1DataSourcesParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV1DataSourcesParams contains all the parameters to send to the API endpoint
for the post learn API public v1 data sources operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV1DataSourcesParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input
	  JSON input object.

	*/
	Input PostLearnAPIPublicV1DataSourcesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV1DataSourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) WithContext(ctx context.Context) *PostLearnAPIPublicV1DataSourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV1DataSourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) WithFields(fields *string) *PostLearnAPIPublicV1DataSourcesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) WithInput(input PostLearnAPIPublicV1DataSourcesBody) *PostLearnAPIPublicV1DataSourcesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the post learn API public v1 data sources params
func (o *PostLearnAPIPublicV1DataSourcesParams) SetInput(input PostLearnAPIPublicV1DataSourcesBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV1DataSourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
