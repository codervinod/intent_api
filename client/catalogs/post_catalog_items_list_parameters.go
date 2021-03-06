package catalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codervinod/intent_api/models"
)

// NewPostCatalogItemsListParams creates a new PostCatalogItemsListParams object
// with the default values initialized.
func NewPostCatalogItemsListParams() *PostCatalogItemsListParams {
	var ()
	return &PostCatalogItemsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCatalogItemsListParamsWithTimeout creates a new PostCatalogItemsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCatalogItemsListParamsWithTimeout(timeout time.Duration) *PostCatalogItemsListParams {
	var ()
	return &PostCatalogItemsListParams{

		timeout: timeout,
	}
}

// NewPostCatalogItemsListParamsWithContext creates a new PostCatalogItemsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCatalogItemsListParamsWithContext(ctx context.Context) *PostCatalogItemsListParams {
	var ()
	return &PostCatalogItemsListParams{

		Context: ctx,
	}
}

/*PostCatalogItemsListParams contains all the parameters to send to the API endpoint
for the post catalog items list operation typically these are written to a http.Request
*/
type PostCatalogItemsListParams struct {

	/*GetEntitiesRequest*/
	GetEntitiesRequest *models.CatalogListMetadata

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post catalog items list params
func (o *PostCatalogItemsListParams) WithTimeout(timeout time.Duration) *PostCatalogItemsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post catalog items list params
func (o *PostCatalogItemsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post catalog items list params
func (o *PostCatalogItemsListParams) WithContext(ctx context.Context) *PostCatalogItemsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post catalog items list params
func (o *PostCatalogItemsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithGetEntitiesRequest adds the getEntitiesRequest to the post catalog items list params
func (o *PostCatalogItemsListParams) WithGetEntitiesRequest(getEntitiesRequest *models.CatalogListMetadata) *PostCatalogItemsListParams {
	o.SetGetEntitiesRequest(getEntitiesRequest)
	return o
}

// SetGetEntitiesRequest adds the getEntitiesRequest to the post catalog items list params
func (o *PostCatalogItemsListParams) SetGetEntitiesRequest(getEntitiesRequest *models.CatalogListMetadata) {
	o.GetEntitiesRequest = getEntitiesRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostCatalogItemsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.GetEntitiesRequest == nil {
		o.GetEntitiesRequest = new(models.CatalogListMetadata)
	}

	if err := r.SetBodyParam(o.GetEntitiesRequest); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
