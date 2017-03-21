package controllers

import (
	"fmt"
	"os"

	"github.com/revel/revel"
	"golang.org/x/oauth2"
)

type Auth struct {
	*revel.Controller
}

var oauthCfg = &oauth2.Config{
	//TODO: put your project's Client Id here.  To be got from https://code.google.com/apis/console
	ClientID:     os.Getenv("LINKEDIN_KEY"),
	ClientSecret: os.Getenv("LINKEDIN_SECRET"),
	Endpoint: oauth2.Endpoint{
		AuthURL:  "https://www.linkedin.com/oauth/v2/authorization",
		TokenURL: "https://www.linkedin.com/oauth/v2/accessToken",
	},

	RedirectURL: "http://localhost:9000/Auth/Callback",
	Scopes:      []string{"r_basicprofile", "r_emailaddress", "w_share"},
}

func (c Auth) Index(code string) revel.Result {
	url := oauthCfg.AuthCodeURL("stuff")
	return c.Redirect(url)
}

func (c Auth) Callback(code string) revel.Result {
	tok, err := oauthCfg.Exchange(oauth2.NoContext, code)
	fmt.Println(tok)
	fmt.Println(err)
	return c.Redirect("/")
}
