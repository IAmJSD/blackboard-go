// Code generated by go-swagger; DO NOT EDIT.

package institutional_hierarchy

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

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParams creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesParams object
// with the default values initialized.
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParams() *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParamsWithTimeout creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParamsWithContext creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParamsWithContext(ctx context.Context) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParamsWithHTTPClient creates a new PostLearnAPIPublicV1InstitutionalHierarchyNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV1InstitutionalHierarchyNodesParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	var ()
	return &PostLearnAPIPublicV1InstitutionalHierarchyNodesParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV1InstitutionalHierarchyNodesParams contains all the parameters to send to the API endpoint
for the post learn API public v1 institutional hierarchy nodes operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV1InstitutionalHierarchyNodesParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input*/
	Input PostLearnAPIPublicV1InstitutionalHierarchyNodesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) WithContext(ctx context.Context) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) WithFields(fields *string) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) WithInput(input PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the post learn API public v1 institutional hierarchy nodes params
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) SetInput(input PostLearnAPIPublicV1InstitutionalHierarchyNodesBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV1InstitutionalHierarchyNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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