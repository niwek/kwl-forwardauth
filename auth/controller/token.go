package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	responseform "github.com/niwek/kwl-forwardauth/auth/form/response"
	"github.com/niwek/kwl-forwardauth/auth/log"
)

// GetSessionData swagger:route GET /token/session-data/{token} auth getSessionData
//
// Get Session Data from token
//
// Responses:
//   default: errorResponse
//     200: authenticatedUser
func GetSessionData(c *gin.Context) {
	const moduleName = "controller.token.getSessionData"
	var err error
	token := c.Param("token")

	if len(token) != 10 {
		err = fmt.Errorf("Token %s is not of Length 10", token)
		log.Error(err, moduleName)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseform.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			ModuleName: moduleName,
		})
		return
	}

	authenticatedUser := responseform.AuthenticatedUser{
		ID:       uuid.New(),
		Username: token + "-username",
		Roles:    []string{"USER", "ADMIN"},
	}

	c.JSON(http.StatusOK, authenticatedUser)
}

// GetSessionDataFromHeader swagger:route GET /token/session-data-from-header auth getSessionDataFromHeader
//
// Get Session Data from Header
//
// Responses:
//   default: errorResponse
//     200: authenticatedUser
func GetSessionDataFromHeader(c *gin.Context) {
	const moduleName = "controller.token.getSessionData"
	var err error
	token := c.GetHeader("Authorization")

	if token == "" {
		err = fmt.Errorf("Authorization Header is empty")
		log.Error(err, moduleName)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseform.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			ModuleName: moduleName,
		})
		return
	}

	extractedToken := strings.Split(token, "Bearer ")

	if len(extractedToken) != 2 {
		err = fmt.Errorf("Authorization Header %s is not Bearer", extractedToken)
		log.Error(err, moduleName)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseform.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			ModuleName: moduleName,
		})
		return
	}

	if len(extractedToken[1]) != 10 {
		err = fmt.Errorf("Token %s is not of Length 10", extractedToken[1])
		log.Error(err, moduleName)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseform.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			ModuleName: moduleName,
		})
		return
	}

	authenticatedUser := responseform.AuthenticatedUser{
		ID:       uuid.New(),
		Username: extractedToken[1] + "-username",
		Roles:    []string{"USER", "ADMIN"},
	}

	userBytes, err := json.Marshal(authenticatedUser)
	if err != nil {
		log.Error(err, moduleName)
		c.AbortWithStatusJSON(http.StatusBadRequest, responseform.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			ModuleName: moduleName,
		})
		return
	}
	c.Header("X-Forwarded-Auth-User", string(userBytes))
	c.JSON(http.StatusOK, authenticatedUser)
	return
}
