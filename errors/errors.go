package errors

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

var (
	ErrInvalidToken               = NewHttpError("invalid or corrupted token", http.StatusUnauthorized)
	ErrInvalidAPIKeyDescription   = NewValidationError("invalid API key description")
	ErrInvalidUserId              = NewValidationError("invalid user id")
	ErrInvalidIPAddress           = NewValidationError("invalid IP address")
	ErrInvalidEmail               = NewValidationError("invalid email")
	ErrInvalidUsername            = NewValidationError("invalid username")
	ErrInvalidPassword            = NewValidationError("invalid password")
	ErrInvalidPermission          = NewValidationError("invalid permission")
	ErrOldUsernameDoesNotMatch    = NewValidationError("old username does not match current one")
	ErrNewUsernameEqualsOld       = NewHttpError("new username must be different from the old one", 409)
	ErrOldPasswordDoesNotMatch    = NewValidationError("old password does not match current one")
	ErrNewPasswordEqualsOld       = NewHttpError("new password must be different from the old one", 409)
	ErrOldEmailDoesNotMatch       = NewValidationError("old email does not match current one")
	ErrNewEmailEqualsOld          = NewHttpError("new email must be different from the old one", 409)
	ErrUserNotFound               = NewHttpError("user not found", http.StatusNotFound)
	ErrWrongPassword              = errors.New("wrong password")
	ErrUsernameAlreadyTaken       = NewValidationError("username already taken")
	ErrPluginNameAlreadyTaken     = NewValidationError("plugin name already taken")
	ErrEmailAlreadyTaken          = NewValidationError("email already taken")
	ErrUnexpectedJWTSigningMethod = NewValidationError("unexpected JWT signing method")
	ErrUserDeleted                = NewHttpError("user deleted", 410)
	ErrMissingAPIKey              = NewValidationError("missing API key")
	ErrQueryFailed                = NewInternalError("db query failed")
	ErrPassTooLong                = NewValidationError("password too long")
	ErrPassTooShort               = NewValidationError("password too short")
	ErrEmailTooLong               = NewValidationError("email too long")
	ErrUsernameTooLong            = NewValidationError("username too long")
	ErrPluginNameTooLong          = NewValidationError("plugin name too long")
	ErrUsernameTooShort           = NewValidationError("username too short")
	ErrPluginNameTooShort         = NewValidationError("plugin name too short")
	ErrUnauthorized               = NewHttpError("unauthorized", http.StatusUnauthorized)
	ErrNotModified                = NewHttpError("", 304)
	ErrInvalidGithubURL           = NewValidationError("invalid github URL")
	ErrGithubRepoDoesNotExist     = NewValidationError("Non existent GitHub repo")
	ErrPluginNotFound             = NewHttpError("plugin not found", http.StatusNotFound)
	ErrInvalidPluginName          = NewValidationError("invalid plugin name")
	ErrReleaseTagDoesNotExist     = NewHttpError("release tag not found in plugin's source repository", http.StatusNotFound)
	ErrReleaseUsernameTooLong     = NewValidationError("release name too long")
	ErrReleaseUsernameTooShort    = NewValidationError("release name too short")
	ErrMissingPluginFile          = NewValidationError("the plugin file is missing in the selected release")
	ErrGithubApiFailed            = NewInternalError("request to github api failed")
	ErrSessionNotFound            = NewHttpError("session not found", 404)
	ErrEmailAlreadyVerified       = NewValidationError("email already verified")
	ErrReleaseNameAlreadyInUse    = NewValidationError("release name already in use")
	ErrInvalidPluginId            = NewValidationError("invalid plugin id")
	ErrReleaseNotFound            = NewHttpError("release not found", 404)
	ErrPluginHasNoReleases        = NewHttpError("plugin has no releases", 404)
	ErrRepoGone                   = NewHttpError("github repository no longer available", 410)
	ErrNoReadme                   = NewHttpError("github repository has no README.md", 410)
	ErrInvalidPageNumber          = NewValidationError("invalid page number")
	ErrNoPluginsMatchSearch       = NewHttpError("no matches found", 404)
	ErrNoRefreshToken             = NewHttpError("no refresh token provided", 401)
)

func GetHTTPErrorResponse(c echo.Context, err error) error {

	if IsValidationError(err) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if IsInternalError(err) {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	// must be last in the chain bc it sets the error parameter, this is done to fallback on code below if it fails
	if IsHTTPError(err) {
		var httpErr *HTTPError
		httpErr, err = ErrorToHttpError(err)

		if err == nil {
			return c.JSON(httpErr.code, echo.Map{
				"message": httpErr.message,
			})
		}
	}

	log.Errorf("uncaught error trying to serve request: %v", err)

	return c.JSON(http.StatusInternalServerError, echo.Map{
		"message": "internal server error",
	})
}
