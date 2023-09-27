package controllers

import (
	"fmt"
	"gin_golang/src/config"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// https://github.com/danilopolani/gocialite/wiki/Multi-provider-example

func RedirectHandler(c *gin.Context) {
	provider := c.Param("provider")
	providerSecrets := map[string]map[string]string{
		"github": {
			"clientID":     os.Getenv("CLIENT_ID_GH"),
			"clientSecret": os.Getenv("CLIENT_SECRET_GH"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/github/callback",
		},
		"google": {
			"clientID":     os.Getenv("CLIENT_ID_G"),
			"clientSecret": os.Getenv("CLIENT_SECRET_G"),
			"redirectURL":  os.Getenv("AUTH_REDIRECT_URL") + "/google/callback",
		},
	}

	providerScopes := map[string][]string{
		"github": []string{"public_repo"},
		"google": []string{},
	}

	providerData := providerSecrets[provider]
	actualScopes := providerScopes[provider]
	authURL, err := config.Gocial.New().
		Driver(provider).
		Scopes(actualScopes).
		Redirect(
			providerData["clientID"],
			providerData["clientSecret"],
			providerData["redirectURL"],
		)

	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	c.Redirect(http.StatusFound, authURL)
}

func CallbackHandler(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")
	provider := c.Param("provider")

	user, token, err := config.Gocial.Handle(state, code)
	if err != nil {
		c.Writer.Write([]byte("Error: " + err.Error()))
		return
	}

	fmt.Printf("%#v", token)
	fmt.Printf("%#v", user)
	fmt.Printf("%#v", provider)

	c.Writer.Write([]byte("Hi, " + user.FullName))
}
