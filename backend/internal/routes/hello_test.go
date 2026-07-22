package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func setup(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/hello", Hello)
	return r
}

func TestHello_Default(t *testing.T) {
	r := setup(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/hello", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d", w.Code)
	}
	var body map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	if body["message"] != "Hello Skilluv!" {
		t.Fatalf("message = %v", body["message"])
	}
}

func TestHello_WithName(t *testing.T) {
	r := setup(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/hello?name=Ada", nil)
	r.ServeHTTP(w, req)
	var body map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	if body["message"] != "Hello Ada!" {
		t.Fatalf("message = %v", body["message"])
	}
}

func TestHello_TrimsWhitespace(t *testing.T) {
	r := setup(t)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/hello?name="+strings.ReplaceAll("   ", " ", "%20"), nil)
	r.ServeHTTP(w, req)
	var body map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	if body["message"] != "Hello Skilluv!" {
		t.Fatalf("message = %v", body["message"])
	}
}
