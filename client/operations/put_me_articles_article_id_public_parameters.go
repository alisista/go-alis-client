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

	models "github.com/alisista/go-alis-client/models"
)

// NewPutMeArticlesArticleIDPublicParams creates a new PutMeArticlesArticleIDPublicParams object
// with the default values initialized.
func NewPutMeArticlesArticleIDPublicParams() *PutMeArticlesArticleIDPublicParams {
	var ()
	return &PutMeArticlesArticleIDPublicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutMeArticlesArticleIDPublicParamsWithTimeout creates a new PutMeArticlesArticleIDPublicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutMeArticlesArticleIDPublicParamsWithTimeout(timeout time.Duration) *PutMeArticlesArticleIDPublicParams {
	var ()
	return &PutMeArticlesArticleIDPublicParams{

		timeout: timeout,
	}
}

// NewPutMeArticlesArticleIDPublicParamsWithContext creates a new PutMeArticlesArticleIDPublicParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutMeArticlesArticleIDPublicParamsWithContext(ctx context.Context) *PutMeArticlesArticleIDPublicParams {
	var ()
	return &PutMeArticlesArticleIDPublicParams{

		Context: ctx,
	}
}

// NewPutMeArticlesArticleIDPublicParamsWithHTTPClient creates a new PutMeArticlesArticleIDPublicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutMeArticlesArticleIDPublicParamsWithHTTPClient(client *http.Client) *PutMeArticlesArticleIDPublicParams {
	var ()
	return &PutMeArticlesArticleIDPublicParams{
		HTTPClient: client,
	}
}

/*PutMeArticlesArticleIDPublicParams contains all the parameters to send to the API endpoint
for the put me articles article ID public operation typically these are written to a http.Request
*/
type PutMeArticlesArticleIDPublicParams struct {

	/*Article
	  article object

	*/
	Article *models.UpdateArticle
	/*ArticleID
	  対象記事の指定するために使用

	*/
	ArticleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) WithTimeout(timeout time.Duration) *PutMeArticlesArticleIDPublicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) WithContext(ctx context.Context) *PutMeArticlesArticleIDPublicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) WithHTTPClient(client *http.Client) *PutMeArticlesArticleIDPublicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArticle adds the article to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) WithArticle(article *models.UpdateArticle) *PutMeArticlesArticleIDPublicParams {
	o.SetArticle(article)
	return o
}

// SetArticle adds the article to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) SetArticle(article *models.UpdateArticle) {
	o.Article = article
}

// WithArticleID adds the articleID to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) WithArticleID(articleID string) *PutMeArticlesArticleIDPublicParams {
	o.SetArticleID(articleID)
	return o
}

// SetArticleID adds the articleId to the put me articles article ID public params
func (o *PutMeArticlesArticleIDPublicParams) SetArticleID(articleID string) {
	o.ArticleID = articleID
}

// WriteToRequest writes these params to a swagger request
func (o *PutMeArticlesArticleIDPublicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Article != nil {
		if err := r.SetBodyParam(o.Article); err != nil {
			return err
		}
	}

	// path param article_id
	if err := r.SetPathParam("article_id", o.ArticleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
