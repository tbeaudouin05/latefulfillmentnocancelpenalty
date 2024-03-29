package homepage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/oauth/authorize"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/row/useraccess"
)

var user useraccess.User

// Start loads the first web page of the application - GET request
func Start(c *gin.Context) {

	authorize.Authorize(c, &user)

	session := sessions.Default(c)
	userError := session.Get("unauthorized")
	if userError != nil {
		user.Error = userError.(string)
		user.Render(c, `template/unauthorized.html`)
		return
	}

	userName := session.Get("userName")
	if userName == nil {
		fmt.Println(fmt.Errorf(`Error: missing userName`))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.Name = userName.(string)
	user.Success = `Welcome back ` + user.Name + `!`
	// render the web page itself given the html template
	user.Render(c, `template/index.html`)

}

func StartUnauthorized(c *gin.Context) {
	session := sessions.Default(c)
	userError := session.Get("unauthorized")
	if userError != nil {
		user.Error = userError.(string)
		user.Render(c, `template/unauthorized.html`)
		return
	}

	user.Render(c, `template/unauthorized.html`)
}

func handleErr(c *gin.Context, err error) {
	if err != nil {
		fmt.Println(fmt.Errorf(`Error: %v`, err))
		c.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
