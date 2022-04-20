package main

import (
	"encoding/json"
	"time"

	"github.com/leplasmo/todo-api"
)

type todoResponse struct {
	Results todo.List `json:"results"`
}

// customize json output by creating our own MarshalJSON method
func (r *todoResponse) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results      todo.List `json:"results"`
		Date         int64     `json:"date"`
		TotalResults int       `json:"total_results"`
	}{
		Results:      r.Results,
		Date:         time.Now().Unix(),
		TotalResults: len(r.Results),
	}
	return json.Marshal(resp)
}
