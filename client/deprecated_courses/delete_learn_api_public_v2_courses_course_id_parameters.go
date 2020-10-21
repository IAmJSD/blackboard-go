// Code generated by go-swagger; DO NOT EDIT.

package deprecated_courses

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

// NewDeleteLearnAPIPublicV2CoursesCourseIDParams creates a new DeleteLearnAPIPublicV2CoursesCourseIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV2CoursesCourseIDParams() *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDParamsWithTimeout creates a new DeleteLearnAPIPublicV2CoursesCourseIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDParamsWithContext creates a new DeleteLearnAPIPublicV2CoursesCourseIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV2CoursesCourseIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV2CoursesCourseIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v2 courses course ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV2CoursesCourseIDParams struct {

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
	/*RemoveFiles
	  Whether to delete course files.  Defaults to true.

	*/
	RemoveFiles *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) WithCourseID(courseID string) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithRemoveFiles adds the removeFiles to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) WithRemoveFiles(removeFiles *bool) *DeleteLearnAPIPublicV2CoursesCourseIDParams {
	o.SetRemoveFiles(removeFiles)
	return o
}

// SetRemoveFiles adds the removeFiles to the delete learn API public v2 courses course ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) SetRemoveFiles(removeFiles *bool) {
	o.RemoveFiles = removeFiles
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV2CoursesCourseIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
	}

	if o.RemoveFiles != nil {

		// query param removeFiles
		var qrRemoveFiles bool
		if o.RemoveFiles != nil {
			qrRemoveFiles = *o.RemoveFiles
		}
		qRemoveFiles := swag.FormatBool(qrRemoveFiles)
		if qRemoveFiles != "" {
			if err := r.SetQueryParam("removeFiles", qRemoveFiles); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
