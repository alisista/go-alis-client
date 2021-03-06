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

// NewGetMeArticlesDraftsParams creates a new GetMeArticlesDraftsParams object
// with the default values initialized.
func NewGetMeArticlesDraftsParams() *GetMeArticlesDraftsParams {
	var ()
	return &GetMeArticlesDraftsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMeArticlesDraftsParamsWithTimeout creates a new GetMeArticlesDraftsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMeArticlesDraftsParamsWithTimeout(timeout time.Duration) *GetMeArticlesDraftsParams {
	var ()
	return &GetMeArticlesDraftsParams{

		timeout: timeout,
	}
}

// NewGetMeArticlesDraftsParamsWithContext creates a new GetMeArticlesDraftsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMeArticlesDraftsParamsWithContext(ctx context.Context) *GetMeArticlesDraftsParams {
	var ()
	return &GetMeArticlesDraftsParams{

		Context: ctx,
	}
}

// NewGetMeArticlesDraftsParamsWithHTTPClient creates a new GetMeArticlesDraftsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMeArticlesDraftsParamsWithHTTPClient(client *http.Client) *GetMeArticlesDraftsParams {
	var ()
	return &GetMeArticlesDraftsParams{
		HTTPClient: client,
	}
}

/*GetMeArticlesDraftsParams contains all the parameters to send to the API endpoint
for the get me articles drafts operation typically these are written to a http.Request
*/
type GetMeArticlesDraftsParams struct {

	/*ArticleID
	  ページング処理における、現在のページの最後の記事のID

	*/
	ArticleID *string
	/*Limit
	  取得件数

	*/
	Limit *int64
	/*SortKey
	  ページング処理における、現在のページの最後の記事のソートキー

	*/
	SortKey *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) WithTimeout(timeout time.Duration) *GetMeArticlesDraftsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) WithContext(ctx context.Context) *GetMeArticlesDraftsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) WithHTTPClient(client *http.Client) *GetMeArticlesDraftsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArticleID adds the articleID to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) WithArticleID(articleID *string) *GetMeArticlesDraftsParams {
	o.SetArticleID(articleID)
	return o
}

// SetArticleID adds the articleId to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) SetArticleID(articleID *string) {
	o.ArticleID = articleID
}

// WithLimit adds the limit to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) WithLimit(limit *int64) *GetMeArticlesDraftsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithSortKey adds the sortKey to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) WithSortKey(sortKey *int64) *GetMeArticlesDraftsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the get me articles drafts params
func (o *GetMeArticlesDraftsParams) SetSortKey(sortKey *int64) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetMeArticlesDraftsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ArticleID != nil {

		// query param article_id
		var qrArticleID string
		if o.ArticleID != nil {
			qrArticleID = *o.ArticleID
		}
		qArticleID := qrArticleID
		if qArticleID != "" {
			if err := r.SetQueryParam("article_id", qArticleID); err != nil {
				return err
			}
		}

	}

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

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey int64
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := swag.FormatInt64(qrSortKey)
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
