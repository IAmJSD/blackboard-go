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

// NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams creates a new PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized.
func NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams() *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithTimeout creates a new PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithTimeout(timeout time.Duration) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{

		timeout: timeout,
	}
}

// NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithContext creates a new PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithContext(ctx context.Context) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{

		Context: ctx,
	}
}

// NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithHTTPClient creates a new PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParamsWithHTTPClient(client *http.Client) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	var ()
	return &PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams{
		HTTPClient: client,
	}
}

/*PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams contains all the parameters to send to the API endpoint
for the patch learn API public v2 courses course ID gradebook columns column ID operation typically these are written to a http.Request
*/
type PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams struct {

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
	/*Fields
	  A comma-delimited list of fields to include in the response. If not specified, all fields will be returned.

	*/
	Fields *string
	/*Input*/
	Input PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithTimeout(timeout time.Duration) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithContext(ctx context.Context) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithHTTPClient(client *http.Client) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumnID adds the columnID to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithColumnID(columnID string) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetColumnID(columnID)
	return o
}

// SetColumnID adds the columnId to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetColumnID(columnID string) {
	o.ColumnID = columnID
}

// WithCourseID adds the courseID to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithCourseID(courseID string) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithFields(fields *string) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInput adds the input to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WithInput(input PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDBody) *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the patch learn API public v2 courses course ID gradebook columns column ID params
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) SetInput(input PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDBody) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *PatchLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
