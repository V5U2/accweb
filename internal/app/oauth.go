// Package app provides the main application functionality including authentication
// and server management for the ACCWeb application.
package app

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var (
	ErrUnsupportedProvider = errors.New("unsupported OAuth provider")
	ErrMissingToken        = errors.New("missing OAuth token")
)

// oauthStateString is used to prevent CSRF
const oauthStateString = "random-string" // TODO: Generate this randomly per session

type OAuthManager struct {
	config *oauth2.Config
}

func NewOAuthManager(provider, clientID, clientSecret, callbackURL string) (*OAuthManager, error) {
	var endpoint oauth2.Endpoint

	switch strings.ToLower(provider) {
	case "github":
		endpoint = github.Endpoint
	case "google":
		endpoint = google.Endpoint
	default:
		return nil, ErrUnsupportedProvider
	}

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  callbackURL,
		Scopes: []string{
			"user:email",
		},
		Endpoint: endpoint,
	}

	return &OAuthManager{config: config}, nil
}

func (m *OAuthManager) HandleLogin(c *gin.Context) {
	url := m.config.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (m *OAuthManager) HandleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid oauth state"})
		return
	}

	code := c.Query("code")
	token, err := m.config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("code exchange failed: %s", err.Error())})
		return
	}

	// Store token in session or cookie
	c.SetCookie("oauth_token", token.AccessToken, int(token.Expiry.Sub(token.Expiry).Seconds()),
		"/", "", true, true)

	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func GetUserFromOAuth(c *gin.Context) (*User, error) {
	token, err := c.Cookie("oauth_token")
	if err != nil {
		return nil, ErrMissingToken
	}

	// TODO: Implement user info fetching from OAuth provider
	// This is a placeholder implementation
	return &User{
		Admin: true, // You should implement proper role mapping based on your requirements
		Mod:   true,
	}, nil
}
