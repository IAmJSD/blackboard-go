// Code generated by go-swagger; DO NOT EDIT.

package course_announcements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new course announcements API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for course announcements API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID(params *DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDNoContent, error)

	GetLearnAPIPublicV1CoursesCourseIDAnnouncements(params *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDAnnouncementsOK, error)

	GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID(params *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDOK, error)

	PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID(params *PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDOK, error)

	PostLearnAPIPublicV1CoursesCourseIDAnnouncements(params *PostLearnAPIPublicV1CoursesCourseIDAnnouncementsParams, authInfo runtime.ClientAuthInfoWriter) (*PostLearnAPIPublicV1CoursesCourseIDAnnouncementsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID deletes announcement

  Delete a Course Announcement. Users with the 'course.announcements.DELETE' and 'course.announcements.VIEW' entitlements can delete Course Announcements.

**Since**: 3500.3.0
*/
func (a *Client) DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID(params *DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID",
		Method:             "DELETE",
		PathPattern:        "/learn/api/public/v1/courses/{courseId}/announcements/{announcementId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLearnAPIPublicV1CoursesCourseIDAnnouncements gets announcements

  Return a list of Course Announcements. Users with the 'course.announcements.VIEW' entitlement can view 'available' Course Announcements. Users with the 'course.announcements.VIEW' & 'course.announcements.MODIFY' entitlement can view 'available' & 'unavailable' Course Announcements.

**Since**: 3500.3.0
*/
func (a *Client) GetLearnAPIPublicV1CoursesCourseIDAnnouncements(params *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDAnnouncementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1CoursesCourseIDAnnouncements",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/courses/{courseId}/announcements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1CoursesCourseIDAnnouncementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1CoursesCourseIDAnnouncementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1CoursesCourseIDAnnouncements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID gets announcement

  Get a Course Announcement. Users with the 'course.announcements.VIEW' entitlement can view 'available' Course Announcements. Users with the 'course.announcements.VIEW' & 'course.announcements.MODIFY' entitlement can view 'available' & 'unavailable' Course Announcements.

**Since**: 3500.3.0
*/
func (a *Client) GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID(params *GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID",
		Method:             "GET",
		PathPattern:        "/learn/api/public/v1/courses/{courseId}/announcements/{announcementId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID updates announcement

  Update a Course Announcement. Users with the 'course.announcements.MODIFY' and 'course.announcements.VIEW' entitlements can update Course Announcements.

**Since**: 3500.3.0
*/
func (a *Client) PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID(params *PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams, authInfo runtime.ClientAuthInfoWriter) (*PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID",
		Method:             "PATCH",
		PathPattern:        "/learn/api/public/v1/courses/{courseId}/announcements/{announcementId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchLearnAPIPublicV1CoursesCourseIDAnnouncementsAnnouncementID: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostLearnAPIPublicV1CoursesCourseIDAnnouncements creates announcement

  Create a Course Announcement. Users with the 'course.announcements.CREATE' and 'course.announcements.VIEW' entitlements can create Course Announcements.

**Since**: 3500.3.0
*/
func (a *Client) PostLearnAPIPublicV1CoursesCourseIDAnnouncements(params *PostLearnAPIPublicV1CoursesCourseIDAnnouncementsParams, authInfo runtime.ClientAuthInfoWriter) (*PostLearnAPIPublicV1CoursesCourseIDAnnouncementsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLearnAPIPublicV1CoursesCourseIDAnnouncementsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostLearnAPIPublicV1CoursesCourseIDAnnouncements",
		Method:             "POST",
		PathPattern:        "/learn/api/public/v1/courses/{courseId}/announcements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLearnAPIPublicV1CoursesCourseIDAnnouncementsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostLearnAPIPublicV1CoursesCourseIDAnnouncementsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostLearnAPIPublicV1CoursesCourseIDAnnouncements: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}