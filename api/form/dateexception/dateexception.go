package dateexception

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/thomas-bamilo/sql/connectdb"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/oauth/authorize"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/dbinteract/baainteract"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/dateexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/lfncpenaltyexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/useraccess"
)

var user useraccess.User

// Start loads the purchase request form web page - GET request
func Start(c *gin.Context) {

	authorize.Authorize(c, &user)

	// only pass form as dateException since we only want a blank form at start
	dateException := &dateexception.DateException{}

	// render the web page itself given the html template and the dateException
	dateException.Render(c, `template/form/dateexception/dateexception.html`)
}

// AnswerForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerForm(c *gin.Context) {

	authorize.Authorize(c, &user)

	r := c.Request

	// pass all the form values input by the user
	// since we want to validate these values and upload them to database
	// in case validation fails, we also want to return these values to the form for good user experience

	startDate, err := time.Parse("2006-1-2", r.FormValue(`StartDate`))
	startDateStr := startDate.Format("1/2/2006")
	handleErr(c, err)
	log.Println(`startDate: ` + startDateStr)
	endDate, err := time.Parse("2006-1-2", r.FormValue(`EndDate`))
	endDateStr := endDate.Format("1/2/2006")
	handleErr(c, err)
	log.Println(`endDate: ` + endDateStr)

	dateException := &dateexception.DateException{
		StartDate: startDateStr,
		EndDate:   endDateStr,
		Reason:    r.FormValue(`Reason`),
	}

	log.Println(`dateException.StartDate: ` + dateException.StartDate)

	// Validate validates the dateException form user inputs
	// if validation fails, reload the purchase request form page with the initial user inputs and error messages
	if dateException.Validate() == false {
		dateException.Render(c, `template/form/dateexception/dateexception.html`)
		return
	}

	iDLfncPenaltyException := xid.New()
	initiator := getVariableFromSession(c, `userName`)

	lfncPenaltyException := &lfncpenaltyexception.LfncPenaltyException{
		IDLfncPenaltyException: iDLfncPenaltyException.String(),
		Initiator:              initiator,
		StartDate:              dateException.StartDate,
		EndDate:                dateException.EndDate,
		Reason:                 dateException.Reason,
	}

	// LoadToDb uploads the purchase request form user inputs (= dateException) to database
	dbBaa := connectdb.ConnectToBaa()
	err = baainteract.LoadDateExceptionToDb(lfncPenaltyException, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to confirmation web page
	http.Redirect(c.Writer, r, `/form/dateexceptionconfirmation`, http.StatusSeeOther)
}

// ConfirmForm loads the purchase request confirmation web page - GET request
func ConfirmForm(c *gin.Context) {
	dateException := &dateexception.DateException{}
	// render confirmation web page
	dateException.Render(c, `template/form/dateexception/dateexceptionconfirmation.html`)
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
