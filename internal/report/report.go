package report

import (
	"encoding/json"
	"fmt"
	"warehouse5s/internal/domain"
	"warehouse5s/internal/service"
)

type Row struct {
	ID        string        `json:"id"`
	Site      string        `json:"site"`
	Inspector string        `json:"inspector"`
	Items     []domain.Item `json:"items"`
}
type Result struct {
	Accepted int
	Rejected int
	Errors   []string
	Records  []domain.Record
}

func Import(s *service.Service, data []byte) (Result, error) {
	var rows []Row
	if e := json.Unmarshal(data, &rows); e != nil {
		return Result{}, e
	}
	res := Result{}
	for n, row := range rows {
		if row.ID == "" || row.Site == "" || row.Inspector == "" {
			res.Rejected++
			res.Errors = append(res.Errors, fmt.Sprintf("row %d invalid", n))
			continue
		}
		r, e := s.Create(row.ID, row.Site, row.Inspector, row.Items)
		if e != nil {
			res.Rejected++
			res.Errors = append(res.Errors, e.Error())
			continue
		}
		res.Accepted++
		res.Records = append(res.Records, r)
	}
	return res, nil
}
func Render(res Result) string {
	return fmt.Sprintf("accepted=%d rejected=%d", res.Accepted, res.Rejected)
}
func ValidateRow(r Row) error {
	if r.ID == "" {
		return fmt.Errorf("id required")
	}
	if len(r.Items) == 0 {
		return fmt.Errorf("items required")
	}
	return nil
}
func Export(records []domain.Record) []byte { b, _ := json.Marshal(records); return b }
