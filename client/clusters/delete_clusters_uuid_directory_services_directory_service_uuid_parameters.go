package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams creates a new DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams object
// with the default values initialized.
func NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams() *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams {
	var ()
	return &DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParamsWithTimeout creates a new DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParamsWithTimeout(timeout time.Duration) *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams {
	var ()
	return &DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams{

		timeout: timeout,
	}
}

// NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParamsWithContext creates a new DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParamsWithContext(ctx context.Context) *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams {
	var ()
	return &DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams{

		Context: ctx,
	}
}

/*DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams contains all the parameters to send to the API endpoint
for the delete clusters UUID directory services directory service UUID operation typically these are written to a http.Request
*/
type DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams struct {

	/*DirectoryServiceUUID*/
	DirectoryServiceUUID string
	/*UUID
	  The UUID of the entity.

	*/
	UUID string

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) WithTimeout(timeout time.Duration) *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) WithContext(ctx context.Context) *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithDirectoryServiceUUID adds the directoryServiceUUID to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) WithDirectoryServiceUUID(directoryServiceUUID string) *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams {
	o.SetDirectoryServiceUUID(directoryServiceUUID)
	return o
}

// SetDirectoryServiceUUID adds the directoryServiceUuid to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) SetDirectoryServiceUUID(directoryServiceUUID string) {
	o.DirectoryServiceUUID = directoryServiceUUID
}

// WithUUID adds the uuid to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) WithUUID(uuid string) *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the delete clusters UUID directory services directory service UUID params
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) SetUUID(uuid string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClustersUUIDDirectoryServicesDirectoryServiceUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param directory_service_uuid
	if err := r.SetPathParam("directory_service_uuid", o.DirectoryServiceUUID); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
