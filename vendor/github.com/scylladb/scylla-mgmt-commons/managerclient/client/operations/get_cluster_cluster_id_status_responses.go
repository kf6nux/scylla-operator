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

// GetClusterClusterIDStatusReader is a Reader for the GetClusterClusterIDStatus structure.
type GetClusterClusterIDStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterClusterIDStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDStatusOK creates a GetClusterClusterIDStatusOK with default headers values
func NewGetClusterClusterIDStatusOK() *GetClusterClusterIDStatusOK {
	return &GetClusterClusterIDStatusOK{}
}

/*GetClusterClusterIDStatusOK handles this case with default header values.

Cluster hosts and their statuses
*/
type GetClusterClusterIDStatusOK struct {
	Payload models.ClusterStatus
}

func (o *GetClusterClusterIDStatusOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/status][%d] getClusterClusterIdStatusOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDStatusOK) GetPayload() models.ClusterStatus {
	return o.Payload
}

func (o *GetClusterClusterIDStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDStatusDefault creates a GetClusterClusterIDStatusDefault with default headers values
func NewGetClusterClusterIDStatusDefault(code int) *GetClusterClusterIDStatusDefault {
	return &GetClusterClusterIDStatusDefault{
		_statusCode: code,
	}
}

/*GetClusterClusterIDStatusDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDStatusDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID status default response
func (o *GetClusterClusterIDStatusDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDStatusDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/status][%d] GetClusterClusterIDStatus default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDStatusDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
