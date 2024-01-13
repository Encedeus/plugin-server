package services

import (
	"github.com/Encedeus/pluginServer/config"
	"github.com/Encedeus/pluginServer/errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"strings"
	"time"
)

const (
	// RefreshTokenExpireTime 1 week
	RefreshTokenExpireTime = 168 * time.Hour
	// AccessTokenExpireTime 15 minutes
	AccessTokenExpireTime = 15 * time.Minute
)

type TokenClaims struct {
	UserId uuid.UUID
	jwt.RegisteredClaims
}

// GenerateAccessToken generates an access token containing the UserId of a user that expires in 15 minutes
func GenerateAccessToken(userId uuid.UUID) (string, error) {
	tokenClaims := TokenClaims{
		UserId: userId,
	}

	tokenClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(AccessTokenExpireTime))
	tokenClaims.IssuedAt = jwt.NewNumericDate(time.Now())

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretAccess))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

// GenerateRefreshToken generates a refresh token containing the UserId of a user that expires in a week
func GenerateRefreshToken(userId uuid.UUID) (string, error) {
	tokenClaims := TokenClaims{
		UserId: userId,
	}

	// generate a token containing the user's UserId
	tokenClaims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(RefreshTokenExpireTime))
	tokenClaims.IssuedAt = jwt.NewNumericDate(time.Now())

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(config.Config.Auth.JWTSecretRefresh))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

// GetTokenPair returns an access and a refresh token
func GetTokenPair(userId uuid.UUID) (access string, refresh string, err error) {
	accessToken, err := GenerateAccessToken(userId)
	if err != nil {
		// log.Errorf("error generating access token %v", err1)
		return "", "", err
	}

	refreshToken, err := GenerateRefreshToken(userId)
	if err != nil {
		// log.Errorf("error generating refresh token %v", err2)
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// GetTokenFromHeader extracts the token from the auth header
func GetTokenFromHeader(ctx echo.Context) string {
	// removes "Bearer" in front of the token and returns the token
	return strings.TrimPrefix(ctx.Request().Header.Get("Authorization"), "Bearer ")
}

// GetRefreshTokenFromCookie extracts the refresh token from a browser cookie
func GetRefreshTokenFromCookie(ctx echo.Context) (string, error) {
	cookie, err := ctx.Cookie("encedeus_plugins_refreshToken")
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

/*type AccessTokenClaims interface {
    AccountAPIKeyClaims | TokenClaims
}*/

func ValidateAccessJWT(tokenString string) (isValid bool, tokenClaims TokenClaims, err error) {
	// parse the JWT and check the signing method

	claims := TokenClaims{}

	_, err = jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.ErrUnexpectedJWTSigningMethod
		}
		return []byte(config.Config.Auth.JWTSecretAccess), nil
	})

	if err != nil {
		return false, claims, err
	}

	//isUpdated, err := IsUserUpdated(claims.UserID, claims.IssuedAt)
	// if err != nil || isUpdated {
	//     return false, claims, err
	// }

	return true, claims, nil
}

func ValidateRefreshJWT(tokenString string) (bool, TokenClaims, error) {
	// parse the JWT and check the signing method

	claims := TokenClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, errors.ErrUnexpectedJWTSigningMethod
		}
		return []byte(config.Config.Auth.JWTSecretRefresh), nil
	})

	if err != nil {
		return false, claims, errors.ErrUnauthorized
	}

	// isUpdated, err := services.IsUserUpdated(claims.UserID, claims.IssuedAt)
	// if err != nil || isUpdated {
	//     return false, claims, err
	// }

	return true, claims, nil
}
