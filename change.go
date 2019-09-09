package servicenow

import (
	"net/url"
)

// TableChangeRequests defines the name of the table withing the JSONv2 web service to interface with
// SNOW CHG items
const TableChangeRequests = "change_request"

// GetChangeRequests method will take a url.Value type argument and call the GetRecordsFor method with
// the change_request table and query as the arguments, then format the response into the ChangeRequest type
func (c Client) GetChangeRequests(query url.Values) ([]struct{ChangeRequest}, error) {
	crs := ChangeRequests{}
	err := c.GetRecordsFor(TableChangeRequests, query, &crs)
	return crs.Result, err
}

type ChangeRequest struct {
	Parent  string `json:"parent"`
	Reason  string `json:"reason"`
	UCfUser struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"u_cf_user"`
	USoc2                 string `json:"u_soc2"`
	WatchList             string `json:"watch_list"`
	UponReject            string `json:"upon_reject"`
	SysUpdatedOn          string `json:"sys_updated_on"`
	Type                  string `json:"type"`
	ApprovalHistory       string `json:"approval_history"`
	UWorkDuration         string `json:"u_work_duration"`
	UChangeTemplate       string `json:"u_change_template"`
	Number                string `json:"number"`
	TestPlan              string `json:"test_plan"`
	CabDelegate           string `json:"cab_delegate"`
	URollback             string `json:"u_rollback"`
	RequestedByDate       string `json:"requested_by_date"`
	State                 string `json:"state"`
	SysCreatedBy          string `json:"sys_created_by"`
	Knowledge             string `json:"knowledge"`
	Order                 string `json:"order"`
	Phase                 string `json:"phase"`
	USLAPercentage        string `json:"u_sla_percentage"`
	USymptomDetectionTime string `json:"u_symptom_detection_time"`
	CmdbCi                string `json:"cmdb_ci"`
	Impact                string `json:"impact"`
	Active                string `json:"active"`
	WorkNotesList         string `json:"work_notes_list"`
	Priority              string `json:"priority"`
	SysDomainPath         string `json:"sys_domain_path"`
	CabRecommendation     string `json:"cab_recommendation"`
	ProductionSystem      string `json:"production_system"`
	ReviewDate            string `json:"review_date"`
	URecordProducer       string `json:"u_record_producer"`
	BusinessDuration      string `json:"business_duration"`
	GroupList             string `json:"group_list"`
	RequestedBy           struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"requested_by"`
	UTemplate                  string `json:"u_template"`
	ChangePlan                 string `json:"change_plan"`
	ApprovalSet                string `json:"approval_set"`
	UOutages                   string `json:"u_outages"`
	WfActivity                 string `json:"wf_activity"`
	ImplementationPlan         string `json:"implementation_plan"`
	EndDate                    string `json:"end_date"`
	ShortDescription           string `json:"short_description"`
	UJiraSubbatches            string `json:"u_jira_subbatches"`
	CorrelationDisplay         string `json:"correlation_display"`
	UImplementationPlanCopy    string `json:"u_implementation_plan_copy"`
	WorkStart                  string `json:"work_start"`
	AdditionalAssigneeList     string `json:"additional_assignee_list"`
	OutsideMaintenanceSchedule string `json:"outside_maintenance_schedule"`
	StdChangeProducerVersion   string `json:"std_change_producer_version"`
	ServiceOffering            string `json:"service_offering"`
	SysClassName               string `json:"sys_class_name"`
	ClosedBy                   struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"closed_by"`
	FollowUp                string `json:"follow_up"`
	UJiraLink               string `json:"u_jira_link"`
	UBridgeRequired         string `json:"u_bridge_required"`
	USecurityReviewRequired string `json:"u_security_review_required"`
	ReassignmentCount       string `json:"reassignment_count"`
	ReviewStatus            string `json:"review_status"`
	AssignedTo              struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"assigned_to"`
	StartDate                string `json:"start_date"`
	SLADue                   string `json:"sla_due"`
	CommentsAndWorkNotes     string `json:"comments_and_work_notes"`
	UPlannedDuration         string `json:"u_planned_duration"`
	Escalation               string `json:"escalation"`
	UponApproval             string `json:"upon_approval"`
	CorrelationID            string `json:"correlation_id"`
	UCfACLSecureBusSvcGroups string `json:"u_cf_acl_secure_bus_svc_groups"`
	UJiraUrls                string `json:"u_jira_urls"`
	UCustomer                string `json:"u_customer"`
	MadeSLA                  string `json:"made_sla"`
	UBusinessServiceSubcat   string `json:"u_business_service_subcat"`
	BackoutPlan              string `json:"backout_plan"`
	UJiraComponent           string `json:"u_jira_component"`
	ConflictStatus           string `json:"conflict_status"`
	SysUpdatedBy             string `json:"sys_updated_by"`
	UJiraFrequency           string `json:"u_jira_frequency"`
	OpenedBy                 struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"opened_by"`
	UCfACLSecureBusSvcUsers string `json:"u_cf_acl_secure_bus_svc_users"`
	UFirstTouch             string `json:"u_first_touch"`
	UserInput               string `json:"user_input"`
	SysCreatedOn            string `json:"sys_created_on"`
	SysDomain               struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"sys_domain"`
	ClosedAt               string `json:"closed_at"`
	UReassignmentUserCount string `json:"u_reassignment_user_count"`
	ReviewComments         string `json:"review_comments"`
	BusinessService        struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"business_service"`
	UOutsidePlannedDates string `json:"u_outside_planned_dates"`
	TimeWorked           string `json:"time_worked"`
	ExpectedStart        string `json:"expected_start"`
	OpenedAt             string `json:"opened_at"`
	WorkEnd              string `json:"work_end"`
	PhaseState           string `json:"phase_state"`
	CabDate              string `json:"cab_date"`
	WorkNotes            string `json:"work_notes"`
	UPotentialImpact     string `json:"u_potential_impact"`
	USubstate            string `json:"u_substate"`
	UDepartment          string `json:"u_department"`
	CloseCode            string `json:"close_code"`
	AssignmentGroup      struct {
		Link  string `json:"link"`
		Value string `json:"value"`
	} `json:"assignment_group"`
	UAction            string `json:"u_action"`
	Description        string `json:"description"`
	OnHoldReason       string `json:"on_hold_reason"`
	CalendarDuration   string `json:"calendar_duration"`
	CloseNotes         string `json:"close_notes"`
	UJiraLocation      string `json:"u_jira_location"`
	SysID              string `json:"sys_id"`
	ContactType        string `json:"contact_type"`
	CabRequired        string `json:"cab_required"`
	Urgency            string `json:"urgency"`
	Scope              string `json:"scope"`
	Company            string `json:"company"`
	Justification      string `json:"justification"`
	ActivityDue        string `json:"activity_due"`
	Comments           string `json:"comments"`
	Approval           string `json:"approval"`
	DueDate            string `json:"due_date"`
	SysModCount        string `json:"sys_mod_count"`
	UBridgeDetails     string `json:"u_bridge_details"`
	OnHold             string `json:"on_hold"`
	SysTags            string `json:"sys_tags"`
	UJiraCorrelationID string `json:"u_jira_correlation_id"`
	ConflictLastRun    string `json:"conflict_last_run"`
	UJiraDueDate       string `json:"u_jira_due_date"`
	Location           string `json:"location"`
	Risk               string `json:"risk"`
	Category           string `json:"category"`
	RiskImpactAnalysis string `json:"risk_impact_analysis"`
}

type ChangeRequests struct {
	Result []struct {
		ChangeRequest
	} `json:"result"`
}
