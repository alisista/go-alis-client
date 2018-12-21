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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetArticlesArticleIDCommentsParams creates a new GetArticlesArticleIDCommentsParams object
// with the default values initialized.
func NewGetArticlesArticleIDCommentsParams() *GetArticlesArticleIDCommentsParams {
	var ()
	return &GetArticlesArticleIDCommentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArticlesArticleIDCommentsParamsWithTimeout creates a new GetArticlesArticleIDCommentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArticlesArticleIDCommentsParamsWithTimeout(timeout time.Duration) *GetArticlesArticleIDCommentsParams {
	var ()
	return &GetArticlesArticleIDCommentsParams{

		timeout: timeout,
	}
}

// NewGetArticlesArticleIDCommentsParamsWithContext creates a new GetArticlesArticleIDCommentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArticlesArticleIDCommentsParamsWithContext(ctx context.Context) *GetArticlesArticleIDCommentsParams {
	var ()
	return &GetArticlesArticleIDCommentsParams{

		Context: ctx,
	}
}

// NewGetArticlesArticleIDCommentsParamsWithHTTPClient creates a new GetArticlesArticleIDCommentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArticlesArticleIDCommentsParamsWithHTTPClient(client *http.Client) *GetArticlesArticleIDCommentsParams {
	var ()
	return &GetArticlesArticleIDCommentsParams{
		HTTPClient: client,
	}
}

/*GetArticlesArticleIDCommentsParams contains all the parameters to send to the API endpoint
for the get articles article ID comments operation typically these are written to a http.Request
*/
type GetArticlesArticleIDCommentsParams struct {

	/*ArticleID
	  対象記事の指定するために使用

	*/
	ArticleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) WithTimeout(timeout time.Duration) *GetArticlesArticleIDCommentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) WithContext(ctx context.Context) *GetArticlesArticleIDCommentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) WithHTTPClient(client *http.Client) *GetArticlesArticleIDCommentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArticleID adds the articleID to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) WithArticleID(articleID string) *GetArticlesArticleIDCommentsParams {
	o.SetArticleID(articleID)
	return o
}

// SetArticleID adds the articleId to the get articles article ID comments params
func (o *GetArticlesArticleIDCommentsParams) SetArticleID(articleID string) {
	o.ArticleID = articleID
}

// WriteToRequest writes these params to a swagger request
func (o *GetArticlesArticleIDCommentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param article_id
	if err := r.SetPathParam("article_id", o.ArticleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}