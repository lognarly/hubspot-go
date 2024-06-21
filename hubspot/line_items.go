package hubspot

import (
	"context"
	"fmt"
)

type LineItems interface {
	ListAssociations(ctx context.Context, query *LineItemAssociationsQuery, lineItem string, toObjectType string) (*LineItemAssociations, error)
	Associate(ctx context.Context, lineItemId string, toObjectType string, toObjectId string, associationType string) (*LineItem, error)
	Disassociate(ctx context.Context, lineItemId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *LineItemListQuery) (*LineItemList, error)
	Create(ctx context.Context, options *LineItemCreateOrUpdateOptions) (*LineItem, error)
	Read(ctx context.Context, query *LineItemReadQuery, lineItemId string) (*LineItem, error)
	Update(ctx context.Context, lineItemId string, options *LineItemCreateOrUpdateOptions) (*LineItem, error)
	Archive(ctx context.Context, lineItemId string) error
	BatchArchive(ctx context.Context, lineItemIds []string) error
	BatchCreate(ctx context.Context, options *LineItemBatchCreateOptions) (*LineItemBatchOutput, error)
	BatchRead(ctx context.Context, options *LineItemBatchReadOptions) (*LineItemBatchOutput, error)
	BatchUpdate(ctx context.Context, options *LineItemBatchUpdateOptions) (*LineItemBatchOutput, error)
	Search(ctx context.Context, options *LineItemSearchOptions) (*LineItemSearchResults, error)
	Merge(ctx context.Context, options *LineItemMergeOptions) (*LineItem, error)
}

type lineItems struct {
	client *Client
}

type LineItemAssociationsQuery struct {
	ListAssociationsQuery
}

type LineItemAssociations struct {
	Results []LineItemAssociation `json:"results"`
	Pagination
}

type LineItemAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type LineItemListQuery struct {
	ListQuery
}

type LineItemList struct {
	LineItems []LineItem `json:"results"`
	Pagination
}

type LineItem struct {
	Id         string             `json:"id"`
	Properties LineItemProperties `json:"properties"`
	CreatedAt  string             `json:"createdAt"`
	UpdatedAt  string             `json:"updatedAt"`
	Archived   bool               `json:"archived"`
}

type LineItemCreateOrUpdateOptions struct {
	Properties LineItemProperties `json:"properties"`
}

type LineItemReadQuery struct {
	ReadQuery
}

type LineItemBatchOutput struct {
	Status      string     `json:"status"`
	Results     []LineItem `json:"results"`
	RequestedAt string     `json:"requestedAt"`
	StartedAt   string     `json:"startedAt"`
	CompletedAt string     `json:"completedAt"`
}

type LineItemBatchReadOptions struct {
	BatchReadOptions
}

type LineItemBatchCreateOptions struct {
	Inputs []LineItemCreateOrUpdateOptions `json:"inputs"`
}

type LineItemBatchUpdateOptions struct {
	Inputs []LineItemBatchUpdateProperties `json:"inputs"`
}

type LineItemBatchUpdateProperties struct {
	Id         string             `json:"id"`
	Properties LineItemProperties `json:"properties"`
}

type LineItemSearchOptions struct {
	SearchOptions
}

type LineItemSearchResults struct {
	Total   int64      `json:"total"`
	Results []LineItem `json:"results"`
	Pagination
}

type LineItemMergeOptions struct {
	MergeOptions
}

func (z *lineItems) ListAssociations(ctx context.Context, query *LineItemAssociationsQuery, lineItemId string, toObjectType string) (*LineItemAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/%s/associations/%s", lineItemId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.ListAssociations(): newHttpRequest(ctx, ): %v", err)
	}

	la := &LineItemAssociations{}

	err = z.client.do(req, la)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.ListAssociations(): do(): %v", err)
	}

	return la, nil
}

func (z *lineItems) Associate(ctx context.Context, lineItemId string, toObjectType string, toObjectId string, associationType string) (*LineItem, error) {
	u := fmt.Sprintf("/crm/v3/objects/line_items/%s/associations/%s/%s/%s", lineItemId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Associate(): newHttpRequest(ctx, ): %v", err)
	}

	lineItem := &LineItem{}

	err = z.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Associate(): do(): %v", err)
	}

	return lineItem, nil
}

func (z *lineItems) Disassociate(ctx context.Context, lineItemId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/lineItems/%s/associations/%s/%s/%s", lineItemId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.lineItems.Disassociate(): newHttpRequest(ctx, ): %v", err)
	}

	return z.client.do(req, nil)
}

func (l *lineItems) List(ctx context.Context, query *LineItemListQuery) (*LineItemList, error) {
	u := "crm/v3/objects/line_items"
	req, err := l.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.List(): newHttpRequest(ctx, ): %v", err)
	}

	pl := &LineItemList{}

	err = l.client.do(req, pl)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.List(): do(): %v", err)
	}

	return pl, nil
}

func (l *lineItems) Create(ctx context.Context, options *LineItemCreateOrUpdateOptions) (*LineItem, error) {
	u := "/crm/v3/objects/line_items"
	req, err := l.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Create(): newHttpRequest(ctx, ): %+v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Create(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Read(ctx context.Context, query *LineItemReadQuery, lineItemId string) (*LineItem, error) {
	u := fmt.Sprintf("crm/v3/objects/line_items/%s", lineItemId)
	req, err := l.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Read(): newHttpRequest(ctx, ): %v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Read(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Update(ctx context.Context, lineItemId string, options *LineItemCreateOrUpdateOptions) (*LineItem, error) {
	u := fmt.Sprintf("crm/v3/objects/line_items/%s", lineItemId)
	req, err := l.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Update(): newHttpRequest(ctx, ): %v", err)
	}

	lineItem := &LineItem{}

	err = l.client.do(req, lineItem)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.Update(): do(): %+v", err)
	}

	return lineItem, nil
}

func (l *lineItems) Archive(ctx context.Context, lineItemId string) error {
	u := fmt.Sprintf("crm/v3/objects/line_items/%s", lineItemId)
	req, err := l.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return fmt.Errorf("client.lineItems.Archive(): newHttpRequest(ctx, ): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchArchive(ctx context.Context, lineItemIds []string) error {
	u := "/crm/v3/objects/line_items/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, lineItemId := range lineItemIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: lineItemId})
	}

	req, err := l.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return fmt.Errorf("client.lineItems.BatchArchive(): newHttpRequest(ctx, ): %v", err)
	}

	return l.client.do(req, nil)
}

func (l *lineItems) BatchCreate(ctx context.Context, options *LineItemBatchCreateOptions) (*LineItemBatchOutput, error) {
	u := "/crm/v3/objects/line_items/batch/create"
	req, err := l.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchCreate(): newHttpRequest(ctx, ): %v", err)
	}

	lineItems := &LineItemBatchOutput{}

	err = l.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItem.BatchCreate(): do(): %+v", err)
	}

	return lineItems, nil
}

func (z *lineItems) BatchRead(ctx context.Context, options *LineItemBatchReadOptions) (*LineItemBatchOutput, error) {
	u := "/crm/v3/objects/line_items/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchRead(): newHttpRequest(ctx, ): %v", err)
	}

	lbrr := &LineItemBatchOutput{}

	err = z.client.do(req, lbrr)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchRead(): do(): %+v", err)
	}

	return lbrr, nil
}

func (z *lineItems) BatchUpdate(ctx context.Context, options *LineItemBatchUpdateOptions) (*LineItemBatchOutput, error) {
	u := "/crm/v3/objects/line_items/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchUpdate(): newHttpRequest(ctx, ): %v", err)
	}

	lineItems := &LineItemBatchOutput{}

	err = z.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.BatchUpdate(): do(): %+v", err)
	}

	return lineItems, nil
}

func (z *lineItems) Search(ctx context.Context, options *LineItemSearchOptions) (*LineItemSearchResults, error) {
	u := "/crm/v3/objects/line_items/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Search(): newHttpRequest(ctx, ): %v", err)
	}

	lineItems := &LineItemSearchResults{}

	err = z.client.do(req, lineItems)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Search(): do(): %+v", err)
	}

	return lineItems, nil
}

func (z *lineItems) Merge(ctx context.Context, options *LineItemMergeOptions) (*LineItem, error) {
	u := "/crm/v3/objects/line_items/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Merge(): newHttpRequest(ctx, ): %v", err)
	}

	company := &LineItem{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, fmt.Errorf("client.lineItems.Merge(): do(): %+v", err)
	}

	return company, nil
}
