package dateexception

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

// DateException represents one purchase request
// it also includes potential error or form to include into the web page used to collect purchase requests
type DateException struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Reason    string `json:"reason"`

	Error string
}

// Validate validates the data of the purchase request sent by the user
func (dateException *DateException) Validate() bool {

	dateException.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(dateException,
		validation.Field(&dateException.StartDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&dateException.EndDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&dateException.Reason, validation.Required),
	)

	// add potential error text to dateException.Error
	if err != nil {
		dateException.Error = err.Error()
	}

	// return true if no error, false otherwise
	return dateException.Error == ""
}

// Render the web page itself given the html template and the dateException
func (dateException *DateException) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the dateException
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`StartDate`: dateException.StartDate,
		`EndDate`:   dateException.EndDate,
		`Reason`:    dateException.Reason,

		`Error`: dateException.Error,
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
