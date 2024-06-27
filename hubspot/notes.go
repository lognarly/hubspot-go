package hubspot

import (
	"context"
	"fmt"
)

type Notes interface {
	ListAssociations(ctx context.Context, query *NoteAssociationsQuery, noteId string, toObjectType string) (*NoteAssociations, error)
	Associate(ctx context.Context, noteId string, toObjectType string, toObjectId string, associationType string) (*Note, error)
	Disassociate(ctx context.Context, noteId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *NoteListQuery) (*NoteList, error)
	Create(ctx context.Context, options *NoteCreateOrUpdateOptions) (*Note, error)
	Read(ctx context.Context, query *NoteReadQuery, noteId string) (*Note, error)
	Update(ctx context.Context, options *NoteCreateOrUpdateOptions, noteId string) (*Note, error)
	Archive(ctx context.Context, noteId string) error
	BatchArchive(ctx context.Context, noteIds []string) error
	BatchCreate(ctx context.Context, options *NoteBatchCreateOptions) (*NoteBatchOutput, error)
	BatchRead(ctx context.Context, options *NoteBatchReadOptions) (*NoteBatchOutput, error)
	BatchUpdate(ctx context.Context, options *NoteBatchUpdateOptions) (*NoteBatchOutput, error)
	Search(ctx context.Context, options *NoteSearchOptions) (*NoteSearchResults, error)
	Merge(ctx context.Context, options *NoteMergeOptions) (*Note, error)
}

type notes struct {
	client *Client
}

type NoteAssociationsQuery struct {
	ListAssociationsQuery
}

type NoteAssociations struct {
	Results []NoteAssociation `json:"results"`
	Pagination
}

type NoteAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type NoteListQuery struct {
	ListQuery
}

type NoteList struct {
	Notes []Note `json:"results"`
	Pagination
}

type Note struct {
	Id         string         `json:"id"`
	Properties NoteProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type NoteProperties struct {
	CreateDate         string `json:"createdate"`
	HsLastModifiedDate string `json:"hs_lastmodifieddate"`
	HsNoteBody         string `json:"hs_note_body,omitempty"`
	HsTimestamp        string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId     string `json:"hubspot_owner_id,omitempty"`
}

type NoteCreateOrUpdateOptions struct {
	Properties NoteCreateOrUpdateProperties `json:"properties"`
}

type NoteCreateOrUpdateProperties struct {
	HsNoteBody     string `json:"hs_note_body,omitempty"`
	HsTimestamp    string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId string `json:"hubspot_owner_id,omitempty"`
}

type NoteReadQuery struct {
	ReadQuery
}

type NoteUpdateQuery struct {
	IdProperty string `url:"idProperty,omitempty"`
}

type NoteBatchOutput struct {
	Status      string `json:"status"`
	Results     []Note `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type NoteBatchReadOptions struct {
	BatchReadOptions
}

type NoteBatchCreateOptions struct {
	Inputs []NoteCreateOrUpdateOptions `json:"inputs"`
}

type NoteBatchUpdateOptions struct {
	Inputs []NoteBatchUpdateProperties `json:"inputs"`
}

type NoteBatchUpdateProperties struct {
	Id         string                       `json:"id"`
	Properties NoteCreateOrUpdateProperties `json:"properties"`
}

type NoteSearchOptions struct {
	SearchOptions
}

type NoteSearchResults struct {
	Total   int64  `json:"total"`
	Results []Note `json:"results"`
	Pagination
}

type NoteMergeOptions struct {
	MergeOptions
}

func (z *notes) ListAssociations(ctx context.Context, query *NoteAssociationsQuery, noteId string, toObjectType string) (*NoteAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/notes/%s/associations/%s", noteId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ca := &NoteAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, err
	}
	return ca, nil
}

func (z *notes) Associate(ctx context.Context, noteId string, toObjectType string, toObjectId string, associationType string) (*Note, error) {
	u := fmt.Sprintf("/crm/v3/objects/notes/%s/associations/%s/%s/%s", noteId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	note := &Note{}

	err = z.client.do(req, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (z *notes) Disassociate(ctx context.Context, noteId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/notes/%s/associations/%s/%s/%s", noteId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *notes) List(ctx context.Context, query *NoteListQuery) (*NoteList, error) {
	u := "/crm/v3/objects/notes"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	nl := &NoteList{}

	err = z.client.do(req, nl)
	if err != nil {
		return nil, err
	}
	return nl, nil
}

func (z *notes) Create(ctx context.Context, options *NoteCreateOrUpdateOptions) (*Note, error) {
	u := "/crm/v3/objects/notes"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	note := &Note{}

	err = z.client.do(req, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (z *notes) Read(ctx context.Context, query *NoteReadQuery, noteId string) (*Note, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", noteId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	note := &Note{}

	err = z.client.do(req, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (z *notes) Update(ctx context.Context, options *NoteCreateOrUpdateOptions, noteId string) (*Note, error) {
	u := fmt.Sprintf("crm/v3/objects/notes/%s", noteId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	note := &Note{}

	err = z.client.do(req, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}

func (z *notes) Archive(ctx context.Context, noteId string) error {
	u := fmt.Sprintf("crm/v3/objects/notes/%s", noteId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *notes) BatchArchive(ctx context.Context, noteIds []string) error {
	u := "/crm/v3/objects/notes/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, noteId := range noteIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: noteId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *notes) BatchCreate(ctx context.Context, options *NoteBatchCreateOptions) (*NoteBatchOutput, error) {
	u := "/crm/v3/objects/notes/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	notes := &NoteBatchOutput{}

	err = z.client.do(req, notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (z *notes) BatchRead(ctx context.Context, options *NoteBatchReadOptions) (*NoteBatchOutput, error) {
	u := "/crm/v3/objects/notes/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	notes := &NoteBatchOutput{}

	err = z.client.do(req, notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (z *notes) BatchUpdate(ctx context.Context, options *NoteBatchUpdateOptions) (*NoteBatchOutput, error) {
	u := "/crm/v3/objects/notes/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	notes := &NoteBatchOutput{}

	err = z.client.do(req, notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (z *notes) Search(ctx context.Context, options *NoteSearchOptions) (*NoteSearchResults, error) {
	u := "/crm/v3/objects/notes/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	notes := &NoteSearchResults{}

	err = z.client.do(req, notes)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (z *notes) Merge(ctx context.Context, options *NoteMergeOptions) (*Note, error) {
	u := "/crm/v3/objects/notes/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	note := &Note{}

	err = z.client.do(req, note)
	if err != nil {
		return nil, err
	}
	return note, nil
}
