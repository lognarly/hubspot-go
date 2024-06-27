package hubspot

import (
	"context"
	"fmt"
)

type Emails interface {
	ListAssociations(ctx context.Context, query *EmailAssociationsQuery, emailId string, toObjectType string) (*EmailAssociations, error)
	Associate(ctx context.Context, emailId string, toObjectType string, toObjectId string, associationType string) (*Email, error)
	Disassociate(ctx context.Context, emailId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *EmailListQuery) (*EmailList, error)
	Create(ctx context.Context, options *EmailCreateOrUpdateOptions) (*Email, error)
	Read(ctx context.Context, query *EmailReadQuery, emailId string) (*Email, error)
	Update(ctx context.Context, options *EmailCreateOrUpdateOptions, emailId string) (*Email, error)
	Archive(ctx context.Context, emailId string) error
	BatchArchive(ctx context.Context, emailIds []string) error
	BatchCreate(ctx context.Context, options *EmailBatchCreateOptions) (*EmailBatchOutput, error)
	BatchRead(ctx context.Context, options *EmailBatchReadOptions) (*EmailBatchOutput, error)
	BatchUpdate(ctx context.Context, options *EmailBatchUpdateOptions) (*EmailBatchOutput, error)
	Search(ctx context.Context, options *EmailSearchOptions) (*EmailSearchResults, error)
	Merge(ctx context.Context, options *EmailMergeOptions) (*Email, error)
}

type emails struct {
	client *Client
}

type EmailAssociationsQuery struct {
	ListAssociationsQuery
}

type EmailAssociations struct {
	Results []EmailAssociation `json:"results"`
	Pagination
}

type EmailAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type EmailListQuery struct {
	ListQuery
}

type EmailList struct {
	Emails []Email `json:"results"`
	Pagination
}

type Email struct {
	Id         string          `json:"id"`
	Properties EmailProperties `json:"properties"`
	CreatedAt  string          `json:"createdAt"`
	UpdatedAt  string          `json:"updatedAt"`
	Archived   bool            `json:"archived"`
}

type EmailProperties struct {
	CreateDate             string `json:"createdate"`
	HsEmailDirection       string `json:"hs_email_direction,omitempty"`
	HsEmailSenderEmail     string `json:"hs_email_sender_email,omitempty"`
	HsEmailSenderFirstName string `json:"hs_email_sender_firstname"`
	HsEmailSenderLastName  string `json:"hs_email_sender_lastname,omitempty"`
	HsEmailStatus          string `json:"hs_email_status,omitempty"`
	HsEmailSubject         string `json:"hs_email_subject,omitempty"`
	HsEmailText            string `json:"hs_email_text,omitempty"`
	HsEmailToEmail         string `json:"hs_email_to_email,omitempty"`
	HsEmailToFirstName     string `json:"hs_email_to_firstname,omitempty"`
	HsEmailToLastName      string `json:"hs_email_to_lastname,omitempty"`
	HsLastModifiedDate     string `json:"hs_lastmodifieddate,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type EmailCreateOrUpdateOptions struct {
	Properties EmailCreateOrUpdateProperties `json:"properties"`
}

type EmailCreateOrUpdateProperties struct {
	HsEmailDirection       string `json:"hs_email_direction,omitempty"`
	HsEmailSenderEmail     string `json:"hs_email_sender_email,omitempty"`
	HsEmailSenderFirstName string `json:"hs_email_sender_firstname,omitempty"`
	HsEmailSenderLastName  string `json:"hs_email_sender_lastname,omitempty"`
	HsEmailStatus          string `json:"hs_email_status,omitempty"`
	HsEmailSubject         string `json:"hs_email_subject,omitempty"`
	HsEmailText            string `json:"hs_email_text,omitempty"`
	HsEmailToEmail         string `json:"hs_email_to_email,omitempty"`
	HsEmailToFirstName     string `json:"hs_email_to_firstname,omitempty"`
	HsEmailToLastName      string `json:"hs_email_to_lastname,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type EmailReadQuery struct {
	ReadQuery
}

type EmailUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type EmailBatchOutput struct {
	Status      string  `json:"status"`
	Results     []Email `json:"results"`
	RequestedAt string  `json:"requestedAt"`
	StartedAt   string  `json:"startedAt"`
	CompletedAt string  `json:"completedAt"`
}

type EmailBatchReadOptions struct {
	BatchReadOptions
}

type EmailBatchCreateOptions struct {
	Inputs []EmailCreateOrUpdateOptions `json:"inputs"`
}

type EmailBatchUpdateOptions struct {
	Inputs []EmailBatchUpdateProperties `json:"inputs"`
}

type EmailBatchUpdateProperties struct {
	Id         string                        `json:"id"`
	Properties EmailCreateOrUpdateProperties `json:"properties"`
}

type EmailSearchOptions struct {
	SearchOptions
}

type EmailSearchResults struct {
	Total   int64   `json:"total"`
	Results []Email `json:"results"`
	Pagination
}

type EmailMergeOptions struct {
	MergeOptions
}

func (z *emails) ListAssociations(ctx context.Context, query *EmailAssociationsQuery, emailId string, toObjectType string) (*EmailAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/%s/associations/%s", emailId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ca := &EmailAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, err
	}
	return ca, nil
}

func (z *emails) Associate(ctx context.Context, emailId string, toObjectType string, toObjectId string, associationType string) (*Email, error) {
	u := fmt.Sprintf("/crm/v3/objects/emails/%s/associations/%s/%s/%s", emailId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (z *emails) Disassociate(ctx context.Context, emailId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/emails/%s/associations/%s/%s/%s", emailId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *emails) List(ctx context.Context, query *EmailListQuery) (*EmailList, error) {
	u := "/crm/v3/objects/emails"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	el := &EmailList{}

	err = z.client.do(req, el)
	if err != nil {
		return nil, err
	}
	return el, nil
}

func (z *emails) Create(ctx context.Context, options *EmailCreateOrUpdateOptions) (*Email, error) {
	u := "/crm/v3/objects/emails"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (z *emails) Read(ctx context.Context, query *EmailReadQuery, emailId string) (*Email, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", emailId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (z *emails) Update(ctx context.Context, options *EmailCreateOrUpdateOptions, emailId string) (*Email, error) {
	u := fmt.Sprintf("crm/v3/objects/emails/%s", emailId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, err
	}
	return email, nil
}

func (z *emails) Archive(ctx context.Context, emailId string) error {
	u := fmt.Sprintf("crm/v3/objects/emails/%s", emailId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *emails) BatchArchive(ctx context.Context, emailIds []string) error {
	u := "/crm/v3/objects/emails/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, emailId := range emailIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: emailId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *emails) BatchCreate(ctx context.Context, options *EmailBatchCreateOptions) (*EmailBatchOutput, error) {
	u := "/crm/v3/objects/emails/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	emails := &EmailBatchOutput{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func (z *emails) BatchRead(ctx context.Context, options *EmailBatchReadOptions) (*EmailBatchOutput, error) {
	u := "/crm/v3/objects/emails/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	emails := &EmailBatchOutput{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func (z *emails) BatchUpdate(ctx context.Context, options *EmailBatchUpdateOptions) (*EmailBatchOutput, error) {
	u := "/crm/v3/objects/emails/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	emails := &EmailBatchOutput{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func (z *emails) Search(ctx context.Context, options *EmailSearchOptions) (*EmailSearchResults, error) {
	u := "/crm/v3/objects/emails/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	emails := &EmailSearchResults{}

	err = z.client.do(req, emails)
	if err != nil {
		return nil, err
	}
	return emails, nil
}

func (z *emails) Merge(ctx context.Context, options *EmailMergeOptions) (*Email, error) {
	u := "/crm/v3/objects/emails/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	email := &Email{}

	err = z.client.do(req, email)
	if err != nil {
		return nil, err
	}
	return email, nil
}
