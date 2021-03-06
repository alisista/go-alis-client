// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutMeArticlesArticleIDPublicRepublishReader is a Reader for the PutMeArticlesArticleIDPublicRepublish structure.
type PutMeArticlesArticleIDPublicRepublishReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutMeArticlesArticleIDPublicRepublishReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutMeArticlesArticleIDPublicRepublishOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutMeArticlesArticleIDPublicRepublishOK creates a PutMeArticlesArticleIDPublicRepublishOK with default headers values
func NewPutMeArticlesArticleIDPublicRepublishOK() *PutMeArticlesArticleIDPublicRepublishOK {
	return &PutMeArticlesArticleIDPublicRepublishOK{}
}

/*PutMeArticlesArticleIDPublicRepublishOK handles this case with default header values.

successful operation
*/
type PutMeArticlesArticleIDPublicRepublishOK struct {
}

func (o *PutMeArticlesArticleIDPublicRepublishOK) Error() string {
	return fmt.Sprintf("[PUT /me/articles/{article_id}/public/republish][%d] putMeArticlesArticleIdPublicRepublishOK ", 200)
}

func (o *PutMeArticlesArticleIDPublicRepublishOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
