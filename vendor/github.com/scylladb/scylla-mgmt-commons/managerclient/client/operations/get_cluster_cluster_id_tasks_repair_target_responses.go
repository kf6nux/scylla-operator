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

// GetClusterClusterIDTasksRepairTargetReader is a Reader for the GetClusterClusterIDTasksRepairTarget structure.
type GetClusterClusterIDTasksRepairTargetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterClusterIDTasksRepairTargetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterClusterIDTasksRepairTargetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetClusterClusterIDTasksRepairTargetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterClusterIDTasksRepairTargetOK creates a GetClusterClusterIDTasksRepairTargetOK with default headers values
func NewGetClusterClusterIDTasksRepairTargetOK() *GetClusterClusterIDTasksRepairTargetOK {
	return &GetClusterClusterIDTasksRepairTargetOK{}
}

/*GetClusterClusterIDTasksRepairTargetOK handles this case with default header values.

Repair target
*/
type GetClusterClusterIDTasksRepairTargetOK struct {
	Payload *models.RepairTarget
}

func (o *GetClusterClusterIDTasksRepairTargetOK) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/tasks/repair/target][%d] getClusterClusterIdTasksRepairTargetOK  %+v", 200, o.Payload)
}

func (o *GetClusterClusterIDTasksRepairTargetOK) GetPayload() *models.RepairTarget {
	return o.Payload
}

func (o *GetClusterClusterIDTasksRepairTargetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RepairTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterClusterIDTasksRepairTargetDefault creates a GetClusterClusterIDTasksRepairTargetDefault with default headers values
func NewGetClusterClusterIDTasksRepairTargetDefault(code int) *GetClusterClusterIDTasksRepairTargetDefault {
	return &GetClusterClusterIDTasksRepairTargetDefault{
		_statusCode: code,
	}
}

/*GetClusterClusterIDTasksRepairTargetDefault handles this case with default header values.

Unexpected error
*/
type GetClusterClusterIDTasksRepairTargetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get cluster cluster ID tasks repair target default response
func (o *GetClusterClusterIDTasksRepairTargetDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterClusterIDTasksRepairTargetDefault) Error() string {
	return fmt.Sprintf("[GET /cluster/{cluster_id}/tasks/repair/target][%d] GetClusterClusterIDTasksRepairTarget default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterClusterIDTasksRepairTargetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetClusterClusterIDTasksRepairTargetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
