package orderexception

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// OrderException represents one purchase request
// it also includes potential error or form to include into the web page used to collect purchase requests
type OrderException struct {
	OrderNr                      string `json:"id_purchase_request"`
	FkLfncPenaltyExceptionStatus string `json:"fk_lfnc_penalty_exception_status"`
	FkLfncPenaltyException       string `json:"fk_lfnc_penalty_exception"`

	Error string
}

// Validate validates the data of the purchase request sent by the user
func (orderException *OrderException) Validate() bool {

	orderException.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(orderException,
		validation.Field(&orderException.OrderNr, validation.Required, is.Int),
		validation.Field(&orderException.FkLfncPenaltyExceptionStatus, validation.Required, is.Int),
		validation.Field(&orderException.FkLfncPenaltyException, validation.Required),
	)

	// add potential error text to orderException.Error
	if err != nil {
		orderException.Error = err.Error()
	}

	// return true if no error, false otherwise
	return orderException.Error == ""
}

// Render the web page itself given the html template and the orderException
func (orderException *OrderException) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the orderException
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`OrderNr`:                      orderException.OrderNr,
		`FkLfncPenaltyExceptionStatus`: orderException.FkLfncPenaltyExceptionStatus,
		`FkLfncPenaltyException`:       orderException.FkLfncPenaltyException,

		`Error`: orderException.Error,
	})
	handleErr(c, err)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
