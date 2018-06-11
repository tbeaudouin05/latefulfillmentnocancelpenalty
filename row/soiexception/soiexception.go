package soiexception

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/thomas-bamilo/sql/connectdb"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/dbinteract/omsinteract"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// SoiException represents one purchase request
// it also includes potential error or form to include into the web page used to collect purchase requests
type SoiException struct {
	SellerName          string `json:"seller_name"`
	IDSeller            string `json:"id_seller"`
	BobIDSalesOrderItem string `json:"bob_id_sales_order_item"`
	Reason              string `json:"reason"`

	Error string
}

// Validate validates the data of the purchase request sent by the user
func (soiException *SoiException) Validate() bool {

	dbOms := connectdb.ConnectToOms()
	sellerInfo := omsinteract.GetSupplierInfoFromBobIDSalesOrderItem(dbOms, soiException.BobIDSalesOrderItem)

	soiException.SellerName = sellerInfo[0]
	soiException.IDSeller = sellerInfo[1]

	log.Println(`SellerName: ` + soiException.SellerName)
	log.Println(`IDSeller: ` + soiException.IDSeller)

	soiException.Error = ""

	// define validation of each field of the purchase request
	err := validation.ValidateStruct(soiException,
		validation.Field(&soiException.SellerName, validation.Required),
		validation.Field(&soiException.IDSeller, validation.Required),
		validation.Field(&soiException.BobIDSalesOrderItem, validation.Required, is.Int),
		validation.Field(&soiException.Reason, validation.Required),
	)

	// add potential error text to soiException.Error
	if err != nil {
		soiException.Error = err.Error()
	}

	// return true if no error, false otherwise
	return soiException.Error == ""
}

// Render the web page itself given the html template and the soiException
func (soiException *SoiException) Render(c *gin.Context, htmlTemplate string) {
	// fetch the htmlTemplate
	tmpl, err := template.ParseFiles(htmlTemplate)
	handleErr(c, err)
	// render the htmlTemplate given the soiException
	err = tmpl.Execute(c.Writer, map[string]interface{}{
		`BobIDSalesOrderItem`: soiException.BobIDSalesOrderItem,
		`Reason`:              soiException.Reason,

		`Error`: soiException.Error,
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
