package hubspot

type ProductProperties struct {
	Amount                                string `json:"amount,omitempty"`
	Createdate                            string `json:"createdate,omitempty"`
	Description                           string `json:"description,omitempty"`
	Discount                              string `json:"discount,omitempty"`
	HsAllAccessibleTeamIds                string `json:"hs_all_accessible_team_ids,omitempty"`
	HsAllAssignedBusinessUnitIds          string `json:"hs_all_assigned_business_unit_ids,omitempty"`
	HsAllOwnerIds                         string `json:"hs_all_owner_ids,omitempty"`
	HsAllTeamIds                          string `json:"hs_all_team_ids,omitempty"`
	HsAvatarFilemanagerKey                string `json:"hs_avatar_filemanager_key,omitempty"`
	HsCostOfGoodsSold                     string `json:"hs_cost_of_goods_sold,omitempty"`
	HsCreatedByUserId                     string `json:"hs_created_by_user_id,omitempty"`
	HsCreatedate                          string `json:"hs_createdate,omitempty"`
	HsDiscountPercentage                  string `json:"hs_discount_percentage,omitempty"`
	HsFolderId                            string `json:"hs_folder_id,omitempty"`
	HsImages                              string `json:"hs_images,omitempty"`
	HsLastmodifieddate                    string `json:"hs_lastmodifieddate,omitempty"`
	HsMergedObjectIds                     string `json:"hs_merged_object_ids,omitempty"`
	HsObjectId                            string `json:"hs_object_id,omitempty"`
	HsProductType                         string `json:"hs_product_type,omitempty"`
	HsRecurringBillingPeriod              string `json:"hs_recurring_billing_period,omitempty"`
	HsRecurringBillingStartDate           string `json:"hs_recurring_billing_start_date,omitempty"`
	HsSku                                 string `json:"hs_sku,omitempty"`
	HsUniqueCreationKey                   string `json:"hs_unique_creation_key,omitempty"`
	HsUpdatedByUserId                     string `json:"hs_updated_by_user_id,omitempty"`
	HsUrl                                 string `json:"hs_url,omitempty"`
	HsUserIdsOfAllNotificationFollowers   string `json:"hs_user_ids_of_all_notification_followers,omitempty"`
	HsUserIdsOfAllNotificationUnfollowers string `json:"hs_user_ids_of_all_notification_unfollowers,omitempty"`
	HsUserIdsOfAllOwners                  string `json:"hs_user_ids_of_all_owners,omitempty"`
	HubspotOwnerAssigneddate              string `json:"hubspot_owner_assigneddate,omitempty"`
	HubspotOwnerId                        string `json:"hubspot_owner_id,omitempty"`
	HubspotTeamId                         string `json:"hubspot_team_id,omitempty"`
	Name                                  string `json:"name,omitempty"`
	Price                                 string `json:"price,omitempty"`
	Quantity                              string `json:"quantity,omitempty"`
	Recurringbillingfrequency             string `json:"recurringbillingfrequency,omitempty"`
	Tax                                   string `json:"tax,omitempty"`
}
