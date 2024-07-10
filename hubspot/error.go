package hubspot

import (
	"encoding/json"
	"fmt"
)

type HubspotErrorResponse struct {
	SubCategory   string                 `json:"subCategory,omitempty"`
	Context       map[string]interface{} `json:"context,omitempty"`
	CorrelationId string                 `json:"correlationId,omitempty"`
	Links         map[string]string      `json:"links,omitempty"`
	Message       string                 `json:"message,omitempty"`
	Category      string                 `json:"category,omitempty"`
	Errors        []HubspotError         `json:"errors,omitempty"`
	Status        string                 `json:"status,omitempty"`
}

type HubspotError struct {
	SubCategory string                 `json:"subCategory,omitempty"`
	Code        string                 `json:"code,omitempty"`
	In          string                 `json:"in,omitempty"`
	Context     map[string]interface{} `json:"context,omitempty"`
	Message     string                 `json:"message,omitempty"`
}

func (e *HubspotErrorResponse) Error() string {
	j, err := json.Marshal(e)
	if err != nil {
		return fmt.Sprintf("failed to marshal HubspotErrorResponse: %v", err)
	}
	return string(j)
}

func GetErrorResponseFromError(respErr error) (*HubspotErrorResponse, error) {
	errResponse := &HubspotErrorResponse{}
	err := json.Unmarshal([]byte(respErr.Error()), &errResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal HubspotErrorResponse: %v, original error: %v", err, respErr)
	}
	return errResponse, nil
}
