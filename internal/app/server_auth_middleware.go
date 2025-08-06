package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/V5U2/accweb/internal/pkg/cfg"
)

var (
	ErrForbidden = errors.New("access denied")
	ErrInvalidAuthMode = errors.New("invalid authentication mode")
)

type ACCWebAuthLevel int

const (
	ACCWebAuthLevel_Mod ACCWebAuthLevel = iota
	ACCWebAuthLevel_Adm
)

func ACCWebAuthMiddleware(config *cfg.Config, lvl ACCWebAuthLevel) gin.HandlerFunc {
	return func(c *gin.Context) {
		// If auth mode is none, allow all requests with admin privileges
		if config.Auth.Mode == cfg.AuthModeNone {
			c.Next()
			return
		}

		var u *User
		var err error

		switch config.Auth.Mode {
		case cfg.AuthModeStandard:
			u = GetUserFromClaims(c)
		case cfg.AuthModeOAuth:
			u, err = GetUserFromOAuth(c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
				c.Abort()
				return
			}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"msg": ErrInvalidAuthMode})
			c.Abort()
			return
		}

		if u == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
			c.Abort()
			return
		}

		if lvl == ACCWebAuthLevel_Mod && (!u.Mod && !u.Admin) {
			c.JSON(http.StatusForbidden, gin.H{"msg": ErrForbidden})
			c.Abort()
			return
		}

		if lvl == ACCWebAuthLevel_Adm && !u.Admin {
			c.JSON(http.StatusForbidden, gin.H{"msg": ErrForbidden})
			c.Abort()
			return
		}

		c.Next()
	}
	return func(c *gin.Context) {
		// If auth mode is none, allow all requests
		if config.Auth.Mode == cfg.AuthModeNone {
			c.Next()
			return
		}

		var u *User
		var err error

		switch config.Auth.Mode {
		case cfg.AuthModeStandard:
			u = GetUserFromClaims(c)
		case cfg.AuthModeOAuth:
			u, err = GetUserFromOAuth(c)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
				c.Abort()
				return
			}
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"msg": ErrInvalidAuthMode})
			c.Abort()
			return
		}

		if u == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "unauthorized"})
			c.Abort()
			return
		}

		if lvl == ACCWebAuthLevel_Mod && (!u.Mod && !u.Admin) {
			c.JSON(http.StatusForbidden, gin.H{"msg": ErrForbidden})
			c.Abort()
			return
		}

		if lvl == ACCWebAuthLevel_Adm && !u.Admin {
			c.JSON(http.StatusForbidden, gin.H{"msg": ErrForbidden})
			c.Abort()
			return
		}

		c.Next()
	}
}
