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

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams object
// with the default values initialized.
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams() *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParamsWithTimeout creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParamsWithContext creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParamsWithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParamsWithHTTPClient creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams contains all the parameters to send to the API endpoint
for the get learn API public v2 courses course ID gradebook columns column ID users last changed operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) WithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumnID adds the columnID to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) WithColumnID(columnID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	o.SetColumnID(columnID)
	return o
}

// SetColumnID adds the columnId to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) SetColumnID(columnID string) {
	o.ColumnID = columnID
}

// WithCourseID adds the courseID to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) WithCourseID(courseID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) WithFields(fields *string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v2 courses course ID gradebook columns column ID users last changed params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersLastChangedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
