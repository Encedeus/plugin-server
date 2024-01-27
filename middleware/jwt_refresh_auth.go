package middleware

import (
	"context"
	"github.com/Encedeus/pluginServer/services"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ContextWithIDFromRefresh(ctx context.Context, userId uuid.UUID) context.Context {
	return context.WithValue(ctx, contextKey(2), userId.String())
}

func IdFromRefreshContext(ctx context.Context) (uuid.UUID, error) {
	return uuid.Parse(ctx.Value(contextKey(2)).(string))
}

// RefreshJWTAuth serves as a middleware for authorization via the refresh token
func RefreshJWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		token, err := services.GetRefreshTokenFromCookie(c)
		if err != nil || len(token) == 0 {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		isValid, refreshToken, err := services.ValidateRefreshJWT(token)
		if !isValid || err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "unauthorised",
			})
		}

		c.SetRequest(
			c.Request().WithContext(
				ContextWithIDFromRefresh(
					c.Request().Context(),
					refreshToken.UserId,
				),
			),
		)

		c.Request().Header.Set("UUID", refreshToken.UserId.String())
		return next(c)
	}
}
