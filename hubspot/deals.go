package hubspot

import (
	"context"
	"fmt"
)

type Deals interface {
	ListAssociations(ctx context.Context, query *DealAssociationsQuery, dealId string, toObjectType string) (*DealAssociations, error)
	Associate(ctx context.Context, dealId string, toObjectType string, toObjectId string, associationType string) (*Deal, error)
	Disassociate(ctx context.Context, dealId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *DealListQuery) (*DealList, error)
	Create(ctx context.Context, options *DealCreateOrUpdateOptions) (*Deal, error)
	Read(ctx context.Context, query *DealReadQuery, dealId string) (*Deal, error)
	Update(ctx context.Context, dealId string, options *DealCreateOrUpdateOptions) (*Deal, error)
	Archive(ctx context.Context, dealId string) error
	BatchArchive(ctx context.Context, dealIds []string) error
	BatchCreate(ctx context.Context, options *DealBatchCreateOptions) (*DealBatchOutput, error)
	BatchRead(ctx context.Context, options *DealBatchReadOptions) (*DealBatchOutput, error)
	BatchUpdate(ctx context.Context, options *DealBatchUpdateOptions) (*DealBatchOutput, error)
	Search(ctx context.Context, options *DealSearchOptions) (*DealSearchResults, error)
	Merge(ctx context.Context, options *DealMergeOptions) (*Deal, error)
}

type deals struct {
	client *Client
}

type DealAssociationsQuery struct {
	ListAssociationsQuery
}

type DealAssociations struct {
	Results []DealAssociation `json:"results"`
	Pagination
}

type DealAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type DealListQuery struct {
	ListQuery
}

type DealList struct {
	Deals []Deal `json:"results"`
	Pagination
}

type Deal struct {
	Id         string         `json:"id"`
	Properties DealProperties `json:"properties"` //This can be found in deal_properties.go as to not clutter this file
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type DealCreateOrUpdateOptions struct {
	Properties   DealProperties `json:"properties"`
	Associations []Association  `json:"associations,omitempty"`
}

type DealReadQuery struct {
	ReadQuery
}

type DealBatchOutput struct {
	Status      string `json:"status"`
	Results     []Deal `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type DealBatchReadOptions struct {
	BatchReadOptions
}

type DealBatchCreateOptions struct {
	Inputs []DealCreateOrUpdateOptions `json:"inputs"`
}

type DealBatchUpdateOptions struct {
	Inputs []DealBatchUpdateProperties `json:"inputs"`
}

type DealBatchUpdateProperties struct {
	Id         string         `json:"id"`
	Properties DealProperties `json:"properties"` //This can be found in deal_properties.go as to not clutter this file
}

type DealSearchOptions struct {
	SearchOptions
}

type DealSearchResults struct {
	Total   int64  `json:"total"`
	Results []Deal `json:"results"`
	Pagination
}

type DealMergeOptions struct {
	MergeOptions
}

func (z *deals) ListAssociations(ctx context.Context, query *DealAssociationsQuery, dealId string, toObjectType string) (*DealAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/%s/associations/%s", dealId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	da := &DealAssociations{}
	err = z.client.do(req, da)
	if err != nil {
		return nil, err
	}
	return da, nil
}

func (z *deals) Associate(ctx context.Context, dealId string, toObjectType string, toObjectId string, associationType string) (*Deal, error) {
	u := fmt.Sprintf("/crm/v3/objects/deals/%s/associations/%s/%s/%s", dealId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	deal := &Deal{}
	err = z.client.do(req, deal)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (z *deals) Disassociate(ctx context.Context, dealId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/deals/%s/associations/%s/%s/%s", dealId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *deals) List(ctx context.Context, query *DealListQuery) (*DealList, error) {
	u := "crm/v3/objects/deals"

	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	dl := &DealList{}
	err = z.client.do(req, dl)
	if err != nil {
		return nil, err
	}
	return dl, nil
}

func (z *deals) Create(ctx context.Context, options *DealCreateOrUpdateOptions) (*Deal, error) {
	u := "crm/v3/objects/deals"

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	deal := &Deal{}
	err = z.client.do(req, deal)
	if err != nil {
		return nil, err
	}

	return deal, nil
}

func (z *deals) Read(ctx context.Context, query *DealReadQuery, dealId string) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)

	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	deal := &Deal{}
	err = z.client.do(req, deal)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (z *deals) Update(ctx context.Context, dealId string, options *DealCreateOrUpdateOptions) (*Deal, error) {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	deal := &Deal{}
	err = z.client.do(req, deal)
	if err != nil {
		return nil, err
	}
	return deal, nil
}

func (z *deals) Archive(ctx context.Context, dealId string) error {
	u := fmt.Sprintf("crm/v3/objects/deals/%s", dealId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *deals) BatchArchive(ctx context.Context, dealIds []string) error {
	u := "/crm/v3/objects/deals/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, dealId := range dealIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: dealId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *deals) BatchCreate(ctx context.Context, options *DealBatchCreateOptions) (*DealBatchOutput, error) {
	u := "/crm/v3/objects/deals/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	deals := &DealBatchOutput{}
	err = z.client.do(req, deals)
	if err != nil {
		return nil, err
	}
	return deals, nil
}

func (z *deals) BatchRead(ctx context.Context, options *DealBatchReadOptions) (*DealBatchOutput, error) {
	u := "/crm/v3/objects/deals/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	deals := &DealBatchOutput{}
	err = z.client.do(req, deals)
	if err != nil {
		return nil, err
	}
	return deals, nil
}

func (z *deals) BatchUpdate(ctx context.Context, options *DealBatchUpdateOptions) (*DealBatchOutput, error) {
	u := "/crm/v3/objects/deals/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	deals := &DealBatchOutput{}
	err = z.client.do(req, deals)
	if err != nil {
		return nil, err
	}
	return deals, nil
}

func (z *deals) Search(ctx context.Context, options *DealSearchOptions) (*DealSearchResults, error) {
	u := "/crm/v3/objects/deals/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	deals := &DealSearchResults{}
	err = z.client.do(req, deals)
	if err != nil {
		return nil, err
	}
	return deals, nil
}

func (z *deals) Merge(ctx context.Context, options *DealMergeOptions) (*Deal, error) {
	u := "/crm/v3/objects/deals/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	deal := &Deal{}
	err = z.client.do(req, deal)
	if err != nil {
		return nil, err
	}
	return deal, nil
}
