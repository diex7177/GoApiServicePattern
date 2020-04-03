package controllers

import (
	mocks "GoApiServicePattern/src/api/_mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestStatusCheckController_Ping(t *testing.T) {
	//Arrange
	expected := "Pong"
	mock := new(mocks.IStatusCheckService)
	mock.On("GetStatusHealth").Return(expected, nil)

	testService := StatusCheckController{mock}

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/ping", testService.Ping)

	//Act
	req, err := http.NewRequest(http.MethodGet, "/ping", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	// Create a response recorder so you can inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	//Assert
	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	mock.AssertExpectations(t)
}
