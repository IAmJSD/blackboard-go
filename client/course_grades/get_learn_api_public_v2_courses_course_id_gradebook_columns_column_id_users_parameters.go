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

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams object
// with the default values initialized.
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams() *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParamsWithTimeout creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParamsWithContext creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParamsWithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParamsWithHTTPClient creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams contains all the parameters to send to the API endpoint
for the get learn API public v2 courses course ID gradebook columns column ID users operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams struct {

	/*ChangeIndex
	  Retrieve only items modified after the given change index.

	**Since**: 3300.4.0

	*/
	ChangeIndex *int64
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
	/*IncludeUnpostedGrades
	  If true, calculated columns exposed in the response will be processed such that any unposted grades are included in their calculations. If false, only posted grades will be included in calculations. Entitlements course.gradebook-grades.VIEW, course.gradebook.MODIFY and course.gradebook-grades.EXECUTE are required to use this parameter.

	**Since**: 3800.4.0

	*/
	IncludeUnpostedGrades *bool
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

// WithTimeout adds the timeout to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChangeIndex adds the changeIndex to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithChangeIndex(changeIndex *int64) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetChangeIndex(changeIndex)
	return o
}

// SetChangeIndex adds the changeIndex to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetChangeIndex(changeIndex *int64) {
	o.ChangeIndex = changeIndex
}

// WithColumnID adds the columnID to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithColumnID(columnID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetColumnID(columnID)
	return o
}

// SetColumnID adds the columnId to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetColumnID(columnID string) {
	o.ColumnID = columnID
}

// WithCourseID adds the courseID to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithCourseID(courseID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithFields(fields *string) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIncludeUnpostedGrades adds the includeUnpostedGrades to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithIncludeUnpostedGrades(includeUnpostedGrades *bool) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetIncludeUnpostedGrades(includeUnpostedGrades)
	return o
}

// SetIncludeUnpostedGrades adds the includeUnpostedGrades to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetIncludeUnpostedGrades(includeUnpostedGrades *bool) {
	o.IncludeUnpostedGrades = includeUnpostedGrades
}

// WithLimit adds the limit to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithLimit(limit *int32) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WithOffset(offset *int32) *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v2 courses course ID gradebook columns column ID users params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookColumnsColumnIDUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ChangeIndex != nil {

		// query param changeIndex
		var qrChangeIndex int64
		if o.ChangeIndex != nil {
			qrChangeIndex = *o.ChangeIndex
		}
		qChangeIndex := swag.FormatInt64(qrChangeIndex)
		if qChangeIndex != "" {
			if err := r.SetQueryParam("changeIndex", qChangeIndex); err != nil {
				return err
			}
		}

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

	if o.IncludeUnpostedGrades != nil {

		// query param includeUnpostedGrades
		var qrIncludeUnpostedGrades bool
		if o.IncludeUnpostedGrades != nil {
			qrIncludeUnpostedGrades = *o.IncludeUnpostedGrades
		}
		qIncludeUnpostedGrades := swag.FormatBool(qrIncludeUnpostedGrades)
		if qIncludeUnpostedGrades != "" {
			if err := r.SetQueryParam("includeUnpostedGrades", qIncludeUnpostedGrades); err != nil {
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