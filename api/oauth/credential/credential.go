package credential

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func InitCred() *oauth2.Config {

	// Credentials which stores google ids.
	type Credentials struct {
		Cid     string `json:"client_id"`
		Csecret string `json:"client_secret"`
	}
	var cred Credentials

	file, err := ioutil.ReadFile("creds.json")
	checkError(err)
	json.Unmarshal(file, &cred)

	conf := &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://192.168.100.160.xip.io:8081/auth",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}

	return conf
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
