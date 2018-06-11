package sellerdateexception

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/thomas-bamilo/sql/connectdb"

	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/dbinteract/omsinteract"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
)

// SellerDateException represents one purchase request
// it also includes potential error or form to include into the web page used to collect purchase requests
type SellerDateException struct {
	SellerName string `json:"seller_name"`
	IDSeller   string `json:"id_seller"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Reason     string `json:"reason"`

	Error string
}

// Validate validates the data of the purchase request sent by the user
func (sellerDateException *SellerDateException) Validate() bool {

	dbOms := connectdb.ConnectToOms()
	sellerDateException.IDSeller = omsinteract.GetBobIDSupplierFromSellerName(dbOms, sellerDateException.SellerName)

	sellerDateException.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(sellerDateException,
		validation.Field(&sellerDateException.SellerName, validation.Required),
		validation.Field(&sellerDateException.IDSeller, validation.Required),
		validation.Field(&sellerDateException.StartDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&sellerDateException.EndDate, validation.Required, validation.Date("1/2/2006")),
		validation.Field(&sellerDateException.Reason, validation.Required),
	)

	// add potential error text to sellerDateException.Error
	if err != nil {
		sellerDateException.Error = err.Error()
	}

	// return true if no error, false otherwise
	return sellerDateException.Error == ""
}

// Render the web page itself given the html template and the sellerDateException
func (sellerDateException *SellerDateException) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the sellerDateException
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`SellerName`: sellerDateException.SellerName,
		`IDSeller`:   sellerDateException.IDSeller,
		`StartDate`:  sellerDateException.StartDate,
		`EndDate`:    sellerDateException.EndDate,
		`Reason`:     sellerDateException.Reason,

		`Error`: sellerDateException.Error,
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
