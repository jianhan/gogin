package error

import "fmt"

type APIError struct {
	Status  int                    `json:"status,omitempty"`
	Details string                 `json:"details,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

func (a *APIError) Error() string {
	return fmt.Sprintf("APIError: %d, %s", a.Status, a.Details)
}

func (a *APIError) AddData(k string, v interface{}) {
	if a.Data == nil {
		a.Data = map[string]interface{}{}
	}
	a.Data[k] = v
}
