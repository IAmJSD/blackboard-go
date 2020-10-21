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

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams() *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParamsWithTimeout creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParamsWithContext creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParamsWithHTTPClient creates a new GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams contains all the parameters to send to the API endpoint
for the get learn API public v2 courses course ID gradebook users user ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams struct {

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
	/*Limit
	  The maximum number of results to be returned. There may be less if the query returned less than the maximum.

	*/
	Limit *int32
	/*Offset
	  The number of rows to skip before beginning to return rows. An offset of 0 is the same as omitting the offset parameter.

	*/
	Offset *int32
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

// WithTimeout adds the timeout to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithCourseID(courseID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithFields(fields *string) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithLimit(limit *int32) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithOffset(offset *int32) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithUserID adds the userID to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WithUserID(userID string) *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get learn API public v2 courses course ID gradebook users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV2CoursesCourseIDGradebookUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
