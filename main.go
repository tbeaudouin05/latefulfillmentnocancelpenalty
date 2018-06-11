package main

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/admin"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/admin/adduser"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/form/dateexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/form/sellerdateexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/form/soiexception"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/homepage"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/oauth/authenticate"
	"github.com/thomas-bamilo/vs/latefulfillmentnocancelpenalty/api/oauth/login"
)

// launch the server and use defined functions to define routes for API calls
func main() {

	router := gin.Default()

	// creating cookie store
	store := sessions.NewCookieStore([]byte(randToken(64)))
	store.Options(sessions.Options{
		Path:   `/`,
		MaxAge: 86400 * 7,
	})

	// using the cookie store:
	router.Use(sessions.Sessions(`goquestsession`, store))

	router.GET(`/`, homepage.Start)
	router.GET(`/unauthorized`, homepage.StartUnauthorized)

	router.GET(`/login`, login.LoginHandler)
	router.GET(`/auth`, authenticate.AuthHandler)

	router.GET(`/form/dateexception`, dateexception.Start)
	router.POST(`/form/dateexception`, dateexception.AnswerForm)
	router.GET(`/form/dateexceptionconfirmation`, dateexception.ConfirmForm)

	router.GET(`/form/sellerdateexception`, sellerdateexception.Start)
	router.POST(`/form/sellerdateexception`, sellerdateexception.AnswerForm)
	router.GET(`/form/sellerdateexceptionconfirmation`, sellerdateexception.ConfirmForm)

	router.GET(`/form/soiexception`, soiexception.Start)
	router.POST(`/form/soiexception`, soiexception.AnswerForm)
	router.GET(`/form/soiexceptionconfirmation`, soiexception.ConfirmForm)

	router.GET(`/admin`, admin.Start)
	router.GET(`/admin/idexceptionrequest`, admin.IDExceptionRequest)
	router.GET(`/admin/exceptionrequest`, admin.ExceptionRequest)
	router.POST(`/admin`, admin.AcceptRejectExceptionRequest)

	router.GET(`/admin/adduser`, adduser.Start)
	router.POST(`/admin/adduser`, adduser.AnswerForm)
	router.GET(`/admin/adduserconfirmation`, adduser.ConfirmForm)

	router.Run(`192.168.100.160.xip.io:8081`)
}

// randToken returns a random token of i bytes
func randToken(i int) string {
	b := make([]byte, i)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}
