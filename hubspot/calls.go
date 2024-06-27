package hubspot

import (
	"context"
	"fmt"
)

type Calls interface {
	ListAssociations(ctx context.Context, query *CallAssociationsQuery, callId string, toObjectType string) (*CallAssociations, error)
	Associate(ctx context.Context, callId string, toObjectType string, toObjectId string, associationType string) (*Call, error)
	Disassociate(ctx context.Context, callId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *CallListQuery) (*CallList, error)
	Create(ctx context.Context, options *CallCreateOrUpdateOptions) (*Call, error)
	Read(ctx context.Context, query *CallReadQuery, callId string) (*Call, error)
	Update(ctx context.Context, options *CallCreateOrUpdateOptions, callId string) (*Call, error)
	Archive(ctx context.Context, callId string) error
	BatchArchive(ctx context.Context, callIds []string) error
	BatchCreate(ctx context.Context, options *CallBatchCreateOptions) (*CallBatchOutput, error)
	BatchRead(ctx context.Context, options *CallBatchReadOptions) (*CallBatchOutput, error)
	BatchUpdate(ctx context.Context, options *CallBatchUpdateOptions) (*CallBatchOutput, error)
	Search(ctx context.Context, options *CallSearchOptions) (*CallSearchResults, error)
	Merge(ctx context.Context, options *CallMergeOptions) (*Call, error)
}

type calls struct {
	client *Client
}

type CallAssociationsQuery struct {
	ListAssociationsQuery
}

type CallAssociations struct {
	Results []CallAssociation `json:"results"`
	Pagination
}

type CallAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CallListQuery struct {
	ListQuery
}

type CallList struct {
	Calls []Call `json:"results"`
	Pagination
}

type Call struct {
	Id         string         `json:"id"`
	Properties CallProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type CallProperties struct {
	CreateDate         string `json:"createdate"`
	HsCallBody         string `json:"hs_call_body,omitempty"`
	HsCallDuration     string `json:"hs_call_duration,omitempty"`
	HsCallFromNumber   string `json:"hs_call_from_number,omitempty"`
	HsCallRecordingUrl string `json:"hs_call_recording_url,omitempty"`
	HsCallStatus       string `json:"hs_call_status,omitempty"`
	HsCallTitle        string `json:"hs_call_title,omitempty"`
	HsCallToNumber     string `json:"hs_call_to_number,omitempty"`
	HsTimestamp        string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId     string `json:"hubspot_owner_id,omitempty"`
}

type CallCreateOrUpdateOptions struct {
	Properties CallCreateOrUpdateProperties `json:"properties"`
}

type CallCreateOrUpdateProperties struct {
	HsCallBody         string `json:"hs_call_body,omitempty"`
	HsCallDuration     string `json:"hs_call_duration,omitempty"`
	HsCallFromNumber   string `json:"hs_call_from_number,omitempty"`
	HsCallRecordingUrl string `json:"hs_call_recording_url,omitempty"`
	HsCallStatus       string `json:"hs_call_status,omitempty"`
	HsCallTitle        string `json:"hs_call_title,omitempty"`
	HsCallToNumber     string `json:"hs_call_to_number,omitempty"`
	HsTimestamp        string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId     string `json:"hubspot_owner_id,omitempty"`
}

type CallReadQuery struct {
	ReadQuery
}

type CallUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type CallBatchOutput struct {
	Status      string `json:"status"`
	Results     []Call `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type CallBatchReadOptions struct {
	BatchReadOptions
}

type CallBatchCreateOptions struct {
	Inputs []CallCreateOrUpdateOptions `json:"inputs"`
}

type CallBatchUpdateOptions struct {
	Inputs []CallBatchUpdateProperties `json:"inputs"`
}

type CallBatchUpdateProperties struct {
	Id         string                       `json:"id"`
	Properties CallCreateOrUpdateProperties `json:"properties"`
}

type CallSearchOptions struct {
	SearchOptions
}

type CallSearchResults struct {
	Total   int64  `json:"total"`
	Results []Call `json:"results"`
	Pagination
}

type CallMergeOptions struct {
	MergeOptions
}

func (z *calls) ListAssociations(ctx context.Context, query *CallAssociationsQuery, callId string, toObjectType string) (*CallAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/%s/associations/%s", callId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ca := &CallAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, err
	}
	return ca, nil
}

func (z *calls) Associate(ctx context.Context, callId string, toObjectType string, toObjectId string, associationType string) (*Call, error) {
	u := fmt.Sprintf("/crm/v3/objects/calls/%s/associations/%s/%s/%s", callId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, err
	}
	return call, nil
}

func (z *calls) Disassociate(ctx context.Context, callId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/calls/%s/associations/%s/%s/%s", callId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *calls) List(ctx context.Context, query *CallListQuery) (*CallList, error) {
	u := "/crm/v3/objects/calls"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	cl := &CallList{}

	err = z.client.do(req, cl)
	if err != nil {
		return nil, err
	}
	return cl, nil
}

func (z *calls) Create(ctx context.Context, options *CallCreateOrUpdateOptions) (*Call, error) {
	u := "/crm/v3/objects/calls"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, err
	}
	return call, nil
}

func (z *calls) Read(ctx context.Context, query *CallReadQuery, callId string) (*Call, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", callId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, err
	}
	return call, nil
}

func (z *calls) Update(ctx context.Context, options *CallCreateOrUpdateOptions, callId string) (*Call, error) {
	u := fmt.Sprintf("crm/v3/objects/calls/%s", callId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, err
	}
	return call, nil
}

func (z *calls) Archive(ctx context.Context, callId string) error {
	u := fmt.Sprintf("crm/v3/objects/calls/%s", callId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *calls) BatchArchive(ctx context.Context, callIds []string) error {
	u := "/crm/v3/objects/calls/batch/archive"
	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, callId := range callIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: callId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *calls) BatchCreate(ctx context.Context, options *CallBatchCreateOptions) (*CallBatchOutput, error) {
	u := "/crm/v3/objects/calls/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	calls := &CallBatchOutput{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, err
	}
	return calls, nil
}

func (z *calls) BatchRead(ctx context.Context, options *CallBatchReadOptions) (*CallBatchOutput, error) {
	u := "/crm/v3/objects/calls/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	calls := &CallBatchOutput{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, err
	}
	return calls, nil
}

func (z *calls) BatchUpdate(ctx context.Context, options *CallBatchUpdateOptions) (*CallBatchOutput, error) {
	u := "/crm/v3/objects/calls/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	calls := &CallBatchOutput{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, err
	}
	return calls, nil
}

func (z *calls) Search(ctx context.Context, options *CallSearchOptions) (*CallSearchResults, error) {
	u := "/crm/v3/objects/calls/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	calls := &CallSearchResults{}

	err = z.client.do(req, calls)
	if err != nil {
		return nil, err
	}
	return calls, nil
}

func (z *calls) Merge(ctx context.Context, options *CallMergeOptions) (*Call, error) {
	u := "/crm/v3/objects/calls/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	call := &Call{}

	err = z.client.do(req, call)
	if err != nil {
		return nil, err
	}
	return call, nil
}
