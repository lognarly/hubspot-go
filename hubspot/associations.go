package hubspot

import (
	"context"
	"fmt"
	"strconv"
)

type Associations interface {
	List(ctx context.Context, fromObjectType string, fromObjectId int64, toObjectType string, query *AssociationListQuery) (*AssociationList, error)
	Create(ctx context.Context, options *[]AssociationCreateOptions, fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) (*AssociationCreateOutput, error)
	Delete(ctx context.Context, fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) error
	ReadDefinition(ctx context.Context, fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error)
	CreateDefinition(ctx context.Context, options *AssociationCreateDefinitionOptions, fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error)
	UpdateDefinition(ctx context.Context, options *AssociationUpdateDefinitionOptions, fromObjectType string, toObjectType string) error
	DeleteDefinition(ctx context.Context, fromObjectType string, toObjectType string, typeId int64) error
}

type associations struct {
	client *Client
}

type AssociationListQuery struct {
	ListAssociationsQuery
}

type AssociationList struct {
	Assocations []AssociationListResult `json:"results"`
	Pagination
}

type AssociationListResult struct {
	ToObjectId       int64             `json:"toObjectId"`
	AssociationTypes []AssociationType `json:"associationTypes"`
}

type HubspotAssociationCategory string
type HubspotAssociationTypeId int

const (
	HubSpotDefined    HubspotAssociationCategory = "HUBSPOT_DEFINED"
	UserDefined       HubspotAssociationCategory = "USER_DEFINED"
	IntegratorDefined HubspotAssociationCategory = "INTEGRATOR_DEFINED"

	LineItemToDealTypeId HubspotAssociationTypeId = 20
)

type AssociationType struct {
	Category HubspotAssociationCategory `json:"category,omitempty"`
	TypeId   int64                      `json:"typeId,omitempty"`
	Label    string                     `json:"label,omitempty"`
}

type AssociationCreateOptions struct {
	Category HubspotAssociationCategory `json:"associationCategory"`
	TypeId   HubspotAssociationTypeId   `json:"associationTypeId"`
}

type AssociationCreateOutput struct {
	FromObjectTypeId string   `json:"fromObjectTypeId"`
	FromObjectId     int64    `json:"fromObjectId"`
	ToObjectTypeId   string   `json:"toObjectTypeId"`
	ToObjectId       int64    `json:"toObjectId"`
	Labels           []string `json:"labels,omitempty"`
}

type AssociationDefinitionOutput struct {
	Results []AssociationDefinition `json:"results"`
}

type AssociationDefinition struct {
	Category HubspotAssociationCategory `json:"category"`
	TypeId   int                        `json:"typeId"`
	Label    string                     `json:"label"`
}

type AssociationCreateDefinitionOptions struct {
	Label string `json:"label"`
	Name  string `json:"name"`
}

type AssociationUpdateDefinitionOptions struct {
	Label  string `json:"label"`
	TypeId int64  `json:"associationTypeId"`
}

type Association struct {
	Types []AssociationCreateOptions `json:"types"`
	To    AssociationTo              `json:"to"`
}

type AssociationTo struct {
	Id string `json:"id"`
}

func (a *associations) List(ctx context.Context, fromObjectType string, fromObjectId int64, toObjectType string, query *AssociationListQuery) (*AssociationList, error) {
	u := fmt.Sprintf("crm/v4/objects/%s/%s/associations/%s", fromObjectType, strconv.FormatInt(fromObjectId, 10), toObjectType)
	req, err := a.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	al := &AssociationList{}

	err = a.client.do(req, al)
	if err != nil {
		return nil, err
	}
	return al, nil
}

func (a *associations) Create(ctx context.Context, options *[]AssociationCreateOptions, fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) (*AssociationCreateOutput, error) {
	u := fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", fromObjectType, strconv.FormatInt(fromObjectId, 10), toObjectType, strconv.FormatInt(toObjectId, 10))
	req, err := a.client.newHttpRequest(ctx, "PUT", u, options)
	if err != nil {
		return nil, err
	}

	aco := &AssociationCreateOutput{}

	err = a.client.do(req, aco)
	if err != nil {
		return nil, err
	}
	return aco, nil
}

func (a *associations) Delete(ctx context.Context, fromObjectType string, fromObjectId int64, toObjectType string, toObjectId int64) error {
	u := fmt.Sprintf("/crm/v4/objects/%s/%s/associations/%s/%s", fromObjectType, strconv.FormatInt(fromObjectId, 10), toObjectType, strconv.FormatInt(toObjectId, 10))
	req, err := a.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return a.client.do(req, nil)
}

func (a *associations) ReadDefinition(ctx context.Context, fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", fromObjectType, toObjectType)
	req, err := a.client.newHttpRequest(ctx, "GET", u, nil)
	if err != nil {
		return nil, err
	}

	ard := &AssociationDefinitionOutput{}

	err = a.client.do(req, ard)
	if err != nil {
		return nil, err
	}
	return ard, nil
}

func (a *associations) CreateDefinition(ctx context.Context, options *AssociationCreateDefinitionOptions, fromObjectType string, toObjectType string) (*AssociationDefinitionOutput, error) {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", fromObjectType, toObjectType)
	req, err := a.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	ard := &AssociationDefinitionOutput{}

	err = a.client.do(req, ard)
	if err != nil {
		return nil, err
	}
	return ard, nil
}

func (a *associations) UpdateDefinition(ctx context.Context, options *AssociationUpdateDefinitionOptions, fromObjectType string, toObjectType string) error {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels", fromObjectType, toObjectType)
	req, err := a.client.newHttpRequest(ctx, "PUT", u, options)
	if err != nil {
		return err
	}

	return a.client.do(req, nil)
}

func (a *associations) DeleteDefinition(ctx context.Context, fromObjectType string, toObjectType string, typeId int64) error {
	u := fmt.Sprintf("/crm/v4/associations/%s/%s/labels/%s", fromObjectType, toObjectType, strconv.FormatInt(typeId, 10))
	req, err := a.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return a.client.do(req, nil)
}
