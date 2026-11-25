package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// An anonymous request must be redirected AND must never reach the handler:
// gin keeps running the chain after c.Redirect unless the middleware aborts.
func TestAuthMiddlewaresStopAnonymousRequests(t *testing.T) {
	gin.SetMode(gin.TestMode)
	for name, mw := range map[string]gin.HandlerFunc{
		"AuthRequired":      AuthRequired,
		"SuperuserRequired": SuperuserRequired,
	} {
		t.Run(name, func(t *testing.T) {
			reached := false
			r := gin.New()
			r.Use(sessions.Sessions("session", cookie.NewStore([]byte("test"))))
			r.POST("/write", mw, func(c *gin.Context) { reached = true })

			w := httptest.NewRecorder()
			r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/write", nil))

			if w.Code != http.StatusFound || w.Header().Get("Location") != "/login" {
				t.Fatalf("want 302 to /login, got %d %q", w.Code, w.Header().Get("Location"))
			}
			if reached {
				t.Fatal("handler ran for an anonymous request")
			}
		})
	}
}
