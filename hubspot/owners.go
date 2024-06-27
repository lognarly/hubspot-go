package hubspot

import (
	"context"
	"fmt"
)

type Owners interface {
	List(ctx context.Context, query *OwnerListQuery) (*OwnerList, error)
	Read(ctx context.Context, ownerId string, query *OwnerReadQuery) (*Owner, error)
}

type owners struct {
	client *Client
}

type OwnerListQuery struct {
	Email    string `url:"email,omitempty"`
	After    string `ur:"after,omitempty"`
	Limit    string `url:"limit,omitempty"`
	Archived bool   `url:"archived,omitempty"`
}

type OwnerList struct {
	Results []Owner `json:"results"`
	Pagination
}

type Owner struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	UserId    int64  `json:"userId,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	Archived  bool   `json:"archived"`
	Teams     []Team `json:"teams,omitempty"`
}

type Team struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Primary bool   `json:"primary"`
}

type OwnerReadQuery struct {
	IdProperty string `url:"idProperty,omitempty"`
	Archived   bool   `url:"archived,omitempty"`
}

func (z *owners) List(ctx context.Context, query *OwnerListQuery) (*OwnerList, error) {
	u := "/crm/v3/owners"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ol := &OwnerList{}

	err = z.client.do(req, ol)
	if err != nil {
		return nil, err
	}

	return ol, nil
}

func (z *owners) Read(ctx context.Context, ownerId string, query *OwnerReadQuery) (*Owner, error) {
	u := fmt.Sprintf("/crm/v3/owners/%s", ownerId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	owner := &Owner{}

	err = z.client.do(req, owner)
	if err != nil {
		return nil, err
	}

	return owner, nil
}
