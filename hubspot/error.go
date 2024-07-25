package hubspot

import (
	"encoding/json"
	"fmt"
)

type ErrorResponse struct {
	SubCategory   string                 `json:"subCategory,omitempty"`
	Context       map[string]interface{} `json:"context,omitempty"`
	CorrelationId string                 `json:"correlationId,omitempty"`
	Links         map[string]string      `json:"links,omitempty"`
	Message       string                 `json:"message,omitempty"`
	Category      string                 `json:"category,omitempty"`
	Errors        []ErrorObject          `json:"errors,omitempty"`
	Status        string                 `json:"status,omitempty"`
}

type ErrorObject struct {
	SubCategory string                 `json:"subCategory,omitempty"`
	Code        string                 `json:"code,omitempty"`
	In          string                 `json:"in,omitempty"`
	Context     map[string]interface{} `json:"context,omitempty"`
	Message     string                 `json:"message,omitempty"`
}

func (e *ErrorResponse) Error() string {
	j, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("failed to marshal ErrorResponse: %v", err)
	}
	return string(j)
}

func GetErrorResponseFromError(respErr error) (*ErrorResponse, error) {
	errResponse := &ErrorResponse{}
	err := json.Unmarshal([]byte(respErr.Error()), &errResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal ErrorResponse: %v, original error: %v", err, respErr)
	}
	return errResponse, nil
}
