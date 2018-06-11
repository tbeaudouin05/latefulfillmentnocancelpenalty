package adduser

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/sql/connectdb"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/oauth/authorize"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/dbinteract/baainteract"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/useraccess"
)

var user useraccess.User

// Start loads the purchase request form web page - GET request
func Start(c *gin.Context) {

	authorize.AuthorizeUserAdmin(c, &user)

	// only pass form as addUserFormInput since we only want a blank form at start
	addUserFormInput := &useraccess.User{}

	// render the web page itself given the html template and the addUserFormInput
	addUserFormInput.Render(c, `template/admin/adduser/adduser.html`)
}

// AnswerForm retrieves user inputs, validate them and upload them to database - POST request
func AnswerForm(c *gin.Context) {

	authorize.AuthorizeUserAdmin(c, &user)

	r := c.Request

	// pass all the form values input by the user
	// since we want to validate these values and upload them to database
	// in case validation fails, we also want to return these values to the form for good user experience
	addUserFormInput := &useraccess.User{
		Email: r.FormValue(`email`),
		Name:  r.FormValue(`name`),
	}

	// Validate validates the addUserFormInput form user inputs
	// if validation fails, reload the purchase request form page with the initial user inputs and error messages
	if addUserFormInput.Validate() == false {
		addUserFormInput.Render(c, `template/admin/adduser.html`)
		return
	}

	// LoadToDb uploads the purchase request form user inputs (= addUserFormInput) to database
	dbBaa := connectdb.ConnectToBaa()
	err := baainteract.CreateNewUser(addUserFormInput, dbBaa)
	handleErr(c, err)

	// if everything goes well, redirect user to adduserconfirmation web page
	http.Redirect(c.Writer, r, `/admin/adduserconfirmation`, http.StatusSeeOther)
}

// ConfirmForm loads the purchase request adduserconfirmation web page - GET request
func ConfirmForm(c *gin.Context) {
	addUserFormInput := &useraccess.User{}
	// render adduserconfirmation web page
	addUserFormInput.Render(c, `template/admin/adduser/adduserconfirmation.html`)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
