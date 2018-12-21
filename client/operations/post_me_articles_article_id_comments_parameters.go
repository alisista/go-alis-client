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

// NewPostMeArticlesArticleIDCommentsParams creates a new PostMeArticlesArticleIDCommentsParams object
// with the default values initialized.
func NewPostMeArticlesArticleIDCommentsParams() *PostMeArticlesArticleIDCommentsParams {
	var ()
	return &PostMeArticlesArticleIDCommentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostMeArticlesArticleIDCommentsParamsWithTimeout creates a new PostMeArticlesArticleIDCommentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostMeArticlesArticleIDCommentsParamsWithTimeout(timeout time.Duration) *PostMeArticlesArticleIDCommentsParams {
	var ()
	return &PostMeArticlesArticleIDCommentsParams{

		timeout: timeout,
	}
}

// NewPostMeArticlesArticleIDCommentsParamsWithContext creates a new PostMeArticlesArticleIDCommentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostMeArticlesArticleIDCommentsParamsWithContext(ctx context.Context) *PostMeArticlesArticleIDCommentsParams {
	var ()
	return &PostMeArticlesArticleIDCommentsParams{

		Context: ctx,
	}
}

// NewPostMeArticlesArticleIDCommentsParamsWithHTTPClient creates a new PostMeArticlesArticleIDCommentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostMeArticlesArticleIDCommentsParamsWithHTTPClient(client *http.Client) *PostMeArticlesArticleIDCommentsParams {
	var ()
	return &PostMeArticlesArticleIDCommentsParams{
		HTTPClient: client,
	}
}

/*PostMeArticlesArticleIDCommentsParams contains all the parameters to send to the API endpoint
for the post me articles article ID comments operation typically these are written to a http.Request
*/
type PostMeArticlesArticleIDCommentsParams struct {

	/*ArticleID
	  対象記事の指定するために使用

	*/
	ArticleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) WithTimeout(timeout time.Duration) *PostMeArticlesArticleIDCommentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) WithContext(ctx context.Context) *PostMeArticlesArticleIDCommentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) WithHTTPClient(client *http.Client) *PostMeArticlesArticleIDCommentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArticleID adds the articleID to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) WithArticleID(articleID string) *PostMeArticlesArticleIDCommentsParams {
	o.SetArticleID(articleID)
	return o
}

// SetArticleID adds the articleId to the post me articles article ID comments params
func (o *PostMeArticlesArticleIDCommentsParams) SetArticleID(articleID string) {
	o.ArticleID = articleID
}

// WriteToRequest writes these params to a swagger request
func (o *PostMeArticlesArticleIDCommentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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