// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/scylladb/scylla-mgmt-commons/managerclient/models"
)

// GetVersionReader is a Reader for the GetVersion structure.
type GetVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVersionOK creates a GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {
	return &GetVersionOK{}
}

/*GetVersionOK handles this case with default header values.

Server version
*/
type GetVersionOK struct {
	Payload *models.Version
}

func (o *GetVersionOK) Error() string {
	return fmt.Sprintf("[GET /version][%d] getVersionOK  %+v", 200, o.Payload)
}

func (o *GetVersionOK) GetPayload() *models.Version {
	return o.Payload
}

func (o *GetVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Version)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVersionDefault creates a GetVersionDefault with default headers values
func NewGetVersionDefault(code int) *GetVersionDefault {
	return &GetVersionDefault{
		_statusCode: code,
	}
}

/*GetVersionDefault handles this case with default header values.

Unexpected error
*/
type GetVersionDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get version default response
func (o *GetVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetVersionDefault) Error() string {
	return fmt.Sprintf("[GET /version][%d] GetVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetVersionDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
