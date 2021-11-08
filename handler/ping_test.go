package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestPingHandler_Status(t *testing.T) {
	responseWriter := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(responseWriter)

	handler := PingHandler{}
	handler.Status(testContext)

	type responseStruct struct {
		Message string `json:"message"`
	}

	expectedResponse := responseStruct{
		Message: "pong",
	}

	actualResponse := responseStruct{}

	json.Unmarshal([]byte(responseWriter.Body.String()), &actualResponse)

	assert.Equal(t, 200, responseWriter.Code)
	assert.Equal(t, expectedResponse, actualResponse)
}
