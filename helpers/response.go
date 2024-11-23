package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/mnc-test-tahap2/params"
	log "github.com/sirupsen/logrus"
)

func HandleErrorController(c *gin.Context, statusCode int, errorMsg string) {
	log.Errorln(ERROR_STRING, errorMsg)

	c.AbortWithStatusJSON(statusCode, params.Response{
		Status:  statusCode,
		Payload: throwCustomMessageIfServerError(statusCode, errorMsg),
	})
}

func HandleErrorService(statusCode int, errMsg string) *params.Response {
	log.Errorln("ERROR:", errMsg)

	return &params.Response{
		Status: statusCode,
		Payload: params.ResponseErrorMessage{
			Message: throwCustomMessageIfServerError(statusCode, errMsg),
		},
	}
}

func throwCustomMessageIfServerError(statusCode int, errMsg string) string {
	if statusCode >= 500 && statusCode < 600 {
		return "An unexpected error occurred. Please try again later."
	}

	return errMsg
}
