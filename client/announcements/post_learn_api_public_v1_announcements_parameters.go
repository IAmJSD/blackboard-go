// Code generated by go-swagger; DO NOT EDIT.

package announcements

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

// NewPostLearnAPIPublicV1AnnouncementsParams creates a new PostLearnAPIPublicV1AnnouncementsParams object
// with the default values initialized.
func NewPostLearnAPIPublicV1AnnouncementsParams() *PostLearnAPIPublicV1AnnouncementsParams {
	var ()
	return &PostLearnAPIPublicV1AnnouncementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV1AnnouncementsParamsWithTimeout creates a new PostLearnAPIPublicV1AnnouncementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV1AnnouncementsParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV1AnnouncementsParams {
	var ()
	return &PostLearnAPIPublicV1AnnouncementsParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV1AnnouncementsParamsWithContext creates a new PostLearnAPIPublicV1AnnouncementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV1AnnouncementsParamsWithContext(ctx context.Context) *PostLearnAPIPublicV1AnnouncementsParams {
	var ()
	return &PostLearnAPIPublicV1AnnouncementsParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV1AnnouncementsParamsWithHTTPClient creates a new PostLearnAPIPublicV1AnnouncementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV1AnnouncementsParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV1AnnouncementsParams {
	var ()
	return &PostLearnAPIPublicV1AnnouncementsParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV1AnnouncementsParams contains all the parameters to send to the API endpoint
for the post learn API public v1 announcements operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV1AnnouncementsParams struct {

	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input*/
	Input PostLearnAPIPublicV1AnnouncementsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV1AnnouncementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) WithContext(ctx context.Context) *PostLearnAPIPublicV1AnnouncementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV1AnnouncementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) WithFields(fields *string) *PostLearnAPIPublicV1AnnouncementsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) WithInput(input PostLearnAPIPublicV1AnnouncementsBody) *PostLearnAPIPublicV1AnnouncementsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the post learn API public v1 announcements params
func (o *PostLearnAPIPublicV1AnnouncementsParams) SetInput(input PostLearnAPIPublicV1AnnouncementsBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV1AnnouncementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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