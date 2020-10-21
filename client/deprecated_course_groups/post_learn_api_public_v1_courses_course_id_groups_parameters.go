// Code generated by go-swagger; DO NOT EDIT.

package deprecated_course_groups

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

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsParams creates a new PostLearnAPIPublicV1CoursesCourseIDGroupsParams object
// with the default values initialized.
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsParams() *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsParamsWithTimeout creates a new PostLearnAPIPublicV1CoursesCourseIDGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsParamsWithContext creates a new PostLearnAPIPublicV1CoursesCourseIDGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsParamsWithContext(ctx context.Context) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV1CoursesCourseIDGroupsParamsWithHTTPClient creates a new PostLearnAPIPublicV1CoursesCourseIDGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV1CoursesCourseIDGroupsParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	var ()
	return &PostLearnAPIPublicV1CoursesCourseIDGroupsParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV1CoursesCourseIDGroupsParams contains all the parameters to send to the API endpoint
for the post learn API public v1 courses course ID groups operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV1CoursesCourseIDGroupsParams struct {

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
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input
	  JSON input object.

	*/
	Input PostLearnAPIPublicV1CoursesCourseIDGroupsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) WithContext(ctx context.Context) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) WithCourseID(courseID string) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) WithFields(fields *string) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) WithInput(input PostLearnAPIPublicV1CoursesCourseIDGroupsBody) *PostLearnAPIPublicV1CoursesCourseIDGroupsParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the post learn API public v1 courses course ID groups params
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) SetInput(input PostLearnAPIPublicV1CoursesCourseIDGroupsBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV1CoursesCourseIDGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
	}

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
