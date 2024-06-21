package hubspot

import ()

type ListAssociationsQuery struct {
	Limit int    `url:"limit,omitempty"`
	After string `url:"after,omitempty"`
}

type ListQuery struct {
	Limit                 int32    `url:"limit,omitempty"`
	After                 string   `url:"after,omitempty"`
	Properties            []string `url:"properties,omitempty"`
	PropertiesWithHistory []string `url:"propertiesWithHistory,omitempty"`
	Associations          []string `url:"associations,omitempty"`
	Archived              bool     `url:"archived,omitempty"`
}

type ReadQuery struct {
	Properties            []string `url:"properties,omitempty"`
	PropertiesWithHistory []string `url:"propertiesWithHistory,omitempty"`
	Associations          []string `url:"associations,omitempty"`
	Archived              bool     `url:"archived,omitempty"`
	IdProperty            string   `url:"idProperty,omitempty"`
}

type Pagination struct {
	Paging Paging `json:"paging"`
}

type Paging struct {
	Next Next `json:"next"`
}

type Next struct {
	After string `json:"after"`
}

type HubSpotError struct {
	Status        string `json:"status"`
	Message       string `json:"message"`
	CorrelationId string `json:"correlationId"`
	Category      string `json:"category"`
}

type FilterOperator string

const (
	EqualTo            FilterOperator = "EQ"
	NotEqualTo         FilterOperator = "NEQ"
	LessThan           FilterOperator = "LT"
	LessThanEqualTo    FilterOperator = "LTE"
	GreaterThan        FilterOperator = "GT"
	GreaterThanEqualTo FilterOperator = "GTE"
	Between            FilterOperator = "BETWEEN"
	In                 FilterOperator = "IN"
	NotIn              FilterOperator = "NOT_IN"
	HasProperty        FilterOperator = "HAS_PROPERTY"
	NotHasProperty     FilterOperator = "NOT_HAS_PROPERTY"
	ContainsToken      FilterOperator = "CONTAINS_TOKEN"
	NotContainsToken   FilterOperator = "NOT_CONTAINS_TOKEN"
)

type SearchOptions struct {
	FilterGroups []FilterGroups `json:"filterGroups,omitempty"`
	Sorts        []string       `json:"sorts,omitempty"`
	Query        string         `json:"query,omitempty"`
	Properties   []string       `json:"properties,omitempty"`
	Limit        int32          `json:"limit,omitempty"`
	After        int32          `json:"after,omitempty"`
}

type FilterGroups struct {
	Filters []Filters `json:"filters"`
}

type Filters struct {
	Value        string         `json:"value,omitempty"`
	Values       []string       `json:"values,omitempty"`
	PropertyName string         `json:"propertyName,omitempty"`
	Operator     FilterOperator `json:"operator,omitempty"`
}

type BatchReadQuery struct {
	Archived bool `url:"archived,omitempty"`
}

type BatchInputOptions struct {
	Inputs []BatchInput `json:"inputs,omitempty"`
}

type BatchInput struct {
	Id string `json:"id"`
}

type BatchReadOptions struct {
	Properties            []string `json:"properties,omitempty"`
	PropertiesWithHistory []string `json:"propertiesWithHistory,omitempty"`
	IdProperty            string   `json:"idProperty,omitempty"`
	BatchInputOptions
}

type MergeOptions struct {
	PrimaryObjectId string `json:"primaryObjectId"`
	ObjectIdToMerge string `json:"objectIdToMerge"`
}
