package models

type FetchBillRequest struct {
	ChID            int          `json:"chId"`
	IsRealTimeFetch bool         `json:"isRealTimeFetch"`
	CustDetails     CustDetails  `json:"custDetails"`
	AgentDetails    AgentDetails `json:"agentDetails"`
	BillDetail      BillDetail   `json:"billDetails"`
}

type BillDetail struct {
	BillerID       string          `json:"billerId"`
	CustomerParams []CustomerParam `json:"customerParams"`
}

type FetchBillResponse struct {
	RespCode string   `json:"respCode"`
	Status   string   `json:"status"`
	Response RespData `json:"response"`
}

type RespData struct {
	ChID             int             `json:"chId"`
	RefID            string          `json:"refId"`
	ApprovalRefNum   string          `json:"approvalRefNum"`
	ResponseCode     string          `json:"responseCode"`
	ResponseReason   string          `json:"responseReason"`
	ComplianceReason string          `json:"complianceReason"`
	ComplianceRespCd string          `json:"complianceRespCd"`
	BillDetails      []CustomerParam `json:"billDetails"`
	BillerResponse   BillerResponse  `json:"billerResponse"`
	AdditionalInfo   []interface{}   `json:"additionalInfo"`
}

type BillParam struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
