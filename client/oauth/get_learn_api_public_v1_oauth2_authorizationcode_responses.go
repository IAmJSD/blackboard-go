// Code generated by go-swagger; DO NOT EDIT.

package oauth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetLearnAPIPublicV1Oauth2AuthorizationcodeReader is a Reader for the GetLearnAPIPublicV1Oauth2Authorizationcode structure.
type GetLearnAPIPublicV1Oauth2AuthorizationcodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLearnAPIPublicV1Oauth2AuthorizationcodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLearnAPIPublicV1Oauth2AuthorizationcodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLearnAPIPublicV1Oauth2AuthorizationcodeOK creates a GetLearnAPIPublicV1Oauth2AuthorizationcodeOK with default headers values
func NewGetLearnAPIPublicV1Oauth2AuthorizationcodeOK() *GetLearnAPIPublicV1Oauth2AuthorizationcodeOK {
	return &GetLearnAPIPublicV1Oauth2AuthorizationcodeOK{}
}

/*GetLearnAPIPublicV1Oauth2AuthorizationcodeOK handles this case with default header values.

OK
*/
type GetLearnAPIPublicV1Oauth2AuthorizationcodeOK struct {
}

func (o *GetLearnAPIPublicV1Oauth2AuthorizationcodeOK) Error() string {
	return fmt.Sprintf("[GET /learn/api/public/v1/oauth2/authorizationcode][%d] getLearnApiPublicV1Oauth2AuthorizationcodeOK ", 200)
}

func (o *GetLearnAPIPublicV1Oauth2AuthorizationcodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
