package test

import (
	_ "api/routers"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".."+string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	logs.Info("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())
}
