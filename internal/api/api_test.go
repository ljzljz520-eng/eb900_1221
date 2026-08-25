package api

import (
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
	"warehouse5s/internal/service"
	"warehouse5s/internal/storage"
)

func TestCreateEndpoint(t *testing.T) {
	st, _ := storage.Open(filepath.Join(t.TempDir(), "d"))
	defer st.Close()
	h := Server{Svc: &service.Service{Store: st, Clock: service.FixedClock{Value: "t"}}}.Routes()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/records", strings.NewReader(`{"ID":"a","Site":"dock","Inspector":"Li","Items":[{"ID":"i","Area":"x","Description":"y","Points":1}]}`))
	h.ServeHTTP(w, r)
	if w.Code != 201 {
		t.Fatal(w.Code)
	}
}
