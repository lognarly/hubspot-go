package hubspot

import (
	"context"
	"fmt"
)

type Meetings interface {
	ListAssociations(ctx context.Context, query *MeetingAssociationsQuery, meetingId string, toObjectType string) (*MeetingAssociations, error)
	Associate(ctx context.Context, meetingId string, toObjectType string, toObjectId string, associationType string) (*Meeting, error)
	Disassociate(ctx context.Context, meetingId string, toObjectType string, toObjectId string, associationType string) error
	List(ctx context.Context, query *MeetingListQuery) (*MeetingList, error)
	Create(ctx context.Context, options *MeetingCreateOrUpdateOptions) (*Meeting, error)
	Read(ctx context.Context, query *MeetingReadQuery, meetingId string) (*Meeting, error)
	Update(ctx context.Context, options *MeetingCreateOrUpdateOptions, meetingId string) (*Meeting, error)
	Archive(ctx context.Context, meetingId string) error
	BatchArchive(ctx context.Context, meetingIds []string) error
	BatchCreate(ctx context.Context, options *MeetingBatchCreateOptions) (*MeetingBatchOutput, error)
	BatchRead(ctx context.Context, options *MeetingBatchReadOptions) (*MeetingBatchOutput, error)
	BatchUpdate(ctx context.Context, options *MeetingBatchUpdateOptions) (*MeetingBatchOutput, error)
	Search(ctx context.Context, options *MeetingSearchOptions) (*MeetingSearchResults, error)
	Merge(ctx context.Context, options *MeetingMergeOptions) (*Meeting, error)
}

type meetings struct {
	client *Client
}

type MeetingAssociationsQuery struct {
	ListAssociationsQuery
}

type MeetingAssociations struct {
	Results []MeetingAssociation `json:"results"`
	Pagination
}

type MeetingAssociation struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type MeetingListQuery struct {
	ListQuery
}

type MeetingList struct {
	Meetings []Meeting `json:"results"`
	Pagination
}

type Meeting struct {
	Id         string            `json:"id"`
	Properties MeetingProperties `json:"properties"`
	CreatedAt  string            `json:"createdAt"`
	UpdatedAt  string            `json:"updatedAt"`
	Archived   bool              `json:"archived"`
}

type MeetingProperties struct {
	CreateDate             string `json:"createdate"`
	HsInternalMeetingNotes string `json:"hs_internal_meeting_notes,omitempty"`
	HsLastModifiedDate     string `json:"hs_lastmodifieddate"`
	HsMeetingBody          string `json:"hs_meeting_body,omitempty"`
	HsMeetingEndTime       string `json:"hs_meeting_end_time,omitempty"`
	HsMeetingExternalUrl   string `json:"hs_meeting_external_url,omitempty"`
	HsMeetingLocation      string `json:"hs_meeting_location,omitempty"`
	HsMeetingOutcome       string `json:"hs_meeting_outcome,omitempty"`
	HsMeetingStartTime     string `json:"hs_meeting_start_time,omitempty"`
	HsMeetingTitle         string `json:"hs_meeting_title,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type MeetingCreateOrUpdateOptions struct {
	Properties MeetingCreateOrUpdateProperties `json:"properties"`
}

type MeetingCreateOrUpdateProperties struct {
	HsInternalMeetingNotes string `json:"hs_internal_meeting_notes,omitempty"`
	HsMeetingBody          string `json:"hs_meeting_body,omitempty"`
	HsMeetingEndTime       string `json:"hs_meeting_end_time,omitempty"`
	HsMeetingExternalUrl   string `json:"hs_meeting_external_url,omitempty"`
	HsMeetingLocation      string `json:"hs_meeting_location,omitempty"`
	HsMeetingOutcome       string `json:"hs_meeting_outcome,omitempty"`
	HsMeetingStartTime     string `json:"hs_meeting_start_time,omitempty"`
	HsMeetingTitle         string `json:"hs_meeting_title,omitempty"`
	HsTimestamp            string `json:"hs_timestamp,omitempty"`
	HubSpotOwnerId         string `json:"hubspot_owner_id,omitempty"`
}

type MeetingReadQuery struct {
	ReadQuery
}

type MeetingUpdateQuery struct {
	IdProperty string `url:"idProperty"`
}

type MeetingBatchOutput struct {
	Status      string    `json:"status"`
	Results     []Meeting `json:"results"`
	RequestedAt string    `json:"requestedAt"`
	StartedAt   string    `json:"startedAt"`
	CompletedAt string    `json:"completedAt"`
}

type MeetingBatchReadOptions struct {
	BatchReadOptions
}

type MeetingBatchCreateOptions struct {
	Inputs []MeetingCreateOrUpdateOptions `json:"inputs"`
}

type MeetingBatchUpdateOptions struct {
	Inputs []MeetingBatchUpdateProperties `json:"inputs"`
}

type MeetingBatchUpdateProperties struct {
	Id         string                          `json:"id"`
	Properties MeetingCreateOrUpdateProperties `json:"properties"`
}

type MeetingSearchOptions struct {
	SearchOptions
}

type MeetingSearchResults struct {
	Total   int64     `json:"total"`
	Results []Meeting `json:"results"`
	Pagination
}

type MeetingMergeOptions struct {
	MergeOptions
}

func (z *meetings) ListAssociations(ctx context.Context, query *MeetingAssociationsQuery, meetingId string, toObjectType string) (*MeetingAssociations, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/%s/associations/%s", meetingId, toObjectType)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ca := &MeetingAssociations{}

	err = z.client.do(req, ca)
	if err != nil {
		return nil, err
	}
	return ca, nil
}

func (z *meetings) Associate(ctx context.Context, meetingId string, toObjectType string, toObjectId string, associationType string) (*Meeting, error) {
	u := fmt.Sprintf("/crm/v3/objects/meetings/%s/associations/%s/%s/%s", meetingId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "PUT", u, nil)
	if err != nil {
		return nil, err
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}

func (z *meetings) Disassociate(ctx context.Context, meetingId string, toObjectType string, toObjectId string, associationType string) error {
	u := fmt.Sprintf("/crm/v3/objects/meetings/%s/associations/%s/%s/%s", meetingId, toObjectType, toObjectId, associationType)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *meetings) List(ctx context.Context, query *MeetingListQuery) (*MeetingList, error) {
	u := "/crm/v3/objects/meetings"
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	ml := &MeetingList{}

	err = z.client.do(req, ml)
	if err != nil {
		return nil, err
	}
	return ml, nil
}

func (z *meetings) Create(ctx context.Context, options *MeetingCreateOrUpdateOptions) (*Meeting, error) {
	u := "/crm/v3/objects/meetings"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}

func (z *meetings) Read(ctx context.Context, query *MeetingReadQuery, meetingId string) (*Meeting, error) {
	u := fmt.Sprintf("crm/v3/objects/lineitems/%s", meetingId)
	req, err := z.client.newHttpRequest(ctx, "GET", u, query)
	if err != nil {
		return nil, err
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}

func (z *meetings) Update(ctx context.Context, options *MeetingCreateOrUpdateOptions, meetingId string) (*Meeting, error) {
	u := fmt.Sprintf("crm/v3/objects/meetings/%s", meetingId)
	req, err := z.client.newHttpRequest(ctx, "PATCH", u, options)
	if err != nil {
		return nil, err
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}

func (z *meetings) Archive(ctx context.Context, meetingId string) error {
	u := fmt.Sprintf("crm/v3/objects/meetings/%s", meetingId)
	req, err := z.client.newHttpRequest(ctx, "DELETE", u, nil)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *meetings) BatchArchive(ctx context.Context, meetingIds []string) error {
	u := "/crm/v3/objects/meetings/batch/archive"

	options := BatchInputOptions{}
	options.Inputs = make([]BatchInput, 0)

	for _, meetingId := range meetingIds {
		options.Inputs = append(options.Inputs, BatchInput{Id: meetingId})
	}

	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return err
	}
	return z.client.do(req, nil)
}

func (z *meetings) BatchCreate(ctx context.Context, options *MeetingBatchCreateOptions) (*MeetingBatchOutput, error) {
	u := "/crm/v3/objects/meetings/batch/create"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	meetings := &MeetingBatchOutput{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, err
	}
	return meetings, nil
}

func (z *meetings) BatchRead(ctx context.Context, options *MeetingBatchReadOptions) (*MeetingBatchOutput, error) {
	u := "/crm/v3/objects/meetings/batch/read"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	meetings := &MeetingBatchOutput{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, err
	}
	return meetings, nil
}

func (z *meetings) BatchUpdate(ctx context.Context, options *MeetingBatchUpdateOptions) (*MeetingBatchOutput, error) {
	u := "/crm/v3/objects/meetings/batch/update"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	meetings := &MeetingBatchOutput{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, err
	}
	return meetings, nil
}

func (z *meetings) Search(ctx context.Context, options *MeetingSearchOptions) (*MeetingSearchResults, error) {
	u := "/crm/v3/objects/meetings/search"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	meetings := &MeetingSearchResults{}

	err = z.client.do(req, meetings)
	if err != nil {
		return nil, err
	}
	return meetings, nil
}

func (z *meetings) Merge(ctx context.Context, options *MeetingMergeOptions) (*Meeting, error) {
	u := "/crm/v3/objects/meetings/merge"
	req, err := z.client.newHttpRequest(ctx, "POST", u, options)
	if err != nil {
		return nil, err
	}

	meeting := &Meeting{}

	err = z.client.do(req, meeting)
	if err != nil {
		return nil, err
	}
	return meeting, nil
}
