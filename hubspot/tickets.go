package hubspot

import (
	"context"
	"fmt"
)

type Tickets interface {
	ListAssociations(ctx context.Context, query *TicketAssociationsQuery, ticketId string, toObjectType string) (*TicketAssociations, error)
	Associate(ctx context.Context, ticketId string, toObjectType string, toObjectId string, associationType string) (*Ticket, error)
	Disassociate(ctx context.Context, ticketId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *TicketsListQuery) (*TicketList, error)
	Create(ctx context.Context, options *TicketCreateOrUpdateOptions) (*Ticket, error)
	Read(ctx context.Context, ticketId string, query *TicketReadQuery) (*Ticket, error)
	Update(ctx context.Context, ticketId string, options *TicketCreateOrUpdateOptions) (*Ticket, error)
	Archive(ctx context.Context, ticketId string) error
	BatchArchive(ctx context.Context, ticketIds []string) error
	BatchCreate(ctx context.Context, options *TicketBatchCreateOptions) (*TicketBatchOutput, error)
	BatchRead(ctx context.Context, options *TicketBatchReadOptions) (*TicketBatchOutput, error)
	BatchUpdate(ctx context.Context, options *TicketBatchUpdateOptions) (*TicketBatchOutput, error)
	Search(ctx context.Context, options *TicketSearchOptions) (*TicketSearchResults, error)
	Merge(ctx context.Context, options *MergeOptions) (*Ticket, error)
}

type tickets struct {
	client *Client
}

type TicketAssociationsQuery struct {
	ListAssociationsQuery
}

type TicketAssociations struct {
	Results []TicketAssociation `json:"results"`
	Pagination
}

type TicketAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type TicketsListQuery struct {
	ListQuery
}

type TicketList struct {
	Results []Ticket `json:"results"`
	Pagination
}

type Ticket struct {
	Id         string           `json:"id"`
	Properties TicketProperties `json:"properties"`
}

type TicketCreateOrUpdateOptions struct {
	Properties TicketProperties `json:"properties"`
}

type TicketReadQuery struct {
	ReadQuery
}

type TicketBatchOutput struct {
	Status      string   `json:"status"`
	Results     []Ticket `json:"results"`
	RequestedAt string   `json:"requestedAt"`
	StartedAt   string   `json:"startedAt"`
	CompletedAt string   `json:"completedAt"`
}

type TicketBatchReadOptions struct {
	BatchReadOptions
}

type TicketBatchCreateOptions struct {
	Inputs []TicketCreateOrUpdateOptions `json:"inputs"`
}

type TicketBatchUpdateOptions struct {
	Inputs []TicketBatchUpdateProperties `json:"inputs"`
}

type TicketBatchUpdateProperties struct {
	Id         string           `json:"id"`
	Properties TicketProperties `json:"properties"`
}

type TicketSearchOptions struct {
	SearchOptions
}

type TicketSearchResults struct {
	Total   int64    `json:"total"`
	Results []Ticket `json:"results"`
}

type TicketMergeOptions struct {
	MergeOptions
}

func (z *tickets) ListAssociations(ctx context.Context, query *TicketAssociationsQuery, ticketId string, toObjectType string) (*TicketAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s/associations/%s", ticketId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ta := &TicketAssociations{}

	err = z.client.do(req, ta)
	if err != nil {
		return nil, err
	}

	return ta, nil
}

func (z *tickets) Associate(ctx context.Context, ticketId string, toObjectType string, toObjectId string, associationType string) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s/associations/%s/%s/%s", ticketId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (z *tickets) Disassociate(ctx context.Context, ticketId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s/associations/%s/%s/%s", ticketId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *tickets) List(ctx context.Context, query *TicketsListQuery) (*TicketList, error) {
	u := "/crm/v3/objects/tickets"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	tl := &TicketList{}

	err = z.client.do(req, tl)
	if err != nil {
		return nil, err
	}

	return tl, nil
}

func (z *tickets) Create(ctx context.Context, options *TicketCreateOrUpdateOptions) (*Ticket, error) {
	u := "/crm/v3/objects/tickets"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (z *tickets) Read(ctx context.Context, ticketId string, query *TicketReadQuery) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s", ticketId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (z *tickets) Update(ctx context.Context, ticketId string, options *TicketCreateOrUpdateOptions) (*Ticket, error) {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s", ticketId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (z *tickets) Archive(ctx context.Context, ticketId string) error {
	u := fmt.Sprintf("/crm/v3/objects/tickets/%s", ticketId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *tickets) BatchArchive(ctx context.Context, ticketIds []string) error {
	u := "/crm/v3/objects/tickets/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, ticketId := range ticketIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: ticketId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *tickets) BatchCreate(ctx context.Context, options *TicketBatchCreateOptions) (*TicketBatchOutput, error) {
	u := "/crm/v3/objects/tickets/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tbr := &TicketBatchOutput{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, err
	}

	return tbr, nil
}

func (z *tickets) BatchRead(ctx context.Context, options *TicketBatchReadOptions) (*TicketBatchOutput, error) {
	u := "/crm/v3/objects/tickets/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tbr := &TicketBatchOutput{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, err
	}

	return tbr, nil
}

func (z *tickets) BatchUpdate(ctx context.Context, options *TicketBatchUpdateOptions) (*TicketBatchOutput, error) {
	u := "/crm/v3/objects/tickets/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tbr := &TicketBatchOutput{}

	err = z.client.do(req, tbr)
	if err != nil {
		return nil, err
	}

	return tbr, nil
}

func (z *tickets) Search(ctx context.Context, options *TicketSearchOptions) (*TicketSearchResults, error) {
	u := "/crm/v3/objects/tickets/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tsr := &TicketSearchResults{}

	err = z.client.do(req, tsr)
	if err != nil {
		return nil, err
	}

	return tsr, nil
}

func (z *tickets) Merge(ctx context.Context, options *MergeOptions) (*Ticket, error) {
	u := "/crm/v3/objects/tickets/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	ticket := &Ticket{}

	err = z.client.do(req, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}
