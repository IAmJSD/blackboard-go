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

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams() *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParamsWithTimeout creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParamsWithContext creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParamsWithHTTPClient creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams contains all the parameters to send to the API endpoint
for the get learn API public v2 courses course ID gradebook columns column ID users user ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams struct {

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
	/*UserID
	 The user ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:jsmith                     |
	| userName   | userName:jsmith                       |
	| uuid       | uuid:915c7567d76d444abf1eed56aad3beb5 |


	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColumnID adds the columnID to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WithColumnID(columnID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	o.SetColumnID(columnID)
	return o
}

// SetColumnID adds the columnId to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) SetColumnID(columnID string) {
	o.ColumnID = columnID
}

// WithCourseID adds the courseID to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WithCourseID(courseID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WithFields(fields *string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithUserID adds the userID to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WithUserID(userID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get learn API public v2 courses course ID gradebook columns column ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
