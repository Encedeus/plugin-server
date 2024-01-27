package middleware

import (
	"context"
	"github.com/Encedeus/pluginServer/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ContextWithIDFromAccess(ctx context.Context, userId uuid.UUID) context.Context {
	return context.WithValue(ctx, contextKey(2), userId.String())
}

func IdFromAccessContext(ctx context.Context) (uuid.UUID, error) {
	return uuid.Parse(ctx.Value(contextKey(2)).(string))
}

// AccessJWTAuth serves as a middleware for authorization via the access token
func AccessJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := c.Request().Context()

		// check if the header is empty
		if c.Request().Header.Get("Authorization") == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "no token supplied",
			})
		}

		// extract and validate JWT
		token := services.GetTokenFromHeader(c)
		isValid, accessToken, err := services.ValidateAccessJWT(token)

		if err != nil {
			return c.JSON(401, echo.Map{
				"message": "invalid or corrupted token",
			})
		}

		isOk, err := services.CanUserBeAuthorized(ctx, accessToken)

		if err != nil || !isValid || !isOk {
			return c.JSON(401, echo.Map{
				"message": "unauthorized",
			})
		}

		c.Request().Header.Set("UUID", accessToken.UserId.String())

		c.SetRequest(
			c.Request().WithContext(
				ContextWithIDFromAccess(
					c.Request().Context(),
					accessToken.UserId,
				),
			),
		)
		return next(c)
	}
}
