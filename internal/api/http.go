package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/query"
	"warehouse5s/internal/report"
	"warehouse5s/internal/service"
)

type Server struct{ Svc *service.Service }

func (s Server) Routes() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/records", s.records)
	m.HandleFunc("/records/", s.record)
	m.HandleFunc("/import", s.importRows)
	m.HandleFunc("/scan/", s.scan)
	return m
}
func write(w http.ResponseWriter, v any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}
func (s Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var in struct {
			ID, Site, Inspector string
			Items               []domain.Item
		}
		if json.NewDecoder(r.Body).Decode(&in) != nil {
			write(w, map[string]string{"error": "invalid json"}, 400)
			return
		}
		v, e := s.Svc.Create(in.ID, in.Site, in.Inspector, in.Items)
		if e != nil {
			write(w, map[string]string{"error": e.Error()}, 400)
			return
		}
		write(w, v, 201)
		return
	}
	f := query.Filter{Site: r.URL.Query().Get("site"), Status: r.URL.Query().Get("status")}
	v, e := query.Search(s.Svc.Store, f)
	if e != nil {
		write(w, map[string]string{"error": e.Error()}, 500)
		return
	}
	write(w, v, 200)
}
func (s Server) record(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/records/")
	var v domain.Record
	var e error
	switch r.Method {
	case "POST":
		if strings.HasSuffix(r.URL.Path, "/review") {
			id = strings.TrimSuffix(id, "/review")
			v, e = s.Svc.Review(id)
		} else if strings.HasSuffix(r.URL.Path, "/archive") {
			id = strings.TrimSuffix(id, "/archive")
			v, e = s.Svc.Archive(id)
		} else {
			v, e = s.Svc.Publish(id)
		}
	default:
		v, e = query.ByID(s.Svc.Store, id)
	}
	if e != nil {
		write(w, map[string]string{"error": e.Error()}, 400)
		return
	}
	write(w, v, 200)
}
func (s Server) importRows(w http.ResponseWriter, r *http.Request) {
	var data []byte
	data = make([]byte, r.ContentLength)
	r.Body.Read(data)
	v, e := report.Import(s.Svc, data)
	if e != nil {
		write(w, map[string]string{"error": e.Error()}, 400)
		return
	}
	write(w, v, 200)
}
func (s Server) scan(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/scan/")
	v, e := s.Svc.Scan(id)
	if e != nil {
		write(w, map[string]string{"error": e.Error()}, 400)
		return
	}
	write(w, v, 200)
}
