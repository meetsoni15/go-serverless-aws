package utils

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

//InternalServerError returns a 500 error
func InternalServerError() events.APIGatewayProxyResponse {
	resp := map[string]interface{}{
		"detail": "Internal Server Error",
	}
	jsonVal, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonVal),
	}
}

//RespondError returns a dynamic status code & error
func RespondError(statusCode int, err string) events.APIGatewayProxyResponse {
	resp := map[string]interface{}{
		"detail": err,
	}
	jsonVal, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonVal),
	}
}

//RespondSuccess returns a data with statuscode
func RespondWithData(resp interface{}, statusCode int) events.APIGatewayProxyResponse {
	jsonVal, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(jsonVal),
	}
}
