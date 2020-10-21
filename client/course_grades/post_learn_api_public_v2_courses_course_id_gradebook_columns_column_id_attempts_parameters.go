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
)

// NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams creates a new PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams object
// with the default values initialized.
func NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams() *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	var ()
	return &PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParamsWithTimeout creates a new PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParamsWithTimeout(timeout time.Duration) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	var ()
	return &PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams{

		timeout: timeout,
	}
}

// NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParamsWithContext creates a new PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParamsWithContext(ctx context.Context) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	var ()
	return &PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams{

		Context: ctx,
	}
}

// NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParamsWithHTTPClient creates a new PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParamsWithHTTPClient(client *http.Client) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	var ()
	return &PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams{
		HTTPClient: client,
	}
}

/*PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams contains all the parameters to send to the API endpoint
for the post learn API public v2 courses course ID gradebook columns column ID attempts operation typically these are written to a http.Request
*/
type PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams struct {

	/*AttemptInput*/
	AttemptInput PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsBody
	/*ColumnID*/
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
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WithTimeout(timeout time.Duration) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WithContext(ctx context.Context) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WithHTTPClient(client *http.Client) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttemptInput adds the attemptInput to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WithAttemptInput(attemptInput PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsBody) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	o.SetAttemptInput(attemptInput)
	return o
}

// SetAttemptInput adds the attemptInput to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) SetAttemptInput(attemptInput PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsBody) {
	o.AttemptInput = attemptInput
}

// WithColumnID adds the columnID to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WithColumnID(columnID string) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	o.SetColumnID(columnID)
	return o
}

// SetColumnID adds the columnId to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) SetColumnID(columnID string) {
	o.ColumnID = columnID
}

// WithCourseID adds the courseID to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WithCourseID(courseID string) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WithFields(fields *string) *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post learn API public v2 courses course ID gradebook columns column ID attempts params
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *PostLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDAttemptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.AttemptInput); err != nil {
		return err
	}

	// path param columnId
	if err := r.SetPathParam("columnId", o.ColumnID); err != nil {
		return err
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}