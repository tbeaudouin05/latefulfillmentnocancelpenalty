package soiexception

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/thomas-bamilo/sql/connectdb"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/oauth/authorize"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/dbinteract/baainteract"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/lfncpenaltyexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/soiexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/useraccess"
)

var user useraccess.User

// Start loads the purchase request form web page - GET request
func Start(c *gin.Context) {

	authorize.Authorize(c, &user)

	// only pass form as soiException since we only want a blank form at start
	soiException := &soiexception.SoiException{}

	// render the web page itself given the html template and the soiException
	soiException.Render(c, `template/form/soiexception/soiexception.html`)
}

// AnswerForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerForm(c *gin.Context) {

	authorize.Authorize(c, &user)

	r := c.Request

	// pass all the form values input by the user
	// since we want to validate these values and upload them to database
	// in case validation fails, we also want to return these values to the form for good user experience
	soiException := &soiexception.SoiException{
		BobIDSalesOrderItem: r.FormValue(`BobIDSalesOrderItem`),
		Reason:              r.FormValue(`Reason`),
	}

	// Validate validates the soiException form user inputs
	// if validation fails, reload the purchase request form page with the initial user inputs and error messages
	if soiException.Validate() == false {
		soiException.Render(c, `template/form/soiexception/soiexception.html`)
		return
	}

	iDLfncPenaltyException := xid.New()
	initiator := getVariableFromSession(c, `userName`)

	lfncPenaltyException := &lfncpenaltyexception.LfncPenaltyException{
		IDLfncPenaltyException: iDLfncPenaltyException.String(),
		Initiator:              initiator,
		SellerName:             soiException.SellerName,
		IDSeller:               soiException.IDSeller,
		BobIDSalesOrderNumber:  soiException.BobIDSalesOrderItem,
		Reason:                 soiException.Reason,
	}

	// LoadToDb uploads the purchase request form user inputs (= soiException) to database
	dbBaa := connectdb.ConnectToBaa()
	err := baainteract.LoadSoiExceptionToDb(lfncPenaltyException, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to confirmation web page
	http.Redirect(c.Writer, r, `/form/soiexceptionconfirmation`, http.StatusSeeOther)
}

// ConfirmForm loads the purchase request confirmation web page - GET request
func ConfirmForm(c *gin.Context) {
	soiException := &soiexception.SoiException{}
	// render confirmation web page
	soiException.Render(c, `template/form/soiexception/soiexceptionconfirmation.html`)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getVariableFromSession(c *gin.Context, sessionVariableStr string) string {
	// c = *gin.Context
	session := sessions.Default(c)
	sessionVariable := session.Get(sessionVariableStr)
	if sessionVariable == nil {
		fmt.Println(fmt.Errorf(`Error: missing ` + sessionVariableStr))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return `Error: missing ` + sessionVariableStr
	}
	return sessionVariable.(string)

}
