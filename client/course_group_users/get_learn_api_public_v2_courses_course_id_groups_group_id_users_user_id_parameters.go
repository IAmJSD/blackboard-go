// Code generated by go-swagger; DO NOT EDIT.

package course_group_users

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

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams object
// with the default values initialized.
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams() *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParamsWithTimeout creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParamsWithContext creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParamsWithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParamsWithHTTPClient creates a new GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	var ()
	return &GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams contains all the parameters to send to the API endpoint
for the get learn API public v2 courses course ID groups group ID users user ID operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams struct {

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
	/*GroupID
	 The group ID.  This may be the primary ID, or any of the secondary IDs prefixed with the ID type.

	| ID type    | Example                               |
	|------------|---------------------------------------|
	| primary    | _123_1                                |
	| externalId | externalId:breakfastClub              |


	*/
	GroupID string
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

// WithTimeout adds the timeout to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WithContext(ctx context.Context) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WithCourseID(courseID string) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WithFields(fields *string) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGroupID adds the groupID to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WithGroupID(groupID string) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithUserID adds the userID to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WithUserID(userID string) *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get learn API public v2 courses course ID groups group ID users user ID params
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV2CoursesCourseIDGroupsGroupIDUsersUserIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
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
