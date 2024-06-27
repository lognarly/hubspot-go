package hubspot

import (
	"context"
	"fmt"
)

type Products interface {
	ListAssociations(ctx context.Context, query *ProductAssociationsQuery, productId string, toObjectType string) (*ProductAssociations, error)
	Associate(ctx context.Context, productId string, toObjectType string, toObjectId string, associationType string) (*Product, error)
	Disassociate(ctx context.Context, productId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *ProductListQuery) (*ProductList, error)
	Create(ctx context.Context, options *ProductCreateOrUpdateOptions) (*Product, error)
	Read(ctx context.Context, query *ProductReadQuery, productId string) (*Product, error)
	Update(ctx context.Context, productId string, options *ProductCreateOrUpdateOptions) (*Product, error)
	Archive(ctx context.Context, productId string) error
	BatchArchive(ctx context.Context, productIds []string) error
	BatchCreate(ctx context.Context, options *ProductBatchCreateOptions) (*ProductBatchOutput, error)
	BatchRead(ctx context.Context, options *ProductBatchReadOptions) (*ProductBatchOutput, error)
	BatchUpdate(ctx context.Context, options *ProductBatchUpdateOptions) (*ProductBatchOutput, error)
	Search(ctx context.Context, options *ProductSearchOptions) (*ProductSearchResults, error)
	Merge(ctx context.Context, options *ProductMergeOptions) (*Product, error)
}

type products struct {
	client *Client
}

type ProductAssociationsQuery struct {
	ListAssociationsQuery
}

type ProductAssociations struct {
	Results []ProductAssociation `json:"results"`
	Pagination
}

type ProductAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ProductListQuery struct {
	ListQuery
}

type ProductList struct {
	Products []Product `json:"results"`
	Pagination
}

type Product struct {
	Id         string            `json:"id"`
	Properties ProductProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type ProductCreateOrUpdateOptions struct {
	Properties ProductProperties `json:"properties"`
}

type ProductReadQuery struct {
	ReadQuery
}

type ProductBatchOutput struct {
	Status      string    `json:"status"`
	Results     []Product `json:"results"`
	RequestedAt string    `json:"requestedAt"`
	StartedAt   string    `json:"startedAt"`
	CompletedAt string    `json:"completedAt"`
}

type ProductBatchReadOptions struct {
	BatchReadOptions
}

type ProductBatchCreateOptions struct {
	Inputs []ProductCreateOrUpdateOptions `json:"inputs"`
}

type ProductBatchUpdateOptions struct {
	Inputs []ProductBatchUpdateProperties `json:"inputs"`
}

type ProductBatchUpdateProperties struct {
	Id         string            `json:"id"`
	Properties ProductProperties `json:"properties"`
}

type ProductSearchOptions struct {
	SearchOptions
}

type ProductSearchResults struct {
	Total   int64     `json:"total"`
	Results []Product `json:"results"`
	Pagination
}

type ProductMergeOptions struct {
	MergeOptions
}

func (z *products) ListAssociations(ctx context.Context, query *ProductAssociationsQuery, productId string, toObjectType string) (*ProductAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/%s/associations/%s", productId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	pa := &ProductAssociations{}

	err = z.client.do(req, pa)
	if err != nil {
		return nil, err
	}
	return pa, nil
}

func (z *products) Associate(ctx context.Context, productId string, toObjectType string, toObjectId string, associationType string) (*Product, error) {
	u := fmt.Sprintf("/crm/v3/objects/products/%s/associations/%s/%s/%s", productId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (z *products) Disassociate(ctx context.Context, productId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/products/%s/associations/%s/%s/%s", productId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *products) List(ctx context.Context, query *ProductListQuery) (*ProductList, error) {
	u := "crm/v3/objects/products"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	pl := &ProductList{}

	err = z.client.do(req, pl)
	if err != nil {
		return nil, err
	}
	return pl, nil
}

func (z *products) Create(ctx context.Context, options *ProductCreateOrUpdateOptions) (*Product, error) {
	u := "/crm/v3/objects/products"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (z *products) Read(ctx context.Context, query *ProductReadQuery, productId string) (*Product, error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (z *products) Update(ctx context.Context, productId string, options *ProductCreateOrUpdateOptions) (*Product, error) {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (z *products) Archive(ctx context.Context, productId string) error {
	u := fmt.Sprintf("crm/v3/objects/products/%s", productId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *products) BatchArchive(ctx context.Context, productIds []string) error {
	u := "/crm/v3/objects/products/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, productId := range productIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: productId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *products) BatchCreate(ctx context.Context, options *ProductBatchCreateOptions) (*ProductBatchOutput, error) {
	u := "/crm/v3/objects/products/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	products := &ProductBatchOutput{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (z *products) BatchRead(ctx context.Context, options *ProductBatchReadOptions) (*ProductBatchOutput, error) {
	u := "/crm/v3/objects/products/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	products := &ProductBatchOutput{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (z *products) BatchUpdate(ctx context.Context, options *ProductBatchUpdateOptions) (*ProductBatchOutput, error) {
	u := "/crm/v3/objects/products/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	products := &ProductBatchOutput{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (z *products) Search(ctx context.Context, options *ProductSearchOptions) (*ProductSearchResults, error) {
	u := "/crm/v3/objects/products/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	products := &ProductSearchResults{}

	err = z.client.do(req, products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (z *products) Merge(ctx context.Context, options *ProductMergeOptions) (*Product, error) {
	u := "/crm/v3/objects/products/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	product := &Product{}

	err = z.client.do(req, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
