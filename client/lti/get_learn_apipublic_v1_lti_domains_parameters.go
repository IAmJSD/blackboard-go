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
	"github.com/go-openapi/swag"
)

// NewGetLearnApipublicV1LtiDomainsParams creates a new GetLearnApipublicV1LtiDomainsParams object
// with the default values initialized.
func NewGetLearnApipublicV1LtiDomainsParams() *GetLearnApipublicV1LtiDomainsParams {
	var ()
	return &GetLearnApipublicV1LtiDomainsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnApipublicV1LtiDomainsParamsWithTimeout creates a new GetLearnApipublicV1LtiDomainsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnApipublicV1LtiDomainsParamsWithTimeout(timeout time.Duration) *GetLearnApipublicV1LtiDomainsParams {
	var ()
	return &GetLearnApipublicV1LtiDomainsParams{

		timeout: timeout,
	}
}

// NewGetLearnApipublicV1LtiDomainsParamsWithContext creates a new GetLearnApipublicV1LtiDomainsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnApipublicV1LtiDomainsParamsWithContext(ctx context.Context) *GetLearnApipublicV1LtiDomainsParams {
	var ()
	return &GetLearnApipublicV1LtiDomainsParams{

		Context: ctx,
	}
}

// NewGetLearnApipublicV1LtiDomainsParamsWithHTTPClient creates a new GetLearnApipublicV1LtiDomainsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnApipublicV1LtiDomainsParamsWithHTTPClient(client *http.Client) *GetLearnApipublicV1LtiDomainsParams {
	var ()
	return &GetLearnApipublicV1LtiDomainsParams{
		HTTPClient: client,
	}
}

/*GetLearnApipublicV1LtiDomainsParams contains all the parameters to send to the API endpoint
for the get learn apipublic v1 lti domains operation typically these are written to a http.Request
*/
type GetLearnApipublicV1LtiDomainsParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Limit
	  The maximum number of results to be returned. There may be less if the query returned less than the maximum.

	*/
	Limit *int32
	/*Offset
	  The number of rows to skip before beginning to return rows. An offset of 0 is the same as omitting the offset parameter.

	*/
	Offset *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) WithTimeout(timeout time.Duration) *GetLearnApipublicV1LtiDomainsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) WithContext(ctx context.Context) *GetLearnApipublicV1LtiDomainsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) WithHTTPClient(client *http.Client) *GetLearnApipublicV1LtiDomainsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) WithFields(fields *string) *GetLearnApipublicV1LtiDomainsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) WithLimit(limit *int32) *GetLearnApipublicV1LtiDomainsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) WithOffset(offset *int32) *GetLearnApipublicV1LtiDomainsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn apipublic v1 lti domains params
func (o *GetLearnApipublicV1LtiDomainsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnApipublicV1LtiDomainsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
