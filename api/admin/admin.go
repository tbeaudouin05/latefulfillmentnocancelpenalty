package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/sql/connectdb"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/oauth/authorize"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/dbinteract/baainteract"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/adminchoice"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/useraccess"
)

var user useraccess.User

// Start loads the admin web page - GET request
func Start(c *gin.Context) {

	authorize.AuthorizeAdmin(c, &user)

	// empty adminChoice to start
	adminChoice := &adminchoice.AdminChoice{}

	// render the web page itself given the html template and the adminChoice
	adminChoice.Render(c, `template/admin/admin.html`)
}

// IDExceptionRequest populates the admin form with iDExceptionRequest options - GET request
func IDExceptionRequest(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetIDExceptionRequest queries baa_application.operation.purchase_request to return all pending iDExceptionRequest
	iDExceptionRequestTable := baainteract.GetIDExceptionRequest(dbBaa)

	//Convert the `iDExceptionRequestTable` variable to json
	iDExceptionRequestTableByte, err := json.Marshal(iDExceptionRequestTable)

	// If there is an error, print it to the console, and return a server error response to the user
	handleErr(c, err)

	// If all goes well, write the JSON list of iDExceptionRequestTable to the response
	c.Writer.Write(iDExceptionRequestTableByte)
}

// ExceptionRequest populates the admin web page with pending purchase requests - GET request
func ExceptionRequest(c *gin.Context) {

	// connect to Baa database
	dbBaa := connectdb.ConnectToBaa()
	defer dbBaa.Close()

	// GetExceptionRequest queries baa_application.operation.purchase_request to return all pending purchase requests
	exceptionRequestTable := baainteract.GetExceptionRequest(dbBaa)

	//Convert the `exceptionRequestTable` variable to json
	exceptionRequestTableByte, err := json.Marshal(exceptionRequestTable)

	// If there is an error, print it to the console, and return a server error response to the user
	handleErr(c, err)

	// If all goes well, write the JSON list of exceptionRequestTable to the response
	c.Writer.Write(exceptionRequestTableByte)
}

// AcceptRejectExceptionRequest records admin user inputs to accept or reject a purchase request
// and accept or reject the given purchase request - POST request
func AcceptRejectExceptionRequest(c *gin.Context) {

	authorize.AuthorizeAdmin(c, &user)

	// pass all the form values input by the user and the form as adminChoice
	// since we want to validate these values and upload them to database
	// in case validation fails, we also want to return these values to the form for good user experience
	adminChoice := &adminchoice.AdminChoice{
		IDExceptionRequest: c.Request.FormValue(`IDExceptionRequest`),
		AcceptReject:       c.Request.FormValue(`acceptReject`),
	}

	// Validate validates the adminChoice form user inputs
	// if validation fails, reload the purchase request form page with the initial user inputs and error messages
	if adminChoice.Validate() == false {
		adminChoice.Render(c, `template/admin/admin.html`)
		return
	}

	if adminChoice.AcceptReject == `Accept` {
		dbBaa := connectdb.ConnectToBaa()
		err := baainteract.AcceptException(adminChoice.IDExceptionRequest, dbBaa)
		handleErr(c, err)
	}

	if adminChoice.AcceptReject == `Reject` {
		dbBaa := connectdb.ConnectToBaa()
		err := baainteract.RejectException(adminChoice.IDExceptionRequest, dbBaa)
		handleErr(c, err)
	}

	// if everything goes well, reload web page
	http.Redirect(c.Writer, c.Request, `/admin`, http.StatusSeeOther)

}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
