package main

import (
	"fmt"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/api/middleware"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)
²
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/Rafael-Sapalo/FullStack-web-application/server/utils"
)

func TestGenerateToken(t *testing.T) {
	testCases := []struct {
		userId int
	}{
		{userId: 13},
		{userId: 24},
		{userId: 35},
	}
	for _, tc := range testCases {
		token, err := utils.GenerateToken(tc.userId)
		if err != nil {
			t.Errorf("Error generating toke for userID %d: %v", tc.userId, err)
		} else if token == "" {
			t.Errorf("Empty token generated for userID %d", tc.userId)
		} else {
			fmt.Printf("Token for userId %d: %s\n", tc.userId, token)
		}
	}
}

func TestAuthMiddleware(t *testing.T)  {
	var router = gin.New()
	var store = cookie.NewStore([]byte("seccret"))
	router.Use(sessions.Sessions("my-session", store))
	router.Use(middleware.Authenticate())
	router.GET("/protected", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Protected Route");
	})
	var req = httptest.NewRequest("GET", "/protected", nil)
	var w = httptest.NewRecorder()
	router.ServeHTTP(w, req);
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, `{"error":"unauthorized"}`, w.Body.String())

	// --- Test Case 2: Authorized Access ---
	// Create a new HTTP request for the protected route with a user_id session
	reqWithSession := httptest.NewRequest("GET", "/protected", nil)

	// Create a new Gin context with a session that contains the user_id
	ctx, _ := gin.CreateTestContext(reqWithSession)
	session := sessions.Default(ctx)
	session.Set("user_id", 123) // Set a user_id in the session
	session.Save()

	// Perform the request
	router.ServeHTTP(w, reqWithSession)

	// Assert that the response code is 200 OK
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Protected Route", w.Body.String())
}
