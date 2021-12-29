package swagger_test

import (
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"testing"

	"github.com/go-mixins/swagger-ui"
)

func TestDir(t *testing.T) {
	var h = http.StripPrefix("/swagger-ui", swagger.Handler(func() ([]byte, error) {
		return []byte("{test json}"), nil
	}))
	for _, url := range []string{"", "swagger.json"} {
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/swagger-ui/"+url, nil)
		h.ServeHTTP(w, r)
		if w.Code != 200 {
			result, _ := httputil.DumpResponse(w.Result(), true)
			t.Fatalf("bad request result for %s: (%d)\n%s", url, w.Code, string(result))
		}
		s := w.Body.String()
		if len(s) == 0 {
			t.Error("zero body for", url)
		}
		t.Log(s)
	}

}
