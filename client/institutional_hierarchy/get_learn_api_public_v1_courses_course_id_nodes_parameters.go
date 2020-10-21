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
	"github.com/go-openapi/swag"
)

// NewGetLearnAPIPublicV1CoursesCourseIDNodesParams creates a new GetLearnAPIPublicV1CoursesCourseIDNodesParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDNodesParams() *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDNodesParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDNodesParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDNodesParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDNodesParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDNodesParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDNodesParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDNodesParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDNodesParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDNodesParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDNodesParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID nodes operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDNodesParams struct {

	/*CourseID
	 The course or organization ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:math101                    |
	| courseId   | courseId:math101                      |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	CourseID string
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

// WithTimeout adds the timeout to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithLimit adds the limit to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) WithLimit(limit *int32) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) WithOffset(offset *int32) *GetLearnAPIPublicV1CoursesCourseIDNodesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 courses course ID nodes params
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
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
