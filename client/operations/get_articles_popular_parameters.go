// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetArticlesPopularParams creates a new GetArticlesPopularParams object
// with the default values initialized.
func NewGetArticlesPopularParams() *GetArticlesPopularParams {
	var ()
	return &GetArticlesPopularParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArticlesPopularParamsWithTimeout creates a new GetArticlesPopularParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArticlesPopularParamsWithTimeout(timeout time.Duration) *GetArticlesPopularParams {
	var ()
	return &GetArticlesPopularParams{

		timeout: timeout,
	}
}

// NewGetArticlesPopularParamsWithContext creates a new GetArticlesPopularParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArticlesPopularParamsWithContext(ctx context.Context) *GetArticlesPopularParams {
	var ()
	return &GetArticlesPopularParams{

		Context: ctx,
	}
}

// NewGetArticlesPopularParamsWithHTTPClient creates a new GetArticlesPopularParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArticlesPopularParamsWithHTTPClient(client *http.Client) *GetArticlesPopularParams {
	var ()
	return &GetArticlesPopularParams{
		HTTPClient: client,
	}
}

/*GetArticlesPopularParams contains all the parameters to send to the API endpoint
for the get articles popular operation typically these are written to a http.Request
*/
type GetArticlesPopularParams struct {

	/*Limit
	  取得件数

	*/
	Limit *int64
	/*Page
	  取得対象ページのoffset

	*/
	Page *int64
	/*Topic
	  検索対象のトピック名

	*/
	Topic *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get articles popular params
func (o *GetArticlesPopularParams) WithTimeout(timeout time.Duration) *GetArticlesPopularParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get articles popular params
func (o *GetArticlesPopularParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get articles popular params
func (o *GetArticlesPopularParams) WithContext(ctx context.Context) *GetArticlesPopularParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get articles popular params
func (o *GetArticlesPopularParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get articles popular params
func (o *GetArticlesPopularParams) WithHTTPClient(client *http.Client) *GetArticlesPopularParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get articles popular params
func (o *GetArticlesPopularParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get articles popular params
func (o *GetArticlesPopularParams) WithLimit(limit *int64) *GetArticlesPopularParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get articles popular params
func (o *GetArticlesPopularParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get articles popular params
func (o *GetArticlesPopularParams) WithPage(page *int64) *GetArticlesPopularParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get articles popular params
func (o *GetArticlesPopularParams) SetPage(page *int64) {
	o.Page = page
}

// WithTopic adds the topic to the get articles popular params
func (o *GetArticlesPopularParams) WithTopic(topic *string) *GetArticlesPopularParams {
	o.SetTopic(topic)
	return o
}

// SetTopic adds the topic to the get articles popular params
func (o *GetArticlesPopularParams) SetTopic(topic *string) {
	o.Topic = topic
}

// WriteToRequest writes these params to a swagger request
func (o *GetArticlesPopularParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Topic != nil {

		// query param topic
		var qrTopic string
		if o.Topic != nil {
			qrTopic = *o.Topic
		}
		qTopic := qrTopic
		if qTopic != "" {
			if err := r.SetQueryParam("topic", qTopic); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
