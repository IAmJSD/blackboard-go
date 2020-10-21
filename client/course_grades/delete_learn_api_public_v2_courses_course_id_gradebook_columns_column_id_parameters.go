// Code generated by go-swagger; DO NOT EDIT.

package course_grades

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

// NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams creates a new DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized.
func NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams() *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithTimeout creates a new DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{

		timeout: timeout,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithContext creates a new DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithContext(ctx context.Context) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{

		Context: ctx,
	}
}

// NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithHTTPClient creates a new DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{
		HTTPClient: client,
	}
}

/*DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams contains all the parameters to send to the API endpoint
for the delete learn API public v2 courses course ID gradebook columns column ID operation typically these are written to a http.Request
*/
type DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams struct {

	/*ColumnID
	 The grade column ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.  The ID may also be the keyword 'finalGrade', which indicates that the course's final grade column is specified.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:column1                    |
	| finalGrade | finalGrade                            |


	*/
	ColumnID string
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
	/*OnlyIfEmpty*/
	OnlyIfEmpty *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithTimeout(timeout time.Duration) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithContext(ctx context.Context) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithHTTPClient(client *http.Client) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumnID adds the columnID to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithColumnID(columnID string) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetColumnID(columnID)
	return o
}

// SetColumnID adds the columnId to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetColumnID(columnID string) {
	o.ColumnID = columnID
}

// WithCourseID adds the courseID to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithCourseID(courseID string) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithOnlyIfEmpty adds the onlyIfEmpty to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithOnlyIfEmpty(onlyIfEmpty *bool) *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetOnlyIfEmpty(onlyIfEmpty)
	return o
}

// SetOnlyIfEmpty adds the onlyIfEmpty to the delete learn API public v2 courses course ID gradebook columns column ID params
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetOnlyIfEmpty(onlyIfEmpty *bool) {
	o.OnlyIfEmpty = onlyIfEmpty
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param columnId
	if err := r.SetPathParam("columnId", o.ColumnID); err != nil {
		return err
	}

	// path param courseId
	if err := r.SetPathParam("courseId", o.CourseID); err != nil {
		return err
	}

	if o.OnlyIfEmpty != nil {

		// query param onlyIfEmpty
		var qrOnlyIfEmpty bool
		if o.OnlyIfEmpty != nil {
			qrOnlyIfEmpty = *o.OnlyIfEmpty
		}
		qOnlyIfEmpty := swag.FormatBool(qrOnlyIfEmpty)
		if qOnlyIfEmpty != "" {
			if err := r.SetQueryParam("onlyIfEmpty", qOnlyIfEmpty); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
