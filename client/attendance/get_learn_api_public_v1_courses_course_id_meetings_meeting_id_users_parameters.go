// Code generated by go-swagger; DO NOT EDIT.

package attendance

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

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams creates a new GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams object
// with the default values initialized.
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams() *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParamsWithTimeout creates a new GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParamsWithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams{

		timeout: timeout,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParamsWithContext creates a new GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParamsWithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams{

		Context: ctx,
	}
}

// NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParamsWithHTTPClient creates a new GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParamsWithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	var ()
	return &GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams{
		HTTPClient: client,
	}
}

/*GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams contains all the parameters to send to the API endpoint
for the get learn API public v1 courses course ID meetings meeting ID users operation typically these are written to a http.Request
*/
type GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams struct {

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
	/*MeetingID*/
	MeetingID string
	/*Offset
	  The number of rows to skip before beginning to return rows. An offset of 0 is the same as omitting the offset parameter.

	*/
	Offset *int32
	/*Sort*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithTimeout(timeout time.Duration) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithContext(ctx context.Context) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithHTTPClient(client *http.Client) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCourseID adds the courseID to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithCourseID(courseID string) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetCourseID(courseID)
	return o
}

// SetCourseID adds the courseId to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetCourseID(courseID string) {
	o.CourseID = courseID
}

// WithFields adds the fields to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithFields(fields *string) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithLimit(limit *int32) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithMeetingID adds the meetingID to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithMeetingID(meetingID string) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetMeetingID(meetingID)
	return o
}

// SetMeetingID adds the meetingId to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetMeetingID(meetingID string) {
	o.MeetingID = meetingID
}

// WithOffset adds the offset to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithOffset(offset *int32) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WithSort(sort *string) *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get learn API public v1 courses course ID meetings meeting ID users params
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetLearnAPIPublicV1CoursesCourseIDMeetingsMeetingIDUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param meetingId
	if err := r.SetPathParam("meetingId", o.MeetingID); err != nil {
		return err
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

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
