package middlewares

import (
	"bitbucket.org/ezserve/ezserve/http/res"
	"github.com/gin-gonic/gin"
)
import "bitbucket.org/ezserve/ezserve/app"
import m "bitbucket.org/ezserve/ezserve/models"

func UserAuth(a *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token != "" {
			if len(token) < 8 {
				res.
					Unauthorized(c).
					Code("invalid_token").
					Message("Token is invalid").
					Send()
				return
			}

			// Remove Bearer in starts of token
			token = token[7:]
		}

		if token == "" {
			token = c.Request.URL.Query().Get("token")
		}

		if token == "" {
			token = c.Request.FormValue("token")
		}

		if token == "" {
			res.
				Unauthorized(c).
				Code("token_required").
				Message("Authorization token is required").
				Send()
			return
		}

		user := m.User{}
		a.DB.First(&user, m.User{ApiToken: token})

		if user.ID == 0 {
			res.
				Unauthorized(c).
				Code("invalid_token").
				Message("Invalid authorization token").
				Send()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
