package hubspot

import (
	"context"
	"fmt"
)

type Tasks interface {
	ListAssociations(ctx context.Context, query *TaskAssociationsQuery, taskId string, toObjectType string) (*TaskAssociations, error)
	Associate(ctx context.Context, taskId string, toObjectType string, toObjectId string, associationType string) (*Task, error)
	Disassociate(ctx context.Context, taskId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *TaskListQuery) (*TaskList, error)
	Create(ctx context.Context, options *TaskCreateOrUpdateOptions) (*Task, error)
	Read(ctx context.Context, query *TaskReadQuery, taskId string) (*Task, error)
	Update(ctx context.Context, options *TaskCreateOrUpdateOptions, taskId string) (*Task, error)
	Archive(ctx context.Context, taskId string) error
	BatchArchive(ctx context.Context, taskIds []string) error
	BatchCreate(ctx context.Context, options *TaskBatchCreateOptions) (*TaskBatchOutput, error)
	BatchRead(ctx context.Context, options *TaskBatchReadOptions) (*TaskBatchOutput, error)
	BatchUpdate(ctx context.Context, options *TaskBatchUpdateOptions) (*TaskBatchOutput, error)
	Search(ctx context.Context, options *TaskSearchOptions) (*TaskSearchResults, error)
	Merge(ctx context.Context, options *TaskMergeOptions) (*Task, error)
}

type tasks struct {
	client *Client
}

type TaskAssociationsQuery struct {
	ListAssociationsQuery
}

type TaskAssociations struct {
	Results []TaskAssociation `json:"results"`
	Pagination
}

type TaskAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type TaskListQuery struct {
	ListQuery
}

type TaskList struct {
	Tasks []Task `json:"results"`
	Pagination
}

type Task struct {
	Id         string         `json:"id"`
	Properties TaskProperties `json:"properties"`
	CreatedAt  string         `json:"createdAt"`
	UpdatedAt  string         `json:"updatedAt"`
	Archived   bool           `json:"archived"`
}

type TaskProperties struct {
	CreateDate         string `json:"createdate"`
	HsLastModifiedDate string `json:"hs_lastmodifieddate"`
	HsTaskBody         string `json:"hs_task_body"`
	HsTaskPriority     string `json:"hs_task_priority"`
	HsTaskStatus       string `json:"hs_task_status"`
	HsTaskSubject      string `json:"hs_task_subject"`
	HsTimestamp        string `json:"hs_timestamp"`
	HubSpotOwnerId     string `json:"hubspot_owner_id"`
}

type TaskCreateOrUpdateOptions struct {
	Properties TaskCreateOrUpdateProperties `json:"properties"`
}

type TaskCreateOrUpdateProperties struct {
	HsTaskBody     string `json:"hs_task_body"`
	HsTaskPriority string `json:"hs_task_priority"`
	HsTaskStatus   string `json:"hs_task_status"`
	HsTaskSubject  string `json:"hs_task_subject"`
	HsTimestamp    string `json:"hs_timestamp"`
	HubSpotOwnerId string `json:"hubspot_owner_id"`
}

type TaskReadQuery struct {
	ReadQuery
}

type TaskUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type TaskBatchOutput struct {
	Status      string `json:"status"`
	Results     []Task `json:"results"`
	RequestedAt string `json:"requestedAt"`
	StartedAt   string `json:"startedAt"`
	CompletedAt string `json:"completedAt"`
}

type TaskBatchReadOptions struct {
	BatchReadOptions
}

type TaskBatchCreateOptions struct {
	Inputs []TaskCreateOrUpdateOptions `json:"inputs"`
}

type TaskBatchUpdateOptions struct {
	Inputs []TaskBatchUpdateProperties `json:"inputs"`
}

type TaskBatchUpdateProperties struct {
	Id         string                       `json:"id"`
	Properties TaskCreateOrUpdateProperties `json:"properties"`
}

type TaskSearchOptions struct {
	SearchOptions
}

type TaskSearchResults struct {
	Total   int64  `json:"total"`
	Results []Task `json:"results"`
	Pagination
}

type TaskMergeOptions struct {
	MergeOptions
}

func (z *tasks) ListAssociations(ctx context.Context, query *TaskAssociationsQuery, taskId string, toObjectType string) (*TaskAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/%s/associations/%s", taskId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ca := &TaskAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, err
	}

	return ca, nil
}

func (z *tasks) Associate(ctx context.Context, taskId string, toObjectType string, toObjectId string, associationType string) (*Task, error) {
	u := fmt.Sprintf("/crm/v3/objects/tasks/%s/associations/%s/%s/%s", taskId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (z *tasks) Disassociate(ctx context.Context, taskId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/tasks/%s/associations/%s/%s/%s", taskId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *tasks) List(ctx context.Context, query *TaskListQuery) (*TaskList, error) {
	u := "/crm/v3/objects/tasks"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	tl := &TaskList{}

	err = z.client.do(req, tl)
	if err != nil {
		return nil, err
	}

	return tl, nil
}

func (z *tasks) Create(ctx context.Context, options *TaskCreateOrUpdateOptions) (*Task, error) {
	u := "/crm/v3/objects/tasks"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (z *tasks) Read(ctx context.Context, query *TaskReadQuery, taskId string) (*Task, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", taskId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (z *tasks) Update(ctx context.Context, options *TaskCreateOrUpdateOptions, taskId string) (*Task, error) {
	u := fmt.Sprintf("crm/v3/objects/tasks/%s", taskId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (z *tasks) Archive(ctx context.Context, taskId string) error {
	u := fmt.Sprintf("crm/v3/objects/tasks/%s", taskId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *tasks) BatchArchive(ctx context.Context, taskIds []string) error {
	u := "/crm/v3/objects/tasks/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, taskId := range taskIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: taskId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}

	return z.client.do(req, nil)
}

func (z *tasks) BatchCreate(ctx context.Context, options *TaskBatchCreateOptions) (*TaskBatchOutput, error) {
	u := "/crm/v3/objects/tasks/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tasks := &TaskBatchOutput{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (z *tasks) BatchRead(ctx context.Context, options *TaskBatchReadOptions) (*TaskBatchOutput, error) {
	u := "/crm/v3/objects/tasks/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tasks := &TaskBatchOutput{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (z *tasks) BatchUpdate(ctx context.Context, options *TaskBatchUpdateOptions) (*TaskBatchOutput, error) {
	u := "/crm/v3/objects/tasks/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tasks := &TaskBatchOutput{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (z *tasks) Search(ctx context.Context, options *TaskSearchOptions) (*TaskSearchResults, error) {
	u := "/crm/v3/objects/tasks/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	tasks := &TaskSearchResults{}

	err = z.client.do(req, tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (z *tasks) Merge(ctx context.Context, options *TaskMergeOptions) (*Task, error) {
	u := "/crm/v3/objects/tasks/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	task := &Task{}

	err = z.client.do(req, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
