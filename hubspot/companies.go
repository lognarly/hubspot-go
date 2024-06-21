package hubspot

import (
	"context"
	"fmt"
)

type Companies interface {
	ListAssociations(ctx context.Context, query *CompanyAssociationsQuery, companyId string, toObjectType string) (*CompanyAssociations, error)
	Associate(ctx context.Context, companyId string, toObjectType string, toObjectId string, associationType string) (*Company, error)
	Disassociate(ctx context.Context, companyId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *CompanyListQuery) (*CompanyList, error)
	Create(ctx context.Context, options *CompanyCreateOrUpdateOptions) (*Company, error)
	Read(ctx context.Context, query *CompanyReadQuery, companyId string) (*Company, error)
	Update(ctx context.Context, options *CompanyCreateOrUpdateOptions, companyId string) (*Company, error)
	Archive(ctx context.Context, companyId string) error
	BatchArchive(ctx context.Context, companyIds []string) error
	BatchCreate(ctx context.Context, options *CompanyBatchCreateOptions) (*CompanyBatchOutput, error)
	BatchRead(ctx context.Context, options *CompanyBatchReadOptions) (*CompanyBatchOutput, error)
	BatchUpdate(ctx context.Context, options *CompanyBatchUpdateOptions) (*CompanyBatchOutput, error)
	Search(ctx context.Context, options *CompanySearchOptions) (*CompanySearchResults, error)
	Merge(ctx context.Context, options *CompanyMergeOptions) (*Company, error)
}

type companies struct {
	client *Client
}

type CompanyAssociationsQuery struct {
	ListAssociationsQuery
}

type CompanyAssociations struct {
	Results []CompanyAssociation `json:"results"`
	Pagination
}

type CompanyAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CompanyListQuery struct {
	ListQuery
}

type CompanyList struct {
	Companies []Company `json:"results"`
	Pagination
}

type Company struct {
	Id         string            `json:"id"`
	Properties CompanyProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type CompanyCreateOrUpdateOptions struct {
	Properties CompanyProperties `json:"properties"`
}

type CompanyReadQuery struct {
	ReadQuery
}

type CompanyUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type CompanyBatchOutput struct {
	Status      string    `json:"status"`
	Results     []Company `json:"results"`
	RequestedAt string    `json:"requestedAt"`
	StartedAt   string    `json:"startedAt"`
	CompletedAt string    `json:"completedAt"`
}

type CompanyBatchReadOptions struct {
	BatchReadOptions
}

type CompanyBatchCreateOptions struct {
	Inputs []CompanyCreateOrUpdateOptions `json:"inputs"`
}

type CompanyBatchUpdateOptions struct {
	Inputs []CompanyBatchUpdateProperties `json:"inputs"`
}

type CompanyBatchUpdateProperties struct {
	Id         string            `json:"id"`
	Properties CompanyProperties `json:"properties"`
}

type CompanySearchOptions struct {
	SearchOptions
}

type CompanySearchResults struct {
	Total   int64     `json:"total"`
	Results []Company `json:"results"`
	Pagination
}

type CompanyMergeOptions struct {
	MergeOptions
}

func (z *companies) ListAssociations(ctx context.Context, query *CompanyAssociationsQuery, companyId string, toObjectType string) (*CompanyAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/%s/associations/%s", companyId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ca := &CompanyAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

func (z *companies) Associate(ctx context.Context, companyId string, toObjectType string, toObjectId string, associationType string) (*Company, error) {
	u := fmt.Sprintf("/crm/v3/objects/companies/%s/associations/%s/%s/%s", companyId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (z *companies) Disassociate(ctx context.Context, companyId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/companies/%s/associations/%s/%s/%s", companyId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *companies) List(ctx context.Context, query *CompanyListQuery) (*CompanyList, error) {
	u := "/crm/v3/objects/companies"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	cl := &CompanyList{}

	err = z.client.do(req, cl)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (z *companies) Create(ctx context.Context, options *CompanyCreateOrUpdateOptions) (*Company, error) {
	u := "/crm/v3/objects/companies"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (z *companies) Read(ctx context.Context, query *CompanyReadQuery, companyId string) (*Company, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", companyId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (z *companies) Update(ctx context.Context, options *CompanyCreateOrUpdateOptions, companyId string) (*Company, error) {
	u := fmt.Sprintf("crm/v3/objects/companies/%s", companyId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}

func (z *companies) Archive(ctx context.Context, companyId string) error {
	u := fmt.Sprintf("crm/v3/objects/companies/%s", companyId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *companies) BatchArchive(ctx context.Context, companyIds []string) error {
	u := "/crm/v3/objects/companies/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, companyId := range companyIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: companyId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *companies) BatchCreate(ctx context.Context, options *CompanyBatchCreateOptions) (*CompanyBatchOutput, error) {
	u := "/crm/v3/objects/companies/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	companies := &CompanyBatchOutput{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (z *companies) BatchRead(ctx context.Context, options *CompanyBatchReadOptions) (*CompanyBatchOutput, error) {
	u := "/crm/v3/objects/companies/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	companies := &CompanyBatchOutput{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (z *companies) BatchUpdate(ctx context.Context, options *CompanyBatchUpdateOptions) (*CompanyBatchOutput, error) {
	u := "/crm/v3/objects/companies/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	companies := &CompanyBatchOutput{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (z *companies) Search(ctx context.Context, options *CompanySearchOptions) (*CompanySearchResults, error) {
	u := "/crm/v3/objects/companies/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	companies := &CompanySearchResults{}

	err = z.client.do(req, companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func (z *companies) Merge(ctx context.Context, options *CompanyMergeOptions) (*Company, error) {
	u := "/crm/v3/objects/companies/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	company := &Company{}

	err = z.client.do(req, company)
	if err != nil {
		return nil, err
	}

	return company, nil
}
